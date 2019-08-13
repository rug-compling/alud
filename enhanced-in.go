// +build ignore

package alud

import (
	"fmt"
	"sort"
	"strings"
)

type depType struct {
	head int
	dep  string
}

func enhancedDependencies(q *context) {

	changed := reconstructEmptyHead(q)

	// add_Edependency_relations
	for _, node := range q.ptnodes {
		// Edependency_relation
		if changed {
			q.depth = 0
			node.udERelation = dependencyLabel(node, q, []trace{trace{s: "enhancedDependencies", node: node}})
			q.depth = 0
			node.udEHeadPosition = externalHeadPosition(list(node), q, []trace{trace{s: "enhancedDependencies", node: node}})
			if node.udEHeadPosition == 0 && node.udERelation != "root" ||
				node.udEHeadPosition != 0 && node.udERelation == "root" {
				panic(fmt.Sprintf("Invalid DEPS %s:%s in %s:%s",
					number(node.udEHeadPosition), node.udERelation,
					number(node.End), node.Word))
			}
		} else {
			node.udERelation = node.udRelation
			node.udEHeadPosition = node.udHeadPosition
		}
	}

	for _, node := range q.ptnodes {
		enhancedDependencies1(node, q)
	}
}

func enhancedDependencies1(node *nodeType, q *context) {
	// iobj2 control : de commissie kan de raad aanbevelen/adviseren/ X te doen
	// rhd : een levend visje dat doorgeslikt moet worden
	q.depth = 0
	var enhanced []depType
	tr := []trace{trace{s: "enhancedDependencies1", node: node}}
	for {

		if TEST(q, `$node[@ud:ERelation=("nsubj","obj","iobj","nsubj:pass")]`) { // TODO: klopt dit? exists binnen [ ]
			so := FIND(q,
				`$node/ancestor::node/node[@rel=("su","obj1","obj2") and local:internal_head_position(.) = $node/@end ]/@index`)
			if len(so) > 0 {
				soIndex := i1(so)
				enhanced = []depType{depType{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q, tr)}} // self
				enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)                                        // self
				enhanced = append(enhanced, distributeConjuncts(node, q, tr)...)                                    // self
				enhanced = append(enhanced, distributeDependents(node, q, tr)...)                                   // self
				enhanced = append(enhanced, xcompControl(node, q, tr, soIndex)...)
				enhanced = append(enhanced, upstairsControl(node, q, tr, soIndex)...)
				enhanced = append(enhanced, passiveVpControl(node, q, tr, soIndex)...)
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
				enhanced = []depType{depType{head: internalHeadPositionWithGapping(rhdNp, q, tr), dep: "ref"}} // rhdref
				enhanced = append(enhanced, xcompControl(node, q, tr, rhdIndex)...)
				enhanced = append(enhanced, passiveVpControl(node, q, tr, rhdIndex)...)
				break
			}
			// if there is no antecedent, lets keep the basic relation
			enhanced = []depType{depType{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q, tr)}} // self
			enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)                                        // self
			enhanced = append(enhanced, distributeConjuncts(node, q, tr)...)                                    // self
			enhanced = append(enhanced, distributeDependents(node, q, tr)...)                                   // self
			enhanced = append(enhanced, xcompControl(node, q, tr, rhdIndex)...)
			enhanced = append(enhanced, passiveVpControl(node, q, tr, rhdIndex)...)
			break
		}

		relSister := FIND(q, `($node/../node[@rel="mod" and @cat="rel"]/node[@rel="rhd"]/@index)[1]`)
		if len(relSister) > 0 {
			relSisterIndex := i1(relSister)
			enhanced = []depType{depType{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q, tr)}} // self
			enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)                                        // self
			enhanced = append(enhanced, distributeConjuncts(node, q, tr)...)                                    // self
			enhanced = append(enhanced, distributeDependents(node, q, tr)...)                                   // self
			enhanced = append(enhanced, xcompControl(node, q, tr, relSisterIndex)...)
			enhanced = append(enhanced, passiveVpControl(node, q, tr, relSisterIndex)...)
			break
		}

		// underscore is resultaat van reconstructEmptyHead()
		if node.udHeadPosition >= 0 || node.udHeadPosition == underscore {
			enhanced = []depType{depType{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q, tr)}} // self
			enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)                                        // self
			enhanced = append(enhanced, distributeConjuncts(node, q, tr)...)                                    // self
			enhanced = append(enhanced, distributeDependents(node, q, tr)...)                                   // self
			break
		}

		enhanced = []depType{depType{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q, tr)}}
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
		if e.head == 0 && e.dep != "root" {
			panic(fmt.Sprintf("Invalid EUD 0:%s in %s:%s", e.dep, number(node.End), node.Word))
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

func enhanceDependencyLabel(node *nodeType, q *context, tr []trace) string {
	tr = append(tr, trace{s: "enhanceDependencyLabel", node: node})
	label := node.udERelation
	if label == "conj" {
		if crd := n1(FIND(q, `($node/ancestor::node[@cat="conj" and
	       not(.//node[@cat="conj"]//node/@begin = $node/@begin)]/node[@rel="crd"])[1]`)); crd != noNode {
			if crd.Lemma != "" {
				return join(label, enhancedLemmaString1(crd, q))
			}
			if crd.Cat == "mwu" {
				return join(label, enhancedLemmaString1(n1(FIND(q, `($crd/node[@rel="mwp"])[1]`)), q))
			}
			panic(tracer("Empty EUD label", tr, q))
		}
	}

	if label == "nmod" || label == "obl" {
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

	panic(tracer("Empty EUD label", tr, q))
}

func anaphoricRelpronoun(node *nodeType, q *context) []depType {
	// works voor waar, and last() picks waar in 'daar waar' cases
	// dont add anything for hij werd voorzitter, wat hij nog steeds is (otherwise self-reference)
	// for loop ensures correct result if N has 2 acl:relcl dependents
	result := []depType{}
	for _, a := range FIND(q, `$node/ancestor::node[@cat="np" and local:internal_head_position(.) = $node/@end]/
	   			       node[@rel="mod"]/node[@rel="rhd"]/descendant-or-self::node[@pt="vnw" and not(@ud:HeadPosition = $node/@end)][last()]`) {
		anrel := a.(*nodeType)
		var label string
		if r := anrel.udRelation; r == "nsubj" || r == "nsubj:pass" {
			label = r + ":relsubj"
		} else if r == "obj" || r == "obl" {
			label = r + ":relobj"
		} else {
			label = r
		}
		result = append(result, depType{head: anrel.udHeadPosition, dep: label})
	}
	return result
}

// Glastra en Terlouw verzonnen een list --> nsubj(verzonnen,Glastra) nsubj(verzonnen,Terlouw)
func distributeConjuncts(node *nodeType, q *context, tr []trace) []depType {
	tr = append(tr, trace{s: "distributeConjuncts", node: node})
	if node.udRelation == "conj" {
		coordHead := n1(FIND(q, `$q.varallnodes[@end = $node/@ud:HeadPosition
	       and @ud:Relation=("amod","appos","nmod","nsubj","nsubj:pass","nummod","obj","iobj","obl","obl:agent","advcl")]`))
		if coordHead != noNode {
			// in A en B vs in A en naast B --> use enh_dep_label($node) in the latter case...
			depLabel := enhanceDependencyLabel(coordHead, q, tr)
			return []depType{depType{head: coordHead.udHeadPosition, dep: depLabel}}
		}
	}
	return []depType{}
}

// de onrust kan een reis vertragen of frustreren  --> obj(vertragen,reis) obj(frustreren,reis)
// todo: passives ze werd ontmanteld en verkocht  su coindexed with two obj1
// done: phrases [np_i [een scoutskameraad] werd .. en _i zocht hem op]
// idem: de hond was gebaseerd op Lassy en verscheen onder de naam Wirel nsubj:pass in conj1, nsubj in conj 2
func distributeDependents(node *nodeType, q *context, tr []trace) []depType {
	tr = append(tr, trace{s: "distributeDependents", node: node})
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
		return []depType{}
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
		return []depType{}
	}

	udRelation := nonLocalDependencyLabel(phrase, n1(FIND(q, `($q.varallnodes[@rel="cnj"]/
	   			    node[
	   			    (: @rel=$phrase/@rel and :)
					not(@pt or @cat) and @index=$phrase/@index])[1]`)), q, tr)

	EudRelation := udRelation
	if TEST(q, `$udRelation = ("nmod","obl") and $phrase[@cat="pp"]//node[@ud:Relation="case" and @ud:HeadPosition=$node/@end]`) {
		EudRelation = udRelation + ":" + enhancedLemmaString(FIND(q, `$phrase//node[@ud:Relation="case" and @ud:HeadPosition=$node/@end]`), q)
	}

	result := []depType{}
	for _, conj_head := range conj_heads {
		result = append(result, depType{head: internalHeadPosition([]interface{}{conj_head.(*nodeType)}, q, tr), dep: EudRelation})

	}
	return result
}

// should work in coordinations like te laten reizen en te laten beleven,
// and recursive cases: Andras blijft ontkennen sexuele relaties met Timea te hebben gehad ,
//    .. of hij ook voor hen wilde komen tekenen :)
func xcompControl(node *nodeType, q *context, tr []trace, so_index int) []depType {

	tr = append(tr, trace{s: "xcompControl", node: node})

	result := []depType{}
	for _, xcomp := range FIND(q, `$node[not(@ud:PronType="Rel")]/ancestor::node//node[(@rel="vc" or (@cat="inf" and @rel="body")) (: covers inf ti oti :)
	   					   and node[@rel=("hd","predc") and @ud:Relation="xcomp"]  (: vrouwen moeten vertegenwoordigd zijn :)
	   					   and node[@rel="su" and @index]/@index = $so_index
	   					  ]`) {
		result = append(result, depType{head: internalHeadPosition([]interface{}{xcomp.(*nodeType)}, q, tr), dep: "nsubj:xsubj"})
	}
	return result
}

// alpino NF specific case, controllers with extraposed content are realized downstairs
func upstairsControl(node *nodeType, q *context, tr []trace, so_index int) []depType {

	tr = append(tr, trace{s: "upstairsControl", node: node})

	result := []depType{}
	for _, upstairs := range FIND(q, `$node/ancestor::node[node[@rel="hd" and @ud:pos="VERB"]
	   						  and node[@rel=("su","obj1","obj2") and not(@pt or @cat)]/@index = $so_index
	   						 ]`) {
		result = append(result, depType{head: internalHeadPosition([]interface{}{upstairs.(*nodeType)}, q, tr), dep: "nsubj:xsubj"})
	}
	return result

}

// een koers waarin de Alsemberg moet worden beklommen
func passiveVpControl(node *nodeType, q *context, tr []trace, so_index int) []depType {

	tr = append(tr, trace{s: "passiveVpControl", node: node})

	result := []depType{}
	for _, passive_vp := range FIND(q, `$q.varallnodes[@rel="vc" and @cat="ppart"
	   				       and node[@rel="hd" and @ud:Relation="xcomp"]
	   				       and node[@rel="obj1" and @index]/@index = $so_index ]`) {
		result = append(result, depType{head: internalHeadPosition([]interface{}{passive_vp.(*nodeType)}, q, tr), dep: "nsubj:pass:xsubj"})
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
	lemma = strings.Replace(lemma, "/", "schuine_streep", -1)
	lemma = strings.Replace(lemma, "-", "_", -1)
	return strings.ToLower(lemma)
}
