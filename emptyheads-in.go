// +build ignore

package main

import (
	"sort"
)

func reconstructEmptyHead(q *Context) {
	/*
	   if ( $node[@rel="hd" and @index and not(@pt or @cat)]  and
	        $antecedent/@word (: onder andere as hd... :)
	        (: and not(local:auxiliary($antecedent) = ("aux","aux:pass","cop")) skip auxiliaries and copulas, prepositions as well? :)
	      )
	*/
	seen := make(map[int]bool)
	found := false
	for _, n := range q.varindexnodes {
		node := n.(*NodeType)

		if node.Rel != "hd" || node.Pt != "" || node.Cat != "" {
			continue
		}

		// let $antecedent := $node/ancestor::node//node[(@pt or @cat) and @index = $node/@index ]
		antecedent := FIND(q, `$q.varindexnodes[(@pt or @cat) and @index = $node/@index ]`)
		if !TEST(q, `$antecedent/@word (: onder andere as hd... :)
		        (: and not(local:auxiliary($antecedent) = ("aux","aux:pass","cop")) skip auxiliaries and copulas, prepositions as well? :)
                   `) {
			continue
		}
		found = true

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
		var copied int
		antenode := n1(antecedent)
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

		q.ptnodes = append(q.ptnodes, node)
		q.varptnodes = append(q.varptnodes, node)
	}
	if found {
		sort.Slice(q.ptnodes, func(i, j int) bool {
			return q.ptnodes[i].End < q.ptnodes[j].End
		})
		sort.Slice(q.varptnodes, func(i, j int) bool {
			return q.varptnodes[i].(*NodeType).End < q.varptnodes[j].(*NodeType).End
		})
	}
}

func leftEdge(node *NodeType, q *Context) int {
	left := 1000000
	for _, n := range FIND(q, `$node/descendant-or-self::node[@pt]`) {
		if begin := n.(*NodeType).Begin; begin < left {
			left = begin
		}
	}
	return left
}
