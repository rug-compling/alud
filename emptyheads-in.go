package main

import (
	"sort"
)

func reconstructEmptyHead(node *NodeType, q *Context) {

	doit := false
	var antecedent []interface{}

	/*
	   if ( $node[@rel="hd" and @index and not(@pt or @cat)]  and
	        $antecedent/@word (: onder andere as hd... :)
	        (: and not(local:auxiliary($antecedent) = ("aux","aux:pass","cop")) skip auxiliaries and copulas, prepositions as well? :)
	      )
	*/
	if node.Rel == "hd" && node.Index > 0 && node.Pt == "" && node.Cat == "" {
		// let $antecedent := $node/ancestor::node//node[(@pt or @cat) and @index = $node/@index ]
		antecedent = FIND(node, q, `$allnodes[(@pt or @cat) and @index = $node/@index ]`)
		if TEST(antecedent, q, `$node/@word (: onder andere as hd... :)
		        (: and not(local:auxiliary($node) = ("aux","aux:pass","cop")) skip auxiliaries and copulas, prepositions as well? :)
                   `) {
			doit = true
		}
	}

	if doit {
		others := FIND(node, q, `$node/../node[@pt or @cat]`)
		var end int
		if len(others) > 0 {
			if TEST(node, q, `$node/../node[@pt or @cat]/@begin = $node/../@begin`) {
				end = lastnode(others).End + 1
			} else {
				end = leftEdge(firstnode(others), q) + 1
			}
		} else {
			end = i1(FIND(node, q, `$node/../@end`)) - 999 // covers cases where there is no sister with content
		}
		var copied int
		antenode := firstnode(antecedent)
		if antenode.udCopiedFrom > 0 {
			copied = antenode.udCopiedFrom
		} else {
			copied = antenode.End
		}

		node.Begin = end - 1
		node.End = end
		node.Word = antenode.Word
		node.Lemma = antenode.Lemma
		node.Postag = antenode.Postag
		node.Pt = antenode.Pt
		node.udRelation = "_"
		node.udHeadPosition = UNDERSCORE
		node.udCopiedFrom = copied

		// copied verder alle ud-attributen
		node.udAbbr = antenode.udAbbr
		node.udCase = antenode.udCase
		//node.udCopiedFrom = antenode.udCopiedFrom
		node.udDefinite = antenode.udDefinite
		node.udDegree = antenode.udDegree
		node.udEnhanced = antenode.udEnhanced
		node.udForeign = antenode.udForeign
		node.udGender = antenode.udGender
		//node.udHeadPosition = antenode.udHeadPosition
		node.udNumber = antenode.udNumber
		node.udPerson = antenode.udPerson
		node.udPos = antenode.udPos
		node.udPoss = antenode.udPoss
		node.udPronType = antenode.udPronType
		node.udReflex = antenode.udReflex
		// node.udRelation = antenode.udRelation
		node.udTense = antenode.udTense
		node.udVerbForm = antenode.udVerbForm
		node.udFirstWordBegin = antenode.udFirstWordBegin

		q.ptnodes = append(q.ptnodes, node)
		q.varptnodes = append(q.varptnodes, node)
		sort.Slice(q.ptnodes, func(i, j int) bool {
			return q.ptnodes[i].End < q.ptnodes[j].End
		})
		sort.Slice(q.varptnodes, func(i, j int) bool {
			return q.varptnodes[i].(*NodeType).End < q.varptnodes[j].(*NodeType).End
		})

	} else {
		for _, child := range node.Node {
			reconstructEmptyHead(child, q)
		}
	}
}

func leftEdge(node *NodeType, q *Context) int {
	left := 1000000
	for _, n := range FIND(node, q, `$node/descendant-or-self::node[@pt]`) {
		if begin := n.(*NodeType).Begin; begin < left {
			left = begin
		}
	}
	return left
}
