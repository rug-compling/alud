// +build ignore

package main

import (
	"sort"
	"strings"
)

type DepType struct {
	head int
	dep  string
}

func enhancedDependencies(q *Context) {

	reconstructEmptyHead(q)

	// add_Edependency_relations
	for _, node := range q.ptnodes {
		// Edependency_relation
		node.udERelation = dependencyLabel(node, q)
		node.udEHeadPosition = externalHeadPosition(node, q)
	}

	for _, node := range q.ptnodes {
		enhancedDependencies1(node, q)
	}
}

func enhancedDependencies1(node *NodeType, q *Context) {
	// iobj2 control : de commissie kan de raad aanbevelen/adviseren/ X te doen
	// rhd : een levend visje dat doorgeslikt moet worden
	q.depth = 0
	var enhanced []DepType
	for {

		if TEST(q, `$node[@ud:ERelation=("nsubj","obj","iobj","nsubj:pass")]`) { // TODO: klopt dit? exists binnen [ ]
			so := FIND(q,
				`$node/ancestor::node/node[@rel=("su","obj1","obj2") and local:internal_head_position(.) = $node/@end ]/@index`)
			if len(so) > 0 {
				soIndex := i1(so)
				enhanced = []DepType{DepType{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}}
				enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)
				enhanced = append(enhanced, distributeConjuncts(node, q)...)
				enhanced = append(enhanced, distributeDependents(node, q)...)
				enhanced = append(enhanced, xcompControl(node, q, soIndex)...)
				enhanced = append(enhanced, upstairsControl(node, q, soIndex)...)
				enhanced = append(enhanced, passiveVpControl(node, q, soIndex)...)
				break
			}
		}

		rhd := FIND(q,
			`$node/ancestor::node/node[@rel="rhd" and local:internal_head_position(.) = $node/@end ]/@index`)
		if len(rhd) > 0 {
			rhdIndex := i1(rhd)
			rhdNp := FIND(q, `$node/ancestor::node[@cat="np" and node[@rel="mod"]/node[@rel="rhd"]/@index = $rhdIndex]`)
			// de enige _i die voldoet aan de eisen -- make sure empty heads are covered as well
			if len(rhdNp) > 0 {
				enhanced = []DepType{DepType{head: internalHeadPositionWithGapping(rhdNp, q), dep: "ref"}}
				enhanced = append(enhanced, xcompControl(node, q, rhdIndex)...)
				enhanced = append(enhanced, passiveVpControl(node, q, rhdIndex)...)
				break
			}
			// if there is no antecedent, lets keep the basic relation
			enhanced = []DepType{DepType{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}}
			enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)
			enhanced = append(enhanced, distributeConjuncts(node, q)...)
			enhanced = append(enhanced, distributeDependents(node, q)...)
			enhanced = append(enhanced, xcompControl(node, q, rhdIndex)...)
			enhanced = append(enhanced, passiveVpControl(node, q, rhdIndex)...)
			break
		}

		relSister := FIND(q, `$node/../node[@rel="mod" and @cat="rel"]/node[@rel="rhd"]/@index`)
		if len(relSister) > 0 {
			relSisterIndex := i1(relSister)
			enhanced = []DepType{DepType{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}}
			enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)
			enhanced = append(enhanced, distributeConjuncts(node, q)...)
			enhanced = append(enhanced, distributeDependents(node, q)...)
			enhanced = append(enhanced, xcompControl(node, q, relSisterIndex)...)
			enhanced = append(enhanced, passiveVpControl(node, q, relSisterIndex)...)
			break
		}

		if node.udHeadPosition > 0 {
			enhanced = []DepType{DepType{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}}
			enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)
			enhanced = append(enhanced, distributeConjuncts(node, q)...)
			enhanced = append(enhanced, distributeDependents(node, q)...)
			break
		}

		enhanced = []DepType{}
		break
	}

	sort.Slice(enhanced, func(i, j int) bool {
		if enhanced[i].head != enhanced[j].head {
			return enhanced[i].head < enhanced[j].head
		}
		return enhanced[i].dep < enhanced[j].dep
	})
	for i := 1; i < len(enhanced); i++ {
		if enhanced[i].head == enhanced[i-1].head && enhanced[i].dep == enhanced[i-1].dep {
			enhanced = append(enhanced[:i], enhanced[1+i:]...)
			i--
		}
	}
	ss := make([]string, 0, len(enhanced))
	for _, e := range enhanced {
		if e.dep != "" {
			ss = append(ss, number(e.head)+":"+e.dep)
		}
	}
	node.udEnhanced = strings.Join(ss, "|")

}

func join(a, b string) string {
	if b == "" {
		return a
	}
	return a + ":" + b
}

func enhanceDependencyLabel(node *NodeType, q *Context) string {
	/*
	   declare function local:enhance_dependency_label($node as node()) as xs:string {
	    let $label := $node/@ud:ERelation
	    (: find the minimal dominating node :)
	    let $crd := ($node/ancestor::node[@cat="conj" and
	   	       not(.//node[@cat="conj"]//node/@begin = $node/@begin)]/node[@rel="crd"])[1]
	    let $case := ($node/ancestor::node//node[@ud:ERelation="case" and @ud:EHeadPosition=$node/@end])[1]
	    let $mark := ($node/ancestor::node//node[@ud:ERelation=("mark","case") and @ud:EHeadPosition=$node/@end])[1]
	    return
	        1if ($label = "conj" and exists($crd))
	        1then 2if ($crd/@lemma)
	   	      2then string-join(($label,local:enhanced_lemma_string($crd)),':')
	   	      2else 3if ($crd/@cat="mwu")
	   	        3then string-join(($label,local:enhanced_lemma_string(($crd/node[@rel="mwp"])[1])),':')
	   			3else "ERROR_empty_eud_label"
	        1else 4if ($label = ("nmod","obl") and exists($case))
	   	      4then string-join(($label,local:enhanced_lemma_string($case)),':')
	   	      4else 5if ($label = ("advcl","acl") and exists($mark))
	   	        5then string-join(($label,local:enhanced_lemma_string($mark)),':')
	   	        5else 6if ( exists($label) )
	   		      6then $label
	   		      6else "ERROR_empty_eud_label"
	   };
	*/
	label := node.udERelation
	if label == "conj" {
		if crd := n1(FIND(q, `$node/ancestor::node[@cat="conj" and
	       not(.//node[@cat="conj"]//node/@begin = $node/@begin)]/node[@rel="crd"]`)); crd != noNode {
			if crd.Lemma != "" {
				return join(label, enhancedLemmaString1(crd, q))
			}
			if crd.Cat == "mwu" {
				return join(label, enhancedLemmaString1(n1(FIND(q, `$crd/node[@rel="mwp"]`)), q))
			}
			return "ERROR_empty_eud_label"
		}
	}

	if label == "nmod" || label == "obl" {
		if casee := n1(FIND(q, `$node/ancestor::node//node[@ud:ERelation="case" and @ud:EHeadPosition=$node/@end]`)); casee != noNode {
			return join(label, enhancedLemmaString1(casee, q))
		}
	}

	if label == "advcl" || label == "acl" {
		if mark := n1(FIND(q, `$node/ancestor::node//node[@ud:ERelation=("mark","case") and @ud:EHeadPosition=$node/@end]`)); mark != noNode {
			return join(label, enhancedLemmaString1(mark, q))
		}
	}

	if label != "" {
		return label
	}

	return "ERROR_empty_eud_label"
}

func anaphoricRelpronoun(node *NodeType, q *Context) []DepType {
	/*

	   declare function local:anaphoric_relpronoun($node as node()) as node()* {
	   	(: works voor waar, and last() picks waar in 'daar waar' cases :)
	   	(: dont add anything for hij werd voorzitter, wat hij nog steeds is (otherwise self-reference) :)
	   	(: for loop ensures correct result if N has 2 acl:relcl dependents :)
	   	for $anaphoric_relpronoun in
	   			$node/ancestor::node[@cat="np" and local:internal_head_position(.) = $node/@end]/
	   			       node[@rel="mod"]/node[@rel="rhd"]/descendant-or-self::node[@pt="vnw" and not(@ud:HeadPosition = $node/@end)][last()]
	     let $label :=
	       if ($anaphoric_relpronoun/@ud:Relation = ("nsubj","nsubj:pass"))
	       then string-join(($anaphoric_relpronoun/@ud:Relation,"relsubj"),":")
	       else if ($anaphoric_relpronoun/@ud:Relation = ("obj","obl"))
	   	 then string-join(($anaphoric_relpronoun/@ud:Relation,"relobj"),":")
	   	 else $anaphoric_relpronoun/@ud:Relation
	   	return
	   	<headdep head="{$anaphoric_relpronoun/@ud:HeadPosition}" dep="{$label}" />
	   };
	*/

	// TODO: for-loop over 1 element (last)?

	// works voor waar, and last() picks waar in 'daar waar' cases
	// dont add anything for hij werd voorzitter, wat hij nog steeds is (otherwise self-reference)
	// for loop ensures correct result if N has 2 acl:relcl dependents
	list := FIND(q, `$node/ancestor::node[@cat="np" and local:internal_head_position(.) = $node/@end]/
	   			       node[@rel="mod"]/node[@rel="rhd"]/descendant-or-self::node[@pt="vnw" and not(@ud:HeadPosition = $node/@end)]`)
	if len(list) > 0 {
		anrel := nZ(list)
		var label string
		if r := anrel.udRelation; r == "nsubj" || r == "nsubj:pass" {
			label = r + "relsubj"
		} else if r == "obj" || r == "obl" {
			label = r + ":relobj"
		} else {
			label = r
		}
		return []DepType{DepType{head: anrel.udHeadPosition, dep: label}}
	}
	return []DepType{}
}

func distributeConjuncts(node *NodeType, q *Context) []DepType {
	/*
	   (: Glastra en Terlouw verzonnen een list --> nsubj(verzonnen,Glastra) nsubj(verzonnen,Terlouw) :)
	   declare function local:distribute_conjuncts($node as node()) as node()* {
	   	let $coord_head := $node/ancestor::node//node[@end = $node/@ud:HeadPosition
	   	       and @ud:Relation=("amod","appos","nmod","nsubj","nsubj:pass","nummod","obj","iobj","obl","obl:agent","advcl")]
	     (: let $case := ($node/ancestor::node//node[@ud:Relation="case" and @ud:HeadPosition=$coord_head/@end])[1] :)
	     let $deplabel :=
	       (: if ($coord_head/@ud:Relation=("nmod","obl") and $case)
	       then string-join($coord_head/@ud:Relation,local:enhanced_lemma_string($case))
	       else :) local:enhance_dependency_label($coord_head)
	   	return
	   	if ($node[@ud:Relation="conj"] and exists($coord_head))
	   	(: in A en B vs in A en naast B --> use enh_dep_label($node) in the latter case... :)
	   	then <headdep head="{$coord_head/@ud:HeadPosition}" dep="{$deplabel}"/>
	   	else ()
	   };
	*/
	if node.udRelation == "conj" {
		coordHead := n1(FIND(q, `$node/ancestor::node//node[@end = $node/@ud:HeadPosition
	       and @ud:Relation=("amod","appos","nmod","nsubj","nsubj:pass","nummod","obj","iobj","obl","obl:agent","advcl")]`))
		if coordHead != noNode {
			// in A en B vs in A en naast B --> use enh_dep_label($node) in the latter case...
			depLabel := enhanceDependencyLabel(coordHead, q)
			return []DepType{DepType{head: coordHead.udHeadPosition, dep: depLabel}}
		}
	}
	return []DepType{}
}

func distributeDependents(node *NodeType, q *Context) []DepType {
	// TODO
	return []DepType{}
}

func xcompControl(node *NodeType, q *Context, idx int) []DepType {
	// TODO
	return []DepType{}
}

func upstairsControl(node *NodeType, q *Context, idx int) []DepType {
	// TODO
	return []DepType{}
}

func passiveVpControl(node *NodeType, q *Context, idx int) []DepType {
	// TODO
	return []DepType{}
}

func enhancedLemmaString(nodes []interface{}, q *Context) string {
	/*
		sort.Slice(nodes, func(i, j int) bool {
			// TODO" NodeType heeft geen head
			return nodes[i].(*NodeType).head < nodes[j].(*NodeType).head
		})
	*/
	lemmas := make([]string, len(nodes))
	for i, node := range nodes {
		lemmas[i] = enhancedLemmaString1(node.(*NodeType), q)
	}
	return strings.Join(lemmas, "_")
}

func enhancedLemmaString1(node *NodeType, q *Context) string {
	var lemma string
	switch node.Lemma {
	case "a.k.a":
		lemma = "also_known_as"
	case "c.q.":
		lemma = "casu_quo"
	case "dwz.", "d.w.z.":
		lemma = "dat_wil_zeggen"
	case "e.d.":
		lemma = "en_dergelijke"
	case "en/of":
		lemma = "en_of"
	case "enz.":
		lemma = "enzovoort"
	case "etc.":
		lemma = "etcetera"
	case "m.a.w.":
		lemma = "met_andere_woorden"
	case "nl.":
		lemma = "namelijk"
	case "resp.":
		lemma = "respectievelijk"
	case "t/m":
		lemma = "tot_en_met"
	case "t.a.v.":
		lemma = "ten_aanzien_van"
	case "t.g.v.":
		lemma = "ten_gunste_van"
	case "t.n.v.":
		lemma = "ten_name_van"
	case "t.o.v.":
		lemma = "ten_opzichte_van"
	default:
		lemma = node.Lemma
	}
	if fixed := n1(FIND(q, `$node/../node[@ud:ERelation="fixed"]`)); fixed.Lemma != "" {
		lemma += "_" + fixed.Lemma
	}
	lemma = strings.Replace(lemma, "/", "schuine_streep", -1)
	lemma = strings.Replace(lemma, "-", "_", -1)
	return strings.ToLower(lemma)
}
