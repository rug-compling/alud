// +build ignore

package alud

import (
	"github.com/rug-compling/alpinods"
)

func reconstructEmptyHead(q *context) bool {
	seen := make(map[int]bool)
	found := false
	for _, n := range q.varindexnodes {
		node := n.(*nodeType)

		if node.Rel != "hd" || node.Pt != "" || node.Cat != "" {
			continue
		}

		antecedent := FIND(q, `$q.varindexnodes[(@pt or @cat) and @index = $node/@index ]`)
		if !TEST(q, `$antecedent[@word or @cat = "mwu"] (: onder andere as hd... :)
		        (: and not(local:auxiliary($antecedent) = ("aux","aux:pass","cop")) skip auxiliaries and copulas, prepositions as well? :)
                   `) {
			continue
		}
		found = true

		antenode := n1(antecedent)
		mwu := validMwu(antenode)

		others := FIND(q, `$node/../node[@pt or @cat]`)
		var end int
		if len(others) > 0 {
			if TEST(q, `$node/../node[@pt or @cat]/@begin = $node/../@begin`) {
				end = nZ(others).End + 1 // + 0.1
			} else {
				end = leftEdge(n1(others), q) + 1 // + 0.1
			}
		} else {
			end = i1(FIND(q, `$node/../@end`)) - 999 // - 0.9 // covers cases where there is no sister with content
		}
		for seen[end] {
			end++
		}
		seen[end] = true

		end2 := end
		if mwu {
			end2 += len(antenode.Node) - 1
		}

		var copied int
		if antenode.udCopiedFrom > 0 {
			copied = antenode.udCopiedFrom
		} else {
			copied = antenode.End
		}

		node.udOldState = &nodeType{
			NodeAttributes: alpinods.NodeAttributes{
				Begin:  node.Begin,
				Cat:    node.Cat,
				End:    node.End,
				Word:   node.Word,
				Lemma:  node.Lemma,
				Postag: node.Postag,
				Pt:     node.Pt,
			},
			Node: node.Node,
		}

		node.Begin = end - 1
		node.End = end2
		node.Word = antenode.Word
		node.Lemma = antenode.Lemma
		node.Postag = antenode.Postag
		node.Pt = antenode.Pt
		node.Cat = antenode.Cat
		node.udRelation = "_"
		node.udHeadPosition = underscore
		node.udCopiedFrom = copied

		// kopieer verder alle ud-attributen
		node.udAbbr = antenode.udAbbr
		node.udCase = antenode.udCase
		//niet: node.udCopiedFrom = antenode.udCopiedFrom
		node.udDefinite = antenode.udDefinite
		node.udDegree = antenode.udDegree
		node.udEnhanced = antenode.udEnhanced
		node.udForeign = antenode.udForeign
		node.udGender = antenode.udGender
		//niet: node.udHeadPosition = antenode.udHeadPosition
		node.udNumber = antenode.udNumber
		node.udPerson = antenode.udPerson
		node.udPos = antenode.udPos
		node.udPoss = antenode.udPoss
		node.udPronType = antenode.udPronType
		node.udReflex = antenode.udReflex
		//niet: node.udRelation = antenode.udRelation
		node.udTense = antenode.udTense
		node.udVerbForm = antenode.udVerbForm
		node.udFirstWordBegin = antenode.udFirstWordBegin
		node.udERelation = antenode.udERelation
		node.udEHeadPosition = antenode.udEHeadPosition

		if mwu {
			node.Node = make([]*nodeType, len(antenode.Node))
			for i, n := range antenode.Node {
				var copied int
				if n.udCopiedFrom > 0 {
					copied = n.udCopiedFrom
				} else {
					copied = n.End
				}
				if i > 0 {
					end++
					seen[end] = true
				}
				n2 := new(nodeType)
				*n2 = *n
				n2.Begin = end - 1
				n2.End = end
				n2.ID = node.ID + 1 + i
				n2.udRelation = "_"
				n2.udHeadPosition = underscore
				n2.udCopiedFrom = copied
				node.Node[i] = n2
			}
		}
	}
	if found {
		inspect(q)
	}
	return found
}

func leftEdge(node *nodeType, q *context) int {
	left := 1000000
	for _, n := range FIND(q, `$node/descendant-or-self::node[@pt]`) {
		if begin := n.(*nodeType).Begin; begin < left {
			left = begin
		}
	}
	return left
}

func validMwu(node *nodeType) bool {
	if node.Cat != "mwu" {
		return false
	}

	if node.Node == nil || len(node.Node) == 0 {
		return false
	}

	for _, n := range node.Node {
		/*
			if i > 0 && node.Node[i-1].End != n.Begin {
				return false
			}
		*/
		if n.Rel != "mwp" {
			return false
		}
		if n.Word == "" {
			return false
		}
	}

	return true
}
