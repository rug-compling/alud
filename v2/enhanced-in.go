//go:build ignore
// +build ignore

package alud

import (
	"fmt"
	"sort"
	"strings"
)

type depT struct {
	head int
	dep  string
}

func enhancedDependencies(q *context) {

	var node *nodeType

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "enhancedDependencies", q, node))
		}
	}()

	changed := reconstructEmptyHead(q)
	//dump(q.alpino)

	// add_Edependency_relations
	for _, node = range q.ptnodes {
		// Edependency_relation
		if changed {
			q.depth = 0
			node.udERelation = dependencyLabel(node, q)
			q.depth = 0
			node.udEHeadPosition = externalHeadPosition(list(node), q)
		} else {
			node.udERelation = node.udRelation
			node.udEHeadPosition = node.udHeadPosition
		}
	}

	for _, node = range q.ptnodes {
		enhancedDependencies1(node, q)
	}
}

func enhancedDependencies1(node *nodeType, q *context) {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "enhancedDependencies1", q, node))
		}
	}()

	// iobj2 control : de commissie kan de raad aanbevelen/adviseren/ X te doen
	// rhd : een levend visje dat doorgeslikt moet worden
	q.depth = 0
	var enhanced []depT
	for {

		if TEST(q, `$node[@ud:ERelation=("nsubj","obj","iobj","nsubj:pass")]`) { // TODO: klopt dit? exists binnen [ ]
			so := FIND(q,
				`$node/ancestor::node/node[@rel=("su","obj1","obj2") and local:internal_head_position(.) = $node/@end ]/@index`)
			if len(so) > 0 {
				soIndex := i1(so)
				//NP
				enhanced = []depT{{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}} // self
				enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)                          // self
				//NP
				enhanced = append(enhanced, distributeConjuncts(node, q)...) // self
				//NP
				enhanced = append(enhanced, distributeDependents(node, q)...) // self
				//NP
				enhanced = append(enhanced, xcompControl(node, q, soIndex)...)
				//NP
				enhanced = append(enhanced, upstairsControl(node, q, soIndex)...)
				//NP
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
				//NP
				enhanced = []depT{{head: internalHeadPositionWithGapping(rhdNp, q), dep: "ref"}} // rhdref
				//NP
				enhanced = append(enhanced, xcompControl(node, q, rhdIndex)...)
				//NP
				enhanced = append(enhanced, passiveVpControl(node, q, rhdIndex)...)
				break
			}
			// if there is no antecedent, lets keep the basic relation
			//NP
			enhanced = []depT{{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}} // self
			enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)                          // self
			//NP
			enhanced = append(enhanced, distributeConjuncts(node, q)...) // self
			//NP
			enhanced = append(enhanced, distributeDependents(node, q)...) // self
			//NP
			enhanced = append(enhanced, xcompControl(node, q, rhdIndex)...)
			//NP
			enhanced = append(enhanced, passiveVpControl(node, q, rhdIndex)...)
			break
		}

		relSister := FIND(q, `($node[@rel="hd"]/../node[@rel="mod" and @cat="rel"]/node[@rel="rhd"]/@index)[1]`) // get rid of xsubj on det etc non head elements
		if len(relSister) > 0 {
			relSisterIndex := i1(relSister)
			//NP
			enhanced = []depT{{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}} // self
			enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)                          // self
			//NP
			enhanced = append(enhanced, distributeConjuncts(node, q)...) // self
			//NP
			enhanced = append(enhanced, distributeDependents(node, q)...) // self
			//NP
			enhanced = append(enhanced, xcompControl(node, q, relSisterIndex)...)
			//NP
			enhanced = append(enhanced, passiveVpControl(node, q, relSisterIndex)...)
			break
		}

		relSister2 := FIND(q, `($node[@rel="mwp" and @begin = ../@begin]/../../node[@rel="mod" and @cat="rel"]/node[@rel="rhd"]/@index)[1]`)
		// heads of MWU, could not get disjunction of paths (|) or disjunctive condition in go language to work
		if len(relSister2) > 0 {
			relSisterIndex := i1(relSister2)
			//NP
			enhanced = []depT{{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}} // self
			enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)                          // self
			//NP
			enhanced = append(enhanced, distributeConjuncts(node, q)...) // self
			//NP
			enhanced = append(enhanced, distributeDependents(node, q)...) // self
			//NP
			enhanced = append(enhanced, xcompControl(node, q, relSisterIndex)...)
			//NP
			enhanced = append(enhanced, passiveVpControl(node, q, relSisterIndex)...)
			break
		}

		// underscore is resultaat van reconstructEmptyHead()
		if node.udHeadPosition >= 0 || node.udHeadPosition == underscore {
			//NP
			enhanced = []depT{{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}} // self
			enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)                          // self
			//NP
			enhanced = append(enhanced, distributeConjuncts(node, q)...) // self
			//NP
			enhanced = append(enhanced, distributeDependents(node, q)...) // self
			break
		}

		//NP
		enhanced = []depT{{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}}
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
		if e.head < 0 ||
			e.head == 0 && e.dep != "root" ||
			e.head != 0 && e.dep == "root" ||
			e.dep == "orphan" {
			panic(fmt.Sprintf("Invalid EUD %s:%s", number(e.head), e.dep))
		}
		if e.head == node.End {
			panic(fmt.Sprintf(
				"EUD to self %s:%s",
				number(e.head),
				e.dep))
		}
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

func enhanceDependencyLabel(node *nodeType, q *context) string {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "enhanceDependencyLabel", q, node))
		}
	}()

	label := node.udERelation
	if label == "conj" {
		// TODO: selecting last() crd sister works for 'zowel X als Y cases (zowel=preconj),
		// requiring that crd does not follow the conjunct works for 'A en B of C' cases
		// ... want zijn ogen waren niet verlicht maar verduisterd -- pick want, not niet
		// superfluous if we use last()?? and  not(.//node[@cat="conj"]//node/@begin = $node/@begin)
		// TODO: A, B, en C --> B should get 'conj:en' as label, check for presence of comma?
		// should generalize to cases where constituent B is preceded by ',' : now approximation: check for next level up node as well (no guarantee node is actually head of this phrase..)
		if crd := n1(FIND(q, `($node/ancestor::node[@cat="conj" 
		]/node[@rel="crd" and not(@ud:Relation="cc:preconj") and (not(@end > $node/@begin) or $node/@begin = $node/ancestor::node[@rel="top"]/node[@word=","]/@end or $node/../@begin = $node/ancestor::node[@rel="top"]/node[@word=","]/@end )])[last()]`)); crd != noNode {
			if crd.Lemma != "" {
				return join(label, enhancedLemmaString1(crd, q))
			}

			if crd.Cat == "mwu" {
				return join(label, enhancedLemmaString1(n1(FIND(q, `($crd/node[@rel="mwp"])[1]`)), q))
			}

			if crd.Lemma == "" { // cases where the crd element is co-indexed
				// zon 200 joodse en 120 arabische doden en gewonden
				return label
			}
			panic("Empty EUD label")
		}
	}

	if label == "nmod" || label == "obl" || label == "obl:arg" {
		if casee := n1(FIND(q, `($q.varptnodes[@ud:ERelation="case" and @ud:EHeadPosition=$node/@end])[1]`)); casee != noNode {
			return join(label, enhancedLemmaString1(casee, q))
		}
	}

	if label == "advcl" || label == "acl" {
		if mark := n1(FIND(q, `($q.varptnodes[@ud:ERelation=("mark","case") and @ud:EHeadPosition=$node/@end])[1]`)); mark != noNode {
			return join(label, enhancedLemmaString1(mark, q))
		}
	}

	if label != "" {
		return label
	}

	panic("Empty EUD label")
}

func anaphoricRelpronoun(node *nodeType, q *context) []depT {
	// works voor waar, and last() picks waar in 'daar waar' cases
	// dont add anything for hij werd voorzitter, wat hij nog steeds is (otherwise self-reference)
	// for loop ensures correct result if N has 2 acl:relcl dependents
	result := []depT{}
	for _, a := range FIND(q, `$node/ancestor::node[@cat="np" and local:internal_head_position(.) = $node/@end]/
	   			       node[@rel="mod"]/node[@rel="rhd"]/descendant-or-self::node[@pt="vnw" and not(@ud:HeadPosition = $node/@end)][last()]`) {
		anrel := a.(*nodeType)
		var label string
		if r := anrel.udRelation; r == "nsubj" || r == "nsubj:pass" {
			label = r // + ":relsubj"  validation suggests relsubj/obj is superfluous
		} else if r == "obj" || r == "obl" {
			label = r // + ":relobj"
		} else {
			label = r
		}
		result = append(result, depT{head: anrel.udHeadPosition, dep: label})
	}
	return result
}

// Glastra en Terlouw verzonnen een list --> nsubj(verzonnen,Glastra) nsubj(verzonnen,Terlouw)
func distributeConjuncts(node *nodeType, q *context) []depT {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "distributeConjuncts", q, node))
		}
	}()

	// ensure obl:agent is only added if conj is introduced by 'door' GB 17/04/23
	if node.udRelation == "conj" {
		coordHead := n1(FIND(q, `$q.varallnodes[@end = $node/@ud:HeadPosition
		and @ud:Relation=("amod","appos","nmod","nsubj","nsubj:pass","nummod","obj","iobj","obl","obl:agent","obl:arg","advcl")]`))
		casee := n1(FIND(q, `($q.varptnodes[@ud:ERelation="case" and @ud:EHeadPosition=$node/@end])[1]`))
		caselemma := enhancedLemmaString1(casee, q)
		if coordHead.udRelation == "obl:agent" && caselemma == "door" {
			depLabel := "obl:agent"
			return []depT{{head: coordHead.udHeadPosition, dep: depLabel}}
		} else if coordHead != noNode {
			// in A en B vs in A en naast B --> use enh_dep_label($node) in the latter case...
			depLabel := enhanceDependencyLabel(coordHead, q)
			return []depT{{head: coordHead.udHeadPosition, dep: depLabel}}
		}
	}
	return []depT{}
}

// de onrust kan een reis vertragen of frustreren  --> obj(vertragen,reis) obj(frustreren,reis)
// todo: passives ze werd ontmanteld en verkocht  su coindexed with two obj1
// done: phrases [np_i [een scoutskameraad] werd .. en _i zocht hem op]
// idem: de hond was gebaseerd op Lassy en verscheen onder de naam Wirel nsubj:pass in conj1, nsubj in conj 2
func distributeDependents(node *nodeType, q *context) []depT {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "distributeDependents", q, node))
		}
	}()

	var phrase *nodeType
	if node.Rel == "hd" {
		if TEST(q, `$node/../../@cat="pp"`) { // door het schilderij
			phrase = node.parent.parent
		} else {
			phrase = node.parent
		}
	} else {
		if TEST(q, `$node[@rel="mwp" and @begin = ../@begin]`) {
			if TEST(q, `$node[ ../@rel="obj1" and ../../@cat="pp"]`) {
				phrase = node.parent.parent
			} else {
				if TEST(q, `$node[../@rel="hd" and ../../@rel="obj1" and ../../../@cat="pp"]`) {
					phrase = node.parent.parent.parent // in en rond het Hoofstedelijk Gewest --> do not distribute Hoofdstedelijk
				} else {
					if TEST(q, `$node[ ../@rel="hd" and not( ../../@cat="pp") ]`) { // mwu as head, but not complex P
						phrase = node.parent.parent
					} else {
						phrase = node.parent
					}
				}
			}
		} else {
			if TEST(q, `$node[@rel="obj1" and ../@cat="pp"]`) {
				phrase = node.parent
			} else {
				phrase = node
				// do not apply to prepositions and auxiliaries, ever. Too strict?
			}
		}
	}

	if !TEST(q, `$phrase[@rel=("obj1","su","mod","pc","det") and @index]`) {
		return []depT{}
	}

	// TODO: dit xpath kan efficiënter?
	conj_heads := FIND(q, `$node[not(@ud:pos=("ADP","AUX"))]/ancestor::node//node[@rel="cnj"
	   						     and node[
	   						    (: @rel=$phrase/@rel
	   							and -- this constraint is too strict for coord of passives:)
	   							not(@pt or @cat)]/@index = $phrase/@index
	   						     and node[@rel=("hd","predc") and not(@ud:pos="AUX") and (@pt or @cat) and	 (: bekende cafes zijn A en B :)
	   							(: not(@ud:pos=("ADP","AUX")) and not(@cat="mwu") :)
	   							not(local:internal_head_position(..) = @end and (@ud:pos=("ADP","AUX") or @cat="mwu") )
	   							]
	   					      ]
	   						      (: not coordination of AUX or (complex) Ps :)`)
	if len(conj_heads) == 0 {
		return []depT{}
	}

	//NP
	udRelation := nonLocalDependencyLabel(phrase, n1(FIND(q, `($q.varallnodes[@rel="cnj"]/
	   			    node[
	   			    (: @rel=$phrase/@rel and :)
					not(@pt or @cat) and @index=$phrase/@index])[1]`)), q)

	EudRelation := udRelation
	if TEST(q, `$udRelation = ("nmod","obl") and $phrase[@cat="pp"]//node[@ud:Relation="case" and @ud:HeadPosition=$node/@end]`) {
		//NP
		EudRelation = udRelation + ":" + enhancedLemmaString(FIND(q, `$phrase//node[@ud:Relation="case" and @ud:HeadPosition=$node/@end]`), q)
	}

	result := []depT{}
	for _, conj_head := range conj_heads {
		//NP
		result = append(result, depT{head: internalHeadPositionWithGapping([]interface{}{conj_head.(*nodeType)}, q), dep: EudRelation}) // added WithGapping to fix evalud/h_suite938, but does this work ??

	}
	return result
}

// note that result is a weird case [er/1 vlak/2 voor en 1  2 na] where er and vlak are not copied, but vlak does become mod of na in EUD now. Certainly not correct either.

// should work in coordinations like te laten reizen en te laten beleven,
// and recursive cases: Andras blijft ontkennen sexuele relaties met Timea te hebben gehad ,
//
//	.. of hij ook voor hen wilde komen tekenen :)
func xcompControl(node *nodeType, q *context, so_index int) []depT {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "xcompControl", q, node))
		}
	}()

	result := []depT{}
	for _, xcomp := range FIND(q, `$node[not(@ud:PronType="Rel")]/ancestor::node//node[(@rel="vc" or (@cat="inf" and @rel="body")) (: covers inf ti oti :)
	   					   and node[@rel=("hd","predc") and @ud:Relation="xcomp"]  (: vrouwen moeten vertegenwoordigd zijn :)
	   					   and node[@rel="su" and @index]/@index = $so_index
	   					  ]`) {
		//NP
		result = append(result, depT{head: internalHeadPosition([]interface{}{xcomp.(*nodeType)}, q), dep: "nsubj:xsubj"})
	}
	return result
}

// alpino NF specific case, controllers with extraposed content are realized downstairs
func upstairsControl(node *nodeType, q *context, so_index int) []depT {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "upstairsControl", q, node))
		}
	}()

	result := []depT{}
	for _, upstairs := range FIND(q, `$node/ancestor::node[node[@rel="hd" and @ud:pos="VERB"]
	   						  and node[@rel=("su","obj1","obj2") and not(@pt or @cat)]/@index = $so_index
	   						 ]`) {
		//NP
		result = append(result, depT{head: internalHeadPosition([]interface{}{upstairs.(*nodeType)}, q), dep: "nsubj:xsubj"})
	}
	return result

}

// een koers waarin de Alsemberg moet worden beklommen
func passiveVpControl(node *nodeType, q *context, so_index int) []depT {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "passiveVpControl", q, node))
		}
	}()

	result := []depT{}
	for _, passive_vp := range FIND(q, `$q.varallnodes[@rel="vc" and @cat="ppart"
	   				       and node[@rel="hd" and @ud:Relation="xcomp"]
	   				       and node[@rel="obj1" and @index]/@index = $so_index ]`) {
		//NP
		result = append(result, depT{head: internalHeadPosition([]interface{}{passive_vp.(*nodeType)}, q), dep: "nsubj:pass:xsubj"})
	}

	return result
}

func enhancedLemmaString(nodes []interface{}, q *context) string {
	sort.Slice(nodes, func(i, j int) bool {
		// TODO: nodeType heeft geen head
		return nodes[i].(*nodeType).udEHeadPosition < nodes[j].(*nodeType).udEHeadPosition
	})
	lemmas := make([]string, len(nodes))
	for i, node := range nodes {
		lemmas[i] = enhancedLemmaString1(node.(*nodeType), q)
	}
	return strings.Join(lemmas, "_")
}

func enhancedLemmaString1(node *nodeType, q *context) string {
	var lemma string
	switch node.Lemma {
	case "&":
		lemma = "en"
	case "+":
		lemma = "plus"
	case "a.k.a":
		lemma = "also_known_as"
	case "c.q.", "casu quo":
		lemma = "casu_quo"
	case "cum suis":
		lemma = "cum_suis"
	case "de dato":
		lemma = "dd"
	case "dwz.", "d.w.z.", "dat wil zeggen":
		lemma = "dat_wil_zeggen"
	case "e.a.", "en andere":
		lemma = "en_andere"
	case "e.d.", "en dergelijke":
		lemma = "en_dergelijk"
	case "en/of":
		lemma = "en_of"
	case "enz.", "enzovoorts":
		lemma = "enzovoort"
	case "etc.":
		lemma = "etcetera"
	case "zien":
		lemma = "gezien"
	case "inplaats":
		lemma = "in_plaats"
	case "in plaats van":
		lemma = "in_plaats_van"
	case "m.a.w.":
		lemma = "met_andere_woorden"
	case "nl.":
		lemma = "namelijk"
	case "resp.":
		lemma = "respectievelijk"
	case "t/m", "tot en met":
		lemma = "tot_en_met"
	case "t.a.v.":
		lemma = "te_aan_zien_van"
	case "t.g.v.":
		lemma = "ten_gunste_van"
	case "ten gevolge van":
		lemma = "tengevolge_van"
	case "t.n.v.":
		lemma = "ten_name_van"
	case "t.o.v.", "ten opzichte van":
		lemma = "te_opzicht_van"
	case "tusssen":
		lemma = "tussen"
	default:
		lemma = node.Lemma
	}
	fixed := FIND(q, `$node/../node[@ud:ERelation="fixed"]`)
	if len(fixed) > 0 {
		sort.Slice(fixed, func(i, j int) bool {
			fi := fixed[i].(*nodeType)
			fj := fixed[j].(*nodeType)
			ei := fi.End
			ej := fj.End
			if fi.udCopiedFrom > 0 {
				ei = fi.udCopiedFrom
			}
			if fj.udCopiedFrom > 0 {
				ej = fj.udCopiedFrom
			}
			return ei < ej
		})
		for _, f := range fixed {
			lemma += "_" + f.(*nodeType).Lemma
		}
	}
	lemma = strings.Replace(lemma, "/", "schuine_streep", -1) // this works for 't/m'
	lemma = strings.Replace(lemma, "-", "_", -1)
	lemma = strings.Replace(lemma, "en_schuine_streep_of", "en_of", -1) //this works for 'en / of' (multi-token)
	lemma = strings.Replace(lemma, "inplaats_van", "in_plaats_van", -1) // fixing typo in cdb/7002
	lemma = strings.Replace(lemma, "tot_schuine_streep_met", "tot_en_met", -1)
	lemma = strings.Replace(lemma, "te_gevolg_van", "tengevolge_van", -1)
	lemma = strings.Replace(lemma, "te_aanzien_van", "te_aan_zien_van", -1)
	lemma = strings.Replace(lemma, "te_gunst_van", "ten_gunste_van", -1)
	lemma = strings.Replace(lemma, "dat_willen_zeggen", "dat_wil_zeggen", -1)
	lemma = strings.Replace(lemma, "en andere", "en_andere", -1)
	lemma = strings.Replace(lemma, "met name", "met_name", -1)
	lemma = strings.Replace(lemma, " ", "_", -1) // default for expanded abbreviations GB 2/3/23
	return strings.ToLower(lemma)
}
