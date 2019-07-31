/*

Dit is een bewerking van fixpunct.py, geschreven door Martin Popel, onderdeel van udapi.

Links:

https://udapi.github.io/
https://udapi.readthedocs.io/en/latest/udapi.block.ud.html#module-udapi.block.ud.fixpunct

Zie ook:

https://udapi.readthedocs.io/en/latest/udapi.core.html#udapi.core.node.Node
https://udapi.readthedocs.io/en/latest/udapi.core.html#udapi.core.root.Root

*/

package alud

import (
	"sort"
	"strings"
)

type pNode struct {
	f         *nodeType
	ord       int
	key1      int
	key2      int
	punctType string
	parent    *pNode
	children  []*pNode
}

var (
	pairedPunct = map[string]string{
		"(":  ")",
		"[":  "]",
		"{":  "}",
		"<":  ">",
		",,": "''",
		"\"": "\"",
		"'":  "'",
		"`":  "'",
		"“":  "”",
		"‘":  "’",
		"„":  "”",
		"”":  "”",
		"«":  "»",
		"‹":  "›",
	}
	finalPunct = ".?!"
)

func fixpunct(q *context) {

	nodes := make([]*pNode, 0)
	rootDescendants := make([]*pNode, 0)
	var root *pNode

	sort.Slice(q.ptnodes, func(i, j int) bool {
		return q.ptnodes[i].End < q.ptnodes[i].End
	})

	ord := 0
	for _, line := range q.ptnodes {
		if line.udCopiedFrom <= 0 {
			ord++
			nodes = append(nodes, &pNode{
				f:        line,
				children: make([]*pNode, 0),
				ord:      ord,
			})
			rootDescendants = append(rootDescendants, nodes[ord-1])
		}
	}
	for _, node := range rootDescendants {
		if node.f.udHeadPosition == 0 {
			root = node
			node.parent = &pNode{f: noNode, children: make([]*pNode, 0)}
			continue
		}
		for _, n2 := range rootDescendants {
			if node.f.udHeadPosition == n2.f.End {
				n2.children = append(n2.children, node)
				node.parent = n2
				break
			}
		}
		if node.parent == nil {
			// TODO: Dit zou niet moeten gebeuren!
			node.parent = &pNode{f: noNode, children: make([]*pNode, 0)}
		}
	}

	if root == nil {
		// Ongeldige parse
		return
	}

	/*
	   # First, make sure no PUNCT has children
	   for node in root.descendants:
	       while node.parent.upos == "PUNCT":
	           node.parent = node.parent.parent
	*/
	for _, node := range rootDescendants {
		for node.parent.f.udPos == "PUNCT" {
			setParent(node, node.parent.parent)
		}
	}

	/*
	   # Second, fix paired punctuations: quotes and brackets, marking them in _punct_type.
	   # This should be done before handling the subordinate punctuation,
	   # in order to prevent non-projectivities e.g. in dot-before-closing-quote style sentences:
	   #   I call him "Bob."
	   # Here both quotes and the sentence-final dot should be attached to "Bob".
	   # (As you can see on the previous line, I don't like this American typographic rule.)
	   self._punct_type = [None] * (1 + len(root.descendants))
	   for node in root.descendants:
	       if self._punct_type[node.ord] != 'closing':
	           closing_punct = PAIRED_PUNCT.get(node.form, None)
	           if closing_punct is not None:
	               self._fix_paired_punct(root, node, closing_punct)
	*/
	for _, node := range rootDescendants {
		if node.punctType != "closing" {
			if closingPunct, ok := pairedPunct[node.f.Word]; ok {
				fixPairedPunct(node, closingPunct, nodes, rootDescendants)
			}
		}
	}

	/*
	   # Third, fix subordinate punctuation (i.e. any punctuation not marked in _punct_type).
	   for node in root.descendants:
	       if node.upos == "PUNCT" and not self._punct_type[node.ord]:
	           self._fix_subord_punct(node)
	*/
	for _, node := range rootDescendants {
		if node.f.udPos == "PUNCT" && node.punctType == "" {
			fixSubordPunct(node, rootDescendants)
		}
	}

	/*
	   # Finally, check if root is still marked with deprel=root.
	   # This may not hold if the original root was a paired punctuation, which was rehanged.
	   for node in root.children:
	       if node.udeprel != 'root':
	           node.udeprel = 'root'
	           for another_node in root.descendants:
	               if another_node.parent != root and another_node.udeprel == 'root':
	                   another_node.udeprel = 'punct'
	*/
	// root.children zijn niet de children van node=root, maar de toppen van de boom, en dat is er maar 1
	// TODO: klopt dit?
	if root.f.udRelation != "root" {
		root.f.udRelation = "root"
		for _, node2 := range rootDescendants {
			if node2.parent.f.End > 0 && node2.f.udRelation == "root" {
				node2.f.udRelation = "punct"
			}
		}
	}

	// update DEPS
	for _, node := range rootDescendants {
		if node.f.udPos == "PUNCT" {
			node.f.udEnhanced = number(node.f.udHeadPosition) + ":" + node.f.udRelation
		}
	}
}

func fixPairedPunct(openingNode *pNode, closingPunct string, nodes, rootDescendants []*pNode) {
	/*
	   if self.check_paired_punct_upos and opening_node.upos != 'PUNCT':
	       return

	   nested_level = 0
	   for node in root.descendants[opening_node.ord:]:
	       if node.form == closing_punct:
	           if nested_level > 0:
	               nested_level -= 0 # ??????
	           else:
	               self._fix_pair(root, opening_node, node)
	               return
	       elif node.form == opening_node.form:
	           nested_level += 1
	*/

	if openingNode.f.udPos != "PUNCT" {
		return
	}

	nestedLevel := 0
	for _, node := range rootDescendants[openingNode.ord:] {
		if node.f.Word == closingPunct {
			if nestedLevel > 0 {
				nestedLevel-- // ?????? (zie boven)
			} else {
				fixPair(openingNode, node, rootDescendants)
				return
			}
		} else if node.f.Word == openingNode.f.Word {
			nestedLevel++
		}
	}

}

func fixPair(openingNode, closingNode *pNode, rootDescendants []*pNode) {
	/*
	   heads = []
	   punct_heads = []
	   for node in root.descendants[opening_node.ord: closing_node.ord - 1]:
	       if node.parent.precedes(opening_node) or closing_node.precedes(node.parent):
	           if node.upos == 'PUNCT':
	               punct_heads.append(node)
	           else:
	               heads.append(node)
	*/
	heads := make([]*pNode, 0)
	punctHeads := make([]*pNode, 0)
	for _, node := range rootDescendants[openingNode.ord : closingNode.ord-1] {
		if node.parent.ord < openingNode.ord || closingNode.ord < node.parent.ord {
			if node.f.udPos == "PUNCT" {
				punctHeads = append(punctHeads, node)
			} else {
				heads = append(heads, node)
			}
		}
	}

	/*
	   # Punctuation should not have children, but if there is no other head candidate,
	   # let's break this rule.
	   if len(heads) == 0:
	       heads = punct_heads
	*/
	if len(heads) == 0 {
		heads = punctHeads
		if len(heads) == 0 {
			return
		}
	}

	/*
	   if len(heads) == 1:
	       opening_node.parent = heads[0]
	       closing_node.parent = heads[0]
	       self._punct_type[opening_node.ord] = 'opening'
	       self._punct_type[closing_node.ord] = 'closing'
	   elif len(heads) > 1:
	       opening_node.parent = sorted(heads, key=lambda n: n.descendants(add_self=1)[0].ord)[0]
	       closing_node.parent = sorted(heads, key=lambda n: -n.descendants(add_self=1)[-1].ord)[0]
	       self._punct_type[opening_node.ord] = 'opening'
	       self._punct_type[closing_node.ord] = 'closing'
	*/
	if len(heads) == 1 {
		setParent(openingNode, heads[0])
		setParent(closingNode, heads[0])
	} else {
		for _, head := range heads {
			desc := getDescendants(head, true)
			head.key1 = desc[0].ord
			head.key2 = -desc[len(desc)-1].ord
		}
		sort.Slice(heads, func(i, j int) bool {
			return heads[i].key1 < heads[j].key1
		})
		setParent(openingNode, heads[0])
		sort.Slice(heads, func(i, j int) bool {
			return heads[i].key2 < heads[j].key2
		})
		setParent(closingNode, heads[0])
	}
	openingNode.punctType = "opening"
	closingNode.punctType = "closing"
}

func fixSubordPunct(node *pNode, rootDescendants []*pNode) {
	/*
	   # Dot used as the ordinal-number marker (in some languages) or abbreviation marker.
	   # TODO: detect these cases somehow
	   # Numbers can be detected with `node.parent.form.isdigit()`,
	   # but abbreviations are more tricky because the Abbr=Yes feature is not always used.
	   if node.form == '.' and node.parent == node.prev_node:
	       return
	*/
	if node.f.Word == "." && node.parent.ord == node.ord-1 {
		return
	}

	/*
	   # Even non-paired punctuation like commas and dashes may work as paired.
	   # Detect such cases and try to preserve, but only if projective.
	   p_desc = node.parent.descendants(add_self=1)
	   if node in (p_desc[0], p_desc[-1]) and len(p_desc) == p_desc[-1].ord - p_desc[0].ord + 1:
	       if (p_desc[0].upos == 'PUNCT' and p_desc[-1].upos == 'PUNCT'
	               and p_desc[0].parent == node.parent and p_desc[-1].parent == node.parent):
	           return
	*/
	pDesc := getDescendants(node.parent, true)
	pDescLen := len(pDesc)
	pDesc0 := pDesc[0]
	pDesc1 := pDesc[pDescLen-1]
	if (node == pDesc0 || node == pDesc1) && pDescLen == pDesc1.ord-pDesc0.ord+1 {
		if pDesc0.f.udPos == "PUNCT" && pDesc1.f.udPos == "PUNCT" &&
			pDesc0.parent == node.parent && pDesc1.parent == node.parent {
			return
		}
	}

	/*
	   # Initialize the candidates (left and right) with the nearest nodes excluding punctuation.
	   # Final punctuation should not be attached to any following, so exclude r_cand there.
	   l_cand, r_cand = node.prev_node, node.next_node
	   if node.form in FINAL_PUNCT:
	       r_cand = None
	   while l_cand.ord > 0 and l_cand.upos == "PUNCT":
	       if self._punct_type[l_cand.ord] == 'opening':
	           l_cand = None
	           break
	       l_cand = l_cand.prev_node
	   while r_cand is not None and r_cand.upos == "PUNCT":
	       if self._punct_type[r_cand.ord] == 'closing':
	           r_cand = None
	           break
	       r_cand = r_cand.next_node
	*/
	var lCand, rCand *pNode
	if node.ord > 1 {
		lCand = rootDescendants[node.ord-2]
	}
	if node.ord < len(rootDescendants) {
		rCand = rootDescendants[node.ord]
	}
	if strings.Contains(finalPunct, node.f.Word) {
		rCand = nil
	}
	for lCand != nil && lCand.f.udPos == "PUNCT" {
		if lCand.punctType == "opening" {
			lCand = nil
			break
		}
		if lCand.ord > 1 {
			lCand = rootDescendants[lCand.ord-2]
		} else {
			lCand = nil
		}
	}
	for rCand != nil && rCand.f.udPos == "PUNCT" {
		if rCand.punctType == "closing" {
			rCand = nil
			break
		}
		if rCand.ord < len(rootDescendants) {
			rCand = rootDescendants[rCand.ord]
		} else {
			rCand = nil
		}
	}

	/*
	   # Climb up from the candidates, until we would reach the root or "cross" the punctuation.
	   # If the candidates' descendants span across the punctuation, we also stop
	   # because climbing higher would cause a non-projectivity (the punct would be the gap).
	   l_path, r_path = [l_cand], [r_cand]
	   if l_cand is None or l_cand.is_root():
	       l_cand = None
	   else:
	       while (not l_cand.parent.is_root() and l_cand.parent.precedes(node)
	              and not node.precedes(l_cand.descendants(add_self=1)[-1])):
	           l_cand = l_cand.parent
	           l_path.append(l_cand)
	   if r_cand is not None:
	       while (not r_cand.parent.is_root() and node.precedes(r_cand.parent)
	              and not r_cand.descendants(add_self=1)[0].precedes(node)):
	           r_cand = r_cand.parent
	           r_path.append(r_cand)
	*/
	lPath := make([]*pNode, 1)
	rPath := make([]*pNode, 1)
	lPath[0] = lCand
	rPath[0] = rCand
	if lCand == nil || lCand.f.End <= 0 {
		lCand = nil
	} else {
		seen := make(map[int]bool)
		for {
			if lCand.parent.f.End <= 0 || lCand.parent.ord >= node.ord {
				break
			}
			desc := getDescendants(lCand, true)
			if node.ord < desc[len(desc)-1].ord {
				break
			}
			lCand = lCand.parent
			lPath = append(lPath, lCand)
			if id := lCand.ord; seen[id] {
				break
			} else {
				seen[id] = true
			}
		}
	}
	if rCand != nil {
		seen := make(map[int]bool)
		for {
			if rCand.parent.f.End <= 0 || node.ord >= rCand.parent.ord {
				break
			}
			desc := getDescendants(rCand, true)
			if desc[0].ord < node.ord {
				break
			}
			rCand = rCand.parent
			rPath = append(rPath, rCand)
			if id := rCand.ord; seen[id] {
				break
			} else {
				seen[id] = true
			}
		}
	}

	/*
	   # Now select between l_cand and r_cand -- which will be the new parent?
	   # The lower one. Note that if neither is descendant of the other and neither is None
	   # (which can happen in rare non-projective cases), we arbitrarily prefer l_cand,
	   # but if the original parent is either on l_path or r_path, we keep it as acceptable.
	   if l_cand is not None and l_cand.is_descendant_of(r_cand):
	       cand, path = l_cand, l_path
	   elif r_cand is not None and r_cand.is_descendant_of(l_cand):
	       cand, path = r_cand, r_path
	   elif l_cand is not None:
	       cand, path = l_cand, l_path + r_path
	   elif r_cand is not None:
	       cand, path = r_cand, l_path + r_path
	   else:
	       return
	*/
	var cand *pNode
	var path []*pNode
	if lCand != nil && isDescendantOf(lCand, rCand) {
		cand, path = lCand, lPath
	} else if rCand != nil && isDescendantOf(rCand, lCand) {
		cand, path = rCand, rPath
	} else if lCand != nil {
		cand, path = lCand, append(lPath, rPath...)
	} else if rCand != nil {
		cand, path = rCand, append(lPath, rPath...)
	} else {
		return
	}

	/*
	   # The guidelines say:
	   #    Within the relevant unit, a punctuation mark is attached
	   #    at the highest possible node that preserves projectivity.
	   # However, sometimes it is difficult to detect the unit (and its head).
	   # E.g. in "Der Mann, den Sie gestern kennengelernt haben, kam wieder."
	   # the second comma should depend on "kennengelernt", not on "Mann"
	   # because the unit is just the relative clause.
	   # We try to be conservative and keep the parent, unless we are sure it is wrong.
	   if node.parent not in path:
	       node.parent = cand
	   node.deprel = "punct"
	*/
	if !contains(path, node.parent) {
		setParent(node, cand)
	}
	node.f.udRelation = "punct"
}

func setParent(node, parent *pNode) {
	// verwijderen van oude parent
	for i, n := range node.parent.children {
		if n == node {
			node.parent.children = append(node.parent.children[:i], node.parent.children[i+1:]...)
			break
		}
	}

	node.parent = parent
	node.f.udHeadPosition = parent.f.End

	parent.children = append(parent.children, node)
	sort.Slice(parent.children, func(i, j int) bool {
		return parent.children[i].ord < parent.children[j].ord
	})
}

func getDescendants(node *pNode, addSelf bool) []*pNode {
	nodes := make([]*pNode, 0)
	if addSelf {
		nodes = append(nodes, node)
	}
	seen := make(map[int]bool)
	var desc func(*pNode)
	desc = func(n *pNode) {
		nodes = append(nodes, n.children...)
		for _, child := range n.children {
			// loops zouden niet moeten mogen, maar zijn er toch...
			if id := child.f.End; !seen[id] {
				seen[id] = true
				desc(child)
			}
		}
	}
	desc(node)
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].ord < nodes[j].ord
	})
	return nodes
}

func isDescendantOf(node1, node2 *pNode) bool {
	if node2 == nil {
		return false
	}
	seen := make(map[int]bool)
	var f func([]*pNode) bool
	f = func(nodes []*pNode) bool {
		for _, node := range nodes {
			if node1 == node {
				return true
			}
		}
		for _, node := range nodes {
			if id := node.f.End; !seen[id] {
				seen[id] = true
				if f(node.children) {
					return true
				}
			}
		}
		return false
	}
	return f(node2.children)
}

func contains(nodes []*pNode, node *pNode) bool {
	for _, n := range nodes {
		if node == n {
			return true
		}
	}
	return false
}
