package main

import (
	//	"github.com/kr/pretty"

	//	"fmt"
	"sort"
)

// voorkwam dat LPF opnieuw of SGP voor het eerst in de regering zou komen  -- gapped LD
func fixMisplacedHeadsInCoordination(q *Context) {

	if len(q.varindexnodes) == 0 {
		return
	}

START:
	for true {
		for _, n1 := range q.varallnodes {
			// FIND op varallnodes niet mogelijk omdat twee keer naar $node wordt verwezen, en dat moet dezelfde node zijn
			for _, n2 := range FIND(n1, q, `
$node[@rel=("hd","ld") and
      @index and
      (@pt or @cat) and
      ancestor::node[@rel="cnj"] and
      ancestor::node[@cat="conj"]/node[@rel="cnj" and
                                       descendant-or-self::node[@rel=("hd","ld") and
                                                                @index=$node/@index and
                                                                not(@cat or @pt) and
                                                                ( @begin        = ..//node[@cat or @pt]/@end or
                                                                  @begin - 1000 = ..//node[@cat or @pt]/@end
                                                                )
                                                               ]
                                       ]]`) {
				node2 := n2.(*NodeType)
				for _, n3 := range FIND(q.varallnodes, q, `
$node[@rel=("hd","ld","vc") and @index and not(@pt or @cat) and
                 ancestor::node[@rel="cnj"]  and
                                    ( @begin        = ..//node[@cat or @pt]/@end or
                                      @begin - 1000 = ..//node[@cat or @pt]/@end
                                     )]`) {
					node3 := n3.(*NodeType)
					if node2.Index == node3.Index {
						// kopieer inhoud van node2 (niet leeg) naar node3 (leeg)
						id, rel := node3.Id, node3.Rel
						*node3 = *node2
						node3.Id, node3.Rel = id, rel
						// maak node2 leeg
						*node2 = NodeType{
							Begin: node2.Begin,
							End:   node2.End,
							Id:    node2.Id,
							Index: node2.Index,
							Rel:   node2.Rel,
							Node:  []*NodeType{},
						}
						// opnieuw beginnen
						inspect(q)
						continue START
					}
				}
			}
		}
		break
	}
}

func addPosTags(q *Context) {
	for _, node := range q.ptnodes {
		node.udPos = universalPosTags(node, q)
	}
}

func universalPosTags(node *NodeType, q *Context) string {
	pt := node.Pt
	rel := node.Rel

	if pt == "let" {
		if rel == "--" {
			for _, n := range node.parent.Node {
				if n.Pt != "let" || n.Begin < node.Begin {
					return "PUNCT"
				}
			}
		}
		return "SYM"
	}
	if pt == "adj" {
		if rel == "det" {
			return "DET"
		}
		if rel == "hd" && node.parent.Cat == "pp" {
			// vol vertrouwen
			return "ADP"
		}
		if rel == "crd" {
			// respectievelijk
			return "CCONJ"
		}
		return "ADJ" // exceptions forced by 2.4 validation
	}
	if pt == "bw" {
		return "ADV"
	}
	if pt == "lid" {
		return "DET"
	}
	if pt == "n" {
		if node.Ntype == "eigen" {
			return "PROPN"
		}
		return "NOUN"
	}
	if pt == "spec" {
		if node.Spectype == "deeleigen" {
			return "PROPN"
		}
		if node.Spectype == "symb" {
			return "SYM"
		}
		return "X" // afk vreemd afgebr enof meta
	}
	if pt == "tsw" {
		return "INTJ"
	}
	if pt == "tw" {
		if node.Numtype == "rang" {
			return "ADJ"
		}
		return "NUM"
	}
	if pt == "vz" {
		return "ADP" // v2: do not use PART for SVPs and complementizers
	}
	if pt == "vnw" {
		if rel == "det" && node.Vwtype != "bez" {
			return "DET"
		}
		if node.Pdtype == "adv-pron" {
			if rel == "pobj1" {
				return "PRON"
			}
			return "ADV"
		}
		if (rel == "mod" || (rel == "hd" && node.parent.Rel == "mod")) && node.Pdtype == "grad" {
			// veel minder
			return "ADV"
		}
		return "PRON"
	}
	if pt == "vg" {
		if node.Conjtype == "neven" {
			return "CCONJ" // V2: CONJ ==> CCONJ
		}
		return "SCONJ"
	}
	if pt == "ww" {
		if auxiliary1(node, q) == "verb" {
			return "VERB"
		}
		return "AUX" // v2: cop and aux:pass --> AUX  (already in place in v1?)
	}
	return "ERROR_NO_POS_FOUND"
}

func addFeatures(q *Context) {
	for _, node := range q.ptnodes {
		switch node.udPos {
		case "NOUN", "PROPN":
			nominalFeatues(node, q)
		case "ADJ":
			adjectiveFeatures(node, q)
		case "PRON":
			pronounFeatures(node, q)
		case "VERB", "AUX":
			verbalFeatures(node, q)
		case "DET":
			determinerFeatures(node, q)
		case "X":
			specialFeatures(node, q)
		}
	}
}

func nominalFeatues(node *NodeType, q *Context) {
	switch node.Genus {
	case "zijd":
		node.udGender = "Com"
	case "onz":
		node.udGender = "Neut"
	case "genus":
		node.udGender = "Com,Neut"
	case "":
		node.udGender = ""
	default:
		node.udGender = "ERROR_IRREGULAR_GENDER"

	}
	switch node.Getal {
	case "ev":
		node.udNumber = "Sing"
	case "mv":
		node.udNumber = "Plur"
	case "":
		node.udNumber = ""
	default:
		node.udNumber = "ERROR_IRREGULAR_NUMBER"
	}
	switch node.Graad {
	case "dim":
		node.udDegree = ""
	case "basis":
		node.udDegree = ""
	case "":
		node.udDegree = ""
	default:
		node.udDegree = "ERROR_IRREGULAR_DEGREE"
	}
}

func adjectiveFeatures(node *NodeType, q *Context) {
	switch node.Graad {
	case "basis":
		node.udDegree = "Pos"
	case "comp":
		node.udDegree = "Cmp"
	case "sup":
		node.udDegree = "Sup"
	case "dim": // netjes
		node.udDegree = "Pos"
	case "":
		node.udDegree = ""
	default:
		node.udDegree = "ERROR_IRREGULAR_DEGREE"
	}
}

func pronounFeatures(node *NodeType, q *Context) {
	switch node.Vwtype {
	case "refl":
		node.udPronType = "Prs"
		node.udReflex = "Yes"
	case "bez":
		node.udPronType = "Prs"
		node.udPoss = "Yes"
	case "pers", "pr":
		node.udPronType = "Prs"
	case "recip":
		node.udPronType = "Rcp"
	case "vb":
		node.udPronType = "Int"
	case "aanw":
		node.udPronType = "Dem"
	case "onbep":
		node.udPronType = "Ind"
	case "betr":
		node.udPronType = "Rel"
	case "excl": // occurs only once
		node.udPronType = ""
	case "":
		node.udPronType = ""
	default:
		node.udPronType = "ERROR_IRREGULAR_PRONTYPE"
	}

	switch node.Persoon {
	case "1":
		node.udPerson = "1"
	case "2", "2b", "2v":
		node.udPerson = "2"
	case "3", "3o", "3v", "3p", "3m":
		node.udPerson = "3"
	case "persoon":
		node.udPerson = ""
	case "":
		node.udPerson = ""
	default:
		node.udPerson = "ERROR_IRREGULAR_PERSON"
	}

	switch node.Naamval {
	case "nomin":
		node.udCase = "Nom"
	case "obl":
		node.udCase = "Acc"
	case "gen":
		node.udCase = "Gen"
	case "dat": // van dien aard
		node.udCase = "Dat"
	case "stan":
		node.udCase = ""
	case "":
		node.udCase = ""
	default:
		node.udCase = "ERROR_IRREGULAR_CASE"

	}
}

func verbalFeatures(node *NodeType, q *Context) {

	switch node.Wvorm {
	case "pv":
		node.udVerbForm = "Fin"
	case "inf":
		node.udVerbForm = "Inf"
	case "vd", "od":
		node.udVerbForm = "Part"
	case "":
		node.udVerbForm = ""
	default:
		node.udVerbForm = "ERROR_IRREGULAR_VERBFORM"
	}

	switch node.Pvtijd {
	case "verl":
		node.udTense = "Past"
	case "tgw", "conj":
		node.udTense = "Pres"
	case "":
		node.udTense = ""
	default:
		node.udTense = "ERROR_IRREGULAR_TENSE"
	}

	switch node.Pvagr {
	case "ev", "met-t":
		node.udNumber = "Sing"
	case "mv":
		node.udNumber = "Plur"
	case "":
		node.udNumber = ""
	default:
		node.udNumber = "ERROR_IRREGULAR_NUMBER"
	}
}

func determinerFeatures(node *NodeType, q *Context) {
	switch node.Lwtype {
	case "bep":
		node.udDefinite = "Def"
	case "onbep":
		node.udDefinite = "Ind"
	case "":
		node.udDefinite = ""
	default:
		node.udDefinite = "ERROR_IRREGULAR_DEFINITE"
	}
}

func specialFeatures(node *NodeType, q *Context) {
	switch node.Spectype {
	case "vreemd":
		node.udForeign = "Yes"
	case "afk":
		node.udAbbr = "Yes"
	}
}

func addDependencyRelations(q *Context) {
	for _, node := range q.ptnodes {
		node.udRelation = dependencyLabel(node, q)
		node.udHeadPosition = externalHeadPosition(node, q)
	}
}

func dependencyLabel(node *NodeType, q *Context) string {
	if node.parent.Cat == "top" && node.parent.End == 1000 {
		return "root"
	}
	if node.Rel == "app" {
		if TEST(node, q, `$node/../node[@rel="hd" and (@pt or @cat)]`) {
			return "appos"
		}
		if TEST(node, q, `$node/../node[@rel="mod" and (@pt or @cat)]`) {
			return "orphan"
		}
		return dependencyLabel(node.parent, q)
	}
	if node.Rel == "cmp" {
		return "mark"
	}
	if node.Rel == "crd" {
		return "cc"
	}
	if node.Rel == "me" && TEST(node, q, `$node[@rel="me" and not(../node[@ud:pos="ADP"]) ]`) {
		return determineNominalModLabel(node, q)
	}
	if node.Rel == "obcomp" {
		return "advcl"
	}
	if node.Rel == "obj2" {
		if node.Cat == "pp" {
			return "obl"
		}
		return "iobj"
	}
	if node.Rel == "pobj1" {
		return "expl"
	}
	if node.Rel == "predc" {
		if TEST(node, q, `$node/../node[@rel=("obj1","se") and (@pt or @cat)] or $node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX")]`) {
			if TEST(node, q, `$node/../@cat="pp"`) { // check for absolutive (met) constructions, https://github.com/UniversalDependencies/docs/issues/408
				if TEST(node, q, `$node/../../@cat="np"`) {
					return "acl"
				}
				return "advcl"
			}
			return "xcomp"
		}
		return dependencyLabel(node.parent, q) // covers gapping cases where predc is promoted to head as well
		/*
		   hack for now: de keuze is gauw gemaakt
		   was amod, is this more accurate??
		   examples of secondary predicates under xcomp suggests so

		*/
	}
	if node.Rel == "se" {
		return "expl:pv"
	}
	if node.Rel == "su" {
		if TEST(node, q, `$node[../@rel="cnj" and ../node[@rel="hd" and not(@pt or @cat)]]`) { // gapping
			return dependencyLabel(node.parent, q)
		}
		if TEST(node, q, `$node[../@rel="vc" and ../node[@rel="hd" and not(@pt or @cat)]
	                                 and ../../self::node[@rel="cnj" and node[@rel="hd" and not(@pt or @cat)]]]`) { // gapping with subj downstairs
			// TODO: ../.. is veranderd in ../../self::node    is dat juist?
			/*
			   In 1909 werd de persoonlijke dienstplicht ingevoerd en in 1913 de algemene persoonlijke dienstplicht .
			   [ hd_i su_j vc [ hd_k [_j pers dienstplicht ]
			*/
			return dependencyLabel(node.parent.parent, q)
		}
		return subjectLabel(node, q)
	}
	if node.Rel == "sup" {
		return "expl"
	}
	if node.Rel == "svp" {
		return "compound:prt" // v2: added prt extension
	}
	aux := auxiliary1(node, q)
	if aux == "aux:pass" {
		if TEST(node, q, `$node[../node[@rel="su" and not(@pt or @cat)] and
	                 ../node[@rel="vc" and not(@pt or @cat)] and
	                 ../@rel="cnj"]`) {
			return "conj"
		}
		return "aux:pass"
	}
	if aux == "aux" {
		return "aux"
	}
	if aux == "cop" {
		return "cop"
	}
	if node.Rel == "det" {
		if TEST(node, q, `$node/../node[@rel="hd" and (@pt or @cat)]`) {
			return detLabel(node, q)
		}
		if TEST(node, q, `$node/../node[@rel="mod" and (@pt or @cat)]`) { // gapping
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping
	}
	if node.Rel == "obj1" || node.Rel == "me" {
		if TEST(node, q, `$node/../@cat="pp" or $node/../node[@rel="hd" and @ud:pos="ADP"]`) { // vol vertrouwen , heel de geschiedenis door (cat=ap!)
			if TEST(node, q, `$node/../node[@rel="predc"]`) { // absolutive met
				return "nsubj"
			}
			return dependencyLabel(node.parent, q)
		}
		if TEST(node, q, `$node[@index = ../../node[@rel="su"]/@index ]`) {
			return "nsubj:pass" // trees where su (with extraposed material) is spelled out at position of obj1
		}
		if TEST(node, q, `$node/../node[@rel="hd" and (@pt or @cat)]`) {
			return "obj"
		}
		if TEST(node, q, `$node/../node[@rel="su" and (@pt or @cat)]`) {
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping
	}
	if node.Rel == "mwp" {
		if node.Begin >= 0 && node.Begin == node.parent.Begin {
			return dependencyLabel(node.parent, q)
		}
		if TEST(node, q, `$node/../node[@ud:pos="PROPN"]`) {
			return "flat"
		}
		return "fixed" // v2 mwe-> fixed
	}
	if node.Rel == "cnj" {
		// TODO als ik hier firstnode gebruik i.p.v. leftmost (wat zou moeten) gaat het vaak mis
		if node == leftmost(FIND(node, q, `$node/../node[@rel="cnj"]`)) {
			return dependencyLabel(node.parent, q)
		}
		return "conj"
	}
	if node.Rel == "dp" {
		if node == leftmost(FIND(node, q, `$node/../node[@rel="dp"]`)) {
			return dependencyLabel(node.parent, q)
		}
		return "parataxis"
	}
	if node.Rel == "tag" || node.Rel == "sat" {
		return "parataxis"
	}
	if node.Rel == "dlink" {
		return "mark"
	}
	if node.Rel == "nucl" {
		return dependencyLabel(node.parent, q)
	}
	if node.Rel == "vc" {
		if TEST(node, q, `$node/../node[@rel="hd" and @ud:pos=("AUX","ADP")] and not($node/../node[@rel="predc"])`) {
			return dependencyLabel(node.parent, q)
		}
		if TEST(node, q, `$node/../node[@rel="hd" and (@pt or @cat)]`) {
			if TEST(node, q, `$node/node[@rel="su" and @index and not(@word or @cat)]
	                           (: or $node[@cat=("ti","oti")] :)
	                           or $node[@cat="ti"]/node[@rel="body"]/node[@rel="su" and @index and not(@word or @cat)]
	                           or $node[@cat="oti"]/node[@cat="ti"]/node[@rel="body"]/node[@rel="su" and @index and not(@word or @cat)]`) {
				return "xcomp"
			}
			if TEST(node, q, `$node/../@cat="np"`) {
				return "acl" // v2: clausal dependents of nouns always acl
			}
			return "ccomp"
		}
		if TEST(node, q, `$node/../node[@rel=("su","obj1") and (@pt or @cat)]`) {
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping
	}
	if (node.Rel == "mod" || node.Rel == "pc" || node.Rel == "ld") && node.parent.Cat == "np" { // [detp niet veel] meer
		// modification of nomimal heads
		// pc and ld occur in nominalizations
		if TEST(node, q, `$node/../node[@rel="hd" and (@pt or @cat)]`) {
			return modLabelInsideNp(node, q)
		}
		if node == leftmost(FIND(node, q, `$node/../node[@rel="mod" and (@pt or @cat)]`)) { // gapping with multiple mods
			return dependencyLabel(node.parent, q) // gapping, where this mod is the head
		}
		return "orphan"
	}
	if TEST(node, q, `$node[@rel=("mod","pc","ld") and ../@cat=("sv1","smain","ssub","inf","ppres","ppart","oti","ap","advp")]`) {
		// modification of verbal, adjectival heads
		// nb some oti's directly dominate (preceding) modifiers
		// [advp weg ermee ]
		if TEST(node, q, `$node/../node[@rel=("hd","body") and (@pt or @cat)]`) { // body for mods dangling outside cmp/body: maar niet om ...
			return labelVmod(node, q)
		}
		if TEST(node, q, `$node/../node[@rel=("su","obj1","predc","vc") and (@pt or @cat)]`) {
			return "orphan"
		}
		if TEST(node, q, `$node[@rel="mod" and ../node[@rel=("pc","ld")]]`) {
			return "orphan"
		}
		if TEST(node, q, `$node[@rel="mod" and ../node[@rel="mod"]/@begin < @begin]`) { // gapping with multiple mods
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping, where this mod is the head
	}
	if TEST(node, q, `$node[@rel="mod" and ../@cat=("pp","detp","advp")]`) {
		return "amod"
	}
	if TEST(node, q, `$node[@rel="mod" and ../@cat=("cp", "whrel", "whq", "whsub")]`) {
		// [cp net  [cmp als] [body de Belgische regering]], net wat het land nodig heeft
		if TEST(node, q, `$node/../node[@rel="body"]/node[@rel="hd" and @ud:pos="VERB"]`) {
			return "advmod"
		}
		return "amod"
	}
	if node.Rel == "pc" && node.parent.Cat == "pp" { //  [[hd,mwu aan het hoofd] [pc van X]]
		return "nmod"
	}
	if node.Rel == "hdf" {
		return "case"
	}
	if node.Rel == "predm" {
		if node.udPos == "VERB" {
			return "xcomp"
		}
		if node.udPos != "" {
			return "advmod"
		}
		return "advcl"
	}
	if node.Rel == "rhd" || node.Rel == "whd" {
		if TEST(node, q, `$node/../node[@rel="body"]//node/@index = $node/@index`) {
			// index is een int groter dan 0
			return nonLocalDependencyLabel(
				node,
				firstnode(FIND(node, q, `$node/../node[@rel="body"]//node[@index = $node/@index]`)),
				q)
		}
		if node.Cat == "pp" {
			return "nmod" // onder wie michael boogerd
		}
		return "advmod" // [whq waarom jij]
	}
	if node.Rel == "body" {
		return dependencyLabel(node.parent, q)
	}
	if node.Rel == "--" {
		if node.udPos == "PUNCT" {
			if TEST(node, q, `$node[not(../node[not(@ud:pos="PUNCT")]) and @begin = ../@begin]`) {
				return "root" // just punctuation
			}
			return "punct"
		}
		if node.udPos == "SYM" || node.udPos == "X" {
			if TEST(node, q, `$node/../node[@cat]`) {
				return "appos" // 1. Jantje is ziek  1-->appos??
			}
			return "root"
		}
		if node.Lemma == "\\" {
			return "punct" // hack for tagging errors in lassy small 250
		}
		/*
			if node.Spectype == "deeleigen" {
				return "punct" // hack for tagging errors in lassy small 250
			}
		*/
		if TEST(node, q, `$node[@ud:pos="NUM" and ../node[@cat] ]`) {
			return "parataxis" // dangling number 1.
		}
		if TEST(node, q, `$node[@ud:pos="CCONJ" and ../node[@cat="smain" or @cat="conj"]]`) {
			return "cc"
		}
		// sentence initial or final 'en'
		if TEST(node, q, `$node[@ud:pos=("NOUN","PROPN","VERB") and ../node[@cat=("du","smain")]]`) {
			return "parataxis" // dangling words
		}
		if len(FIND(node, q, `$node/../node[not(@ud:pos=("PUNCT","SYM","X"))]`)) < 2 {
			return "root" // only one non-punct/sym/foreign element in the string
		}
		if node.Cat == "mwu" {
			if node.Begin == node.parent.Begin && node.End == node.parent.End {
				return "root"
			}
			if TEST(node, q, `$node/node[@ud:pos=("PUNCT","SYM")]`) { // fix for mwu punctuation in Alpino output
				return "punct"
			}
			return "ERROR_NO_LABEL_--"
		}
		if TEST(node, q, `$node[not(@ud:pos)]/../@rel="top"`) {
			return "root"
		}
		if TEST(node, q, `$node[@ud:pos="PROPN" and not(../node[@cat]) ]`) {
			return "root" // Arthur .
		}
		if TEST(node, q, `$node[@ud:pos=("ADP","ADV","ADJ","DET","PRON","CCONJ","NOUN","VERB","INTJ")]`) {
			return "parataxis"
		}
		return "ERROR_NO_LABEL_--"
	}
	if node.Rel == "hd" {
		if node.udPos == "ADP" {
			if TEST(node, q, `$node/../node[@rel="predc"]`) {
				return "mark" // absolute met constructie -- analoog aan with X being Y
			}
			if TEST(node, q, `not($node/../node[@rel="pc"])`) {
				return "case" // er blijft weinig over van het lijk : over heads a predc and has pc as sister
			}
			return dependencyLabel(node.parent, q) // not sure about this one
		}
		if TEST(node, q, `$node[(@ud:pos=("ADJ","X","ADV") or @cat="mwu")
	                            and ../@cat="pp"
	                            and ../node[@rel=("obj1","se","vc")]]`) {
			if TEST(node, q, `$node[../@rel="cnj" and ../node[@rel="obj1" and @index and not(@pt or @cat)]]`) {
				return "conj"
			}
			return "case" // vol vertrouwen in, Ad U3... , op het gebied van
		}
		if TEST(node, q, `$node/../node[@rel="hd"]/@begin < $node/@begin`) {
			return "conj"
		}
		return dependencyLabel(node.parent, q)
	}
	return "ERROR_NO_LABEL"
}

func determineNominalModLabel(node *NodeType, q *Context) string {
	if TEST(node, q, `$node/../node[@rel="hd" and (@ud:pos="VERB" or @ud:pos="ADJ")]`) {
		return "obl"
	}
	return "nmod"
}

func subjectLabel(subj *NodeType, q *Context) string {
	pass := passiveSubject(subj, q)
	if SUTEST(subj, q, `$subj[@cat=("whsub","ssub","ti","cp","oti")] or
	            $subj[@cat="conj" and node/@cat=("whsub","ssub","ti","cp","oti")]`) {
		return "csubj" + pass
	}
	// weather verbs and other expletive subject verbs
	if SUTEST(subj, q, `$subj[@lemma="het"] and
	                   ( $subj/../node[@rel="hd" and
	                                   @lemma=("dooien", "gieten", "hagelen", "miezeren",
	                                           "misten", "motregenen", "onweren", "plenzen",
	                                           "regenen", "sneeuwen", "stormen", "stortregenen",
	                                           "ijzelen", "vriezen", "weerlichten", "winteren",
	                                           "zomeren") ] or
	                     $subj/../node[@rel="hd" and
	                                   @lemma="ontbreken" and
	                                   ../node[@rel="pc" and node[@rel="hd" and @lemma="aan"] ] ] or
	                     $subj[@index = ../node//node[@rel="sup"]/@index]    (: het kan voorkomen dat ... :)
	                   )`) {
		return "expl" + pass
	}
	return "nsubj" + pass
}

func passiveSubject(subj *NodeType, q *Context) string {
	aux := auxiliary1(firstnode(SUFIND(subj, q, `$subj/../node[@rel="hd"]`)), q)
	if aux == "aux:pass" { // de carriere had gered kunnen worden
		return ":pass"
	}
	if aux == "aux" && SUTEST(subj, q, `$subj/@index = $subj/../node[@rel="vc"]/node[@rel="su"]/@index`) {
		return passiveSubject(firstnode(SUFIND(subj, q, `$subj/../node[@rel="vc"]/node[@rel="su"]`)), q)
	}
	return ""
}

func detLabel(node *NodeType, q *Context) string {
	// zijn boek, diens boek, ieders boek, aller landen, Ron's probleem, Fidel Castro's belang
	if TEST(node, q, `$node[@ud:pos = "PRON" and @vwtype="bez"] or
	          $node[@ud:pos = ("PRON","PROPN") and @naamval="gen"] or
	          $node[@cat="mwu" and node[@spectype="deeleigen"]]`) {
		return "nmod:poss"
	}
	if TEST(node, q, `$node/@ud:pos = ("DET","PROPN","NOUN","ADJ","PRON","ADV","X")`) {
		return "det" // meer  genoeg the
	}
	if TEST(node, q, `$node/@cat = ("mwu","np","pp","ap","detp","smain")`) {
		return "det"
	}
	// tussen 5 en 6 .., needs more principled solution
	// nog meer mensen dan anders
	// hetzelfde tijdstip als anders , nogal wat,
	// hij heeft ik weet niet hoeveel boeken
	if node.udPos == "NUM" || node.udPos == "SYM" {
		return "nummod"
	}
	if node.Cat == "conj" {
		if TEST(node, q, `$node/node[@rel="cnj"][1]/@ud:pos="NUM"`) {
			return "nummod"
		}
		return "det"
	}
	return "ERROR_NO_LABEL_DET"
}

func modLabelInsideNp(node *NodeType, q *Context) string {
	if TEST(node, q, `$node[@cat="pp"]/node[@rel="vc"]`) {
		return "acl" // pp containing a clause
	}
	if TEST(node, q, `$node[@ud:pos="ADJ" or @cat="ap" or node[@cat="conj" and node[@ud:pos="ADJ" or @cat="ap"] ]]`) {
		return "amod"
	}
	if TEST(node, q, `$node[@cat=("pp","np","conj","mwu") or @ud:pos=("NOUN","PRON","PROPN","X","PUNCT","SYM","INTJ") ]`) {
		return "nmod"
	}
	if node.udPos == "NUM" {
		return "nummod"
	}
	if TEST(node, q, `$node[@cat="detp"]/node[@rel="hd" and @ud:pos="NUM"]`) {
		return "nummod"
	}
	if node.Cat == "detp" {
		return "det" // [detp niet veel] meer error?
	}
	if node.Cat == "rel" || node.Cat == "whrel" {
		return "acl:relcl"
	}
	// v2 added relcl -- whrel= met name waar ...
	if TEST(node, q, `$node[@cat="cp"]/node[@rel="body" and (@ud:pos = ("NOUN","PROPN") or @cat=("np","conj"))]`) {
		return "nmod"
	}
	// zijn loopbaan [CP als schrijver]
	if TEST(node, q, `$node[@cat=("cp","sv1","smain","ppres","ppart","inf","ti","oti","du","whq") or @ud:pos="SCONJ"]`) {
		return "acl"
	}
	// oa zinnen tussen haakjes
	if TEST(node, q, `$node[@ud:pos= ("ADV","ADP","VERB","CCONJ") or @cat="advp"]`) {
		return "amod"
	}
	// VERB= aanstormend etc -> amod, ADV = nagenoeg alle prijzen, slechts 4 euro --> amod
	// CCONJ = opdrachten zoals:   --> amod
	if node.Rel == "det" {
		return "det" // empty determiners in gapping?
	}
	if node.Index > 0 {
		return "ERROR_INDEX_NMOD"
	}
	return "ERROR_NO_LABEL_NMOD"
}

func labelVmod(node *NodeType, q *Context) string {
	if TEST(node, q, `$node[@cat="pp"]/node[@rel="vc"]`) {
		return "advcl"
	}
	if TEST(node, q, `$node[ (  node[@rel="hd" and @lemma="door"]
	                                  or (@pt="bw" and ends-with(@lemma,"door"))
	                                  )
	                                  and ../self::*[@cat="ppart"]/../node[@rel="hd" and @lemma=("zijn","worden")]
	                                  and ../../node[@rel="su"]/@index = ../node[@rel="obj1"]/@index
	                              ]`) {
		return "obl:agent"
		/*
			but NOT: door rookontwikkeling om het leven gekomen
			-- already filtered by constraint of su/obj1 control
			NOT: bij Bakema is een stoeptegel door de ruit gegooid
			NO/YES: hierdoor werd Prince door het grote publiek ontdekt
		*/
	}
	if TEST(node, q, `$node[@cat=("pp","np","conj","mwu") or @ud:pos=("NOUN","PRON","PROPN","X","PUNCT","SYM") ]`) {
		return "obl"
	}
	if TEST(node, q, `$node[@cat=("cp","sv1","smain","ssub","ppres","ppart","ti","oti","inf","du","whq","whrel","rel")]`) {
		return "advcl"
	}
	if TEST(node, q, `$node[@ud:pos= ("ADJ","ADV","ADP","VERB","SCONJ","INTJ")
	                      or @cat=("advp","ap")
	                      or (@cat="conj" and node/@ud:pos="ADV")]`) {
		return "advmod" // niet of nauwelijks
	}
	if node.udPos == "NUM" {
		return "nummod"
	}
	if node.Index > 0 {
		return "ERROR_INDEX_VMOD"
	}
	return "ERROR_NO_LABEL_VMOD"
}

// this function is now also used to distribute dependents in coordination in Enhanced UD , so lot more rels and contexts are possible
// and passives, as in " hun taal werd gediscrimineerd en verboden"
func nonLocalDependencyLabel(head, gap *NodeType, q *Context) string {
	if gap.Rel == "su" {
		return subjectLabel(gap, q)
	}
	if gap.Rel == "obj1" {
		if head.Rel == "su" {
			return "nsubj:pass"
		}
		return "obj" // ppart coordination -- see comment above
	}
	if gap.Rel == "obj2" {
		if head.udPos == "ADV" {
			return "advmod"
		}
		return "iobj"
	}
	if gap.Rel == "me" {
		return determineNominalModLabel(gap, q)
	}
	if gap.Rel == "predc" || gap.Rel == "predm" {
		return dependencyLabel(gap, q)
	}
	if gap.Rel == "pc" || gap.Rel == "ld" {
		if HGTEST(head, gap, q, `$head/node[@rel="obj1"]`) {
			return "nmod"
		}
		if HGTEST(head, gap, q, `$head[@ud:pos=("ADV", "ADP") or @cat=("advp","ap")]`) {
			return "advmod" // waar precies zit je ..
		}
		return "ERROR_NO_LABEL_INDEX_PC"
	}
	if gap.Rel == "sup" || gap.Rel == "pobj1" {
		return "expl" // waar het om gaat is dat hij scoort, het is 1881 en dertien jaar geleden dat ...
	}
	if gap.Rel == "mwp" {
		return dependencyLabel(gap.parent, q) //wat heb je voor boeken gelezen
	}
	if gap.Rel == "vc" {
		return "ccomp" // wat ik me afvraag is of hij komt -- CLEFT
	}
	if HGTEST(head, gap, q, `$gap[@rel="mod" and ../@cat=("np","pp")]`) { // voornamelijk in kloosters en door vrouwen
		return modLabelInsideNp(head, q)
	}
	if HGTEST(head, gap, q, `$gap[@rel="mod" and ../@cat=("sv1","smain","ssub","inf","ppres","ppart","oti","ap","advp")]`) {
		return labelVmod(head, q)
	}
	if gap.Rel == "mod" || gap.Rel == "spec" { // spec only used for funny coord
		if HGTEST(head, gap, q, `$head[@cat=("pp","np") or @ud:pos=("NOUN","PRON")]`) {
			return "nmod"
		}
		if HGTEST(head, gap, q, `$head[@ud:pos=("ADV","ADP") or @cat= ("advp","ap","mwu","conj")]`) {
			return "advmod" // hoe vaak -- AP, daar waar, waar en wanneer, voor als rhd
		}
		return "ERROR_NO_LABEL_INDEX_MOD"
	}
	if gap.Rel == "det" {
		return detLabel(head, q)
	}
	if HGTEST(head, gap, q, `$gap[@rel="hd"] and $head[@ud:pos=("ADV","ADP")]`) { // waaronder A, B, en C
		return "case"
	}
	if gap.Rel == "du" || gap.Rel == "dp" {
		return "parataxis"
	}
	return "ERROR_NO_LABEL_INDEX"
}

func externalHeadPosition(node *NodeType, q *Context) int {

	if node.Rel == "hd" && (node.udPos == "ADP" || node.parent.Cat == "pp") {
		// vol vertrouwen
		if TEST(node, q, `$node/../node[@rel="predc"]`) { // TODO: dit kan efficiënter: meerdere keren zoeken naar zelfde set nodes
			// met als titel
			return internalHeadPosition(firstnode(FIND(node, q, `$node/../node[@rel="predc"]`)), q)
		}
		if TEST(node, q, `$node/../node[@rel=("obj1","vc","se","me")]`) {
			// adding pt/cat enough for gapping cases?
			if TEST(node, q, `$node/../node[@rel=("obj1","vc","se","me") and (@pt or @cat)]`) {
				return internalHeadPositionWithGapping(firstnode(FIND(node, q, `$node/../node[@rel=("obj1","vc","se","me")]`)), q)
			}
			if TEST(node, q, `$node/../node[@rel=("obj1","vc","se","me") and @index = ancestor::node/node[@rel=("rhd","whd")]/@index]`) {
				return internalHeadPosition(firstnode(FIND(node, q, `$node/ancestor::node/node[@rel=("rhd","whd")
	                                                             and @index = $node/../node[@rel=("obj1","vc","se","me")]/@index]`)), q)
			}
			if TEST(node, q, `$node/../node[@rel="pobj1"]`) {
				return internalHeadPosition(firstnode(FIND(node, q, `$node/../node[@rel="pobj1"]`)), q)
			}
			// in de eerste rond --> typo in LassySmall/Wiki , binnen en [advp later buiten ]
			return externalHeadPosition(node.parent, q)
		} else {
			return externalHeadPosition(node.parent, q)
		}
	}
	aux := auxiliary1(node, q)
	if node.Rel == "hd" && (aux == "aux" || aux == "aux:pass") {
		// aux aux:pass cop
		if TEST(node, q, `$node/../node[@rel=("vc","predc") and (@pt or (@cat and node[@pt or @cat]))]`) {
			// skip vc with just empty nodes
			return internalHeadPositionWithGapping(firstnode(FIND(node, q, `$node/../node[@rel=("vc","predc")]`)), q)
		}
		// if ($node/../node[@rel="predc"]/@index = $node/../../node[@rel="whd"]/@index)
		//     then local:internal_head_position($node/../../node[@rel="whd"])
		return externalHeadPosition(node.parent, q) // gapping, but does it ever occur with aux?? with cop: hij was en blijft nog steeds een omstreden figuur
	}
	/*

	      else if ($node[@rel="hd" and local:auxiliary($node) eq 'cop'] )
	             then if ($node/../node[@rel="predc" and (@pt or @cat)])
	       	        then local:internal_head_position_with_gapping(($node/../node[@rel="predc"])[1])
	       			else if ($node/../node[@rel="predc"]/@index = $node/ancestor::node/node[@rel=("rhd","whd")]/@index)
	       			      then local:internal_head_position($node/ancestor::node/node[@rel=("rhd","whd") and @index = $node/../node[@rel="predc"]/@index] )
	       			      else local:external_head_position($node/..)  (: gapping, but could it??:)

	      else if ($node[@rel=("hd","nucl","body") ] )
	       then if ($node/../node[@rel="hd"]/number(@begin) < $node/number(@begin) )
	             then local:internal_head_position($node/../node[@rel="hd" and number(@begin) < $node/number(@begin)] ) (: dan moet je moet :)
	             else local:external_head_position($node/..)

	      else if ( $node[@rel="predc"] )
	       then if   ($node[../node[@rel=("obj1","se","vc")] and ../node[@rel="hd" and (@pt or @cat)]] )
	             then if ( $node/../node[@rel="hd" and @ud:pos="ADP"] )
	                   then local:external_head_position($node/..) (: met als presentator Bruno W , met als gevolg [vc dat ...]:)
	                   else local:internal_head_position($node/../node[@rel="hd"])
	             else if  ( $node/..[@cat=("np","ap") and node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX") ]  ]  )
	                             (: reduced relatives , make sure head is not empty (ellipsis) :)
	                             (: also : ap with predc: actief als schrijver :)
	                   then local:internal_head_position($node/../node[@rel="hd"])
	                   else if ($node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos=("AUX","ADP"))] )  (: [met als titel] -- obj1/vc missing :)
	                   	 then local:internal_head_position($node/../node[@rel="hd"])
	                   	 else local:external_head_position($node/..) (: covers gapping as well? :)

	      else if ( $node[@rel=("obj1","se","me") and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] )
	       then if ($node/../node[@rel="predc"] )
	             then local:internal_head_position($node/../node[@rel="predc"])
	             else local:external_head_position($node/..)

	      else if ( $node[@rel="pobj1" and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] )
	       then if ($node/../node[@rel="vc"])
	             then local:internal_head_position($node/../node[@rel="vc"])
	             else local:external_head_position($node/..)

	      else if ($node[@rel="mod" and not(../node[@rel=("obj1","pobj1","se","me")]) and (../@cat="pp" or ../node[@rel="hd" and @ud:pos="ADP"])])   (: mede op grond hiervan :)
	       (: daarom dus :)
	            then if ($node/../node[@rel=("hd","su","obj1","vc") and (@pt or @cat)] )
	                  then local:internal_head_position_with_gapping($node/..)
	                  else local:external_head_position($node/..) (: gapping :)


	     else if ($node[@rel=("cnj","dp","mwp")])
	      then if ( deep-equal($node,local:leftmost($node/../node[@rel=("cnj","dp","mwp")])) )
	            then local:external_head_position($node/..)
	            else if ($node[@rel="cnj"])
	                 then local:head_position_of_conjunction($node)
	                 else local:internal_head_position_with_gapping($node/..)

	     else if ($node[@rel="cmp" and ../node[@rel="body"]])
	      then local:internal_head_position_with_gapping($node/../node[@rel="body"][1])

	     else if ($node[@rel="--" and @cat] )
	     	then if ($node[@cat="mwu"])
	             then if ($node/../node[@cat and not(@cat="mwu")]  )    (: fix for multiword punctuation in Alpino output :)
	                   then local:internal_head_position($node/../node[@cat and not(@cat="mwu")][1])
	                   else local:external_head_position($node/..)
	       else local:external_head_position($node/..)

	     else if ( $node[@rel="--" and @ud:pos] )
	      then if ($node[@ud:pos = ("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")
	                     and ../node[@rel="--" and
	                                 not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")) ]
	                    ] )
	            then local:internal_head_position_with_gapping($node/../node[@rel="--" and not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","ADP","ADV","DET","PROPN","NUM","PRON"))][1])
	            else if ( $node/../node[@cat]  )
	                  then local:internal_head_position($node/../node[@cat][1])
	                  else if ($node[@ud:pos="PUNCT" and count(../node) > 1])
	                        then if ($node/../node[not(@ud:pos="PUNCT")] )
	                              then local:internal_head_position($node/../node[not(@ud:pos="PUNCT")][1])
	                              else if ( deep-equal($node,local:leftmost($node/../node[@rel="--" and (@cat or @pt)]) ) )
	                                    then local:external_head_position($node/..)
	                                    else "1" (: ie end of first punct token :)
	                        else if ($node/..) then local:external_head_position($node/..)
	       else "ERROR_NO_HEAD_FOUND"

	     else if ($node[@rel=("dlink","sat","tag")])
	      then if ($node/../node[@rel="nucl"])
	            then local:internal_head_position_with_gapping($node/../node[@rel="nucl"])
	            else "ERROR_NO_EXTERNAL_HEAD"

	     else if ($node[@rel="vc"])
	       then if ($node/../node[@rel="hd" and
	                              ( @ud:pos="AUX" or
	                                $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index
	                              )
	                          ]
	                     and not($node/../node[@rel="predc"]) )
	             then local:external_head_position($node/..)
	             else if ($node/../@cat="pp") (: eraan dat .. :)
	                    then local:external_head_position($node/..)
	                    else if ($node/../node[@rel=("hd","su") and (@pt or @cat)] )
	                         then local:internal_head_position_with_gapping($node/..)
	                         else local:external_head_position($node/..)

	     else if ($node[@rel="whd" or @rel="rhd"])
	      then if ($node[@index])
	            then local:external_head_position( ($node/../node[@rel="body"]//node[@index = $node/@index ])[1] )
	            else local:internal_head_position($node/../node[@rel="body"])

	   (: we need to select the original node and not the result of
	      following-cnj-sister, as that has no global context
	      and global context is needed where the hd is an index node...
	      unfortunately, nodes are completely identical in some
	      elliptical cases, select last() as brute force solution :)
	     else if ($node[@rel="crd"])
	      then local:internal_head_position_with_gapping(
	      	           $node/../node[@rel="cnj" and
	            	                 @begin=local:following-cnj-sister($node)/@begin and
	            	                 @end=local:following-cnj-sister($node)/@end
	            	                ][last()] )

	     else if ($node[@rel="su"])
	      then if ($node/../node[@rel="hd" and (@pt or @cat)]) (: gapping :)
	           then local:internal_head_position_with_gapping($node/..) (: ud head could still be a predc :)
	            (: only for 1 case where verb is missing -- die eigendom ... (no verb)) :)
	           else if ($node[../node[@rel="predc"] and not(../node[@rel="hd"])] )
	                then local:internal_head_position($node/../node[@rel="predc"])
	            	 else local:external_head_position($node/..) (: this probably does no change anything, as we are still attaching to head of left conjunct :)

	     else if ($node[@rel="obj1"])
	      then if ($node/../node[@rel=("hd","su") and (@pt or @cat)]) (: gapping, as su but now su could be head as well :)
	            then local:internal_head_position_with_gapping($node/..)
	            else local:external_head_position($node/..)

	     else if ($node[@rel="pc"])
	      then if ($node/../node[@rel=("hd","su","obj1") and (@pt or @cat)]) (: gapping, as su but now su could be head as well :)
	            then local:internal_head_position_with_gapping($node/..)
	            else local:external_head_position($node/..)

	     else if ($node[@rel="mod"])
	      then if ($node/../node[@rel=("hd","su","obj1","pc","predc","body") and (@pt or @cat)]) (: gapping, as su but now su or obj1  could be head as well :)
	            then local:internal_head_position_with_gapping($node/..)
	            else if ($node/../node[@rel="mod" and (@cat or @pt)])
	                  then if  (deep-equal($node,local:leftmost($node/../node[@rel="mod" and (@pt or @cat)])) ) (: gapping with multiple mods :)
	                        then local:external_head_position($node/..)
	                        else local:internal_head_position_with_gapping($node/..)
	                  else if ( $node/../../node[@rel="su" and (@pt or @cat)]  )  (: an mod in an otherwise empty tree (after fixing heads in conj) :)
	                       then local:internal_head_position($node/../../node[@rel="su"])
	                       else local:external_head_position($node/..) (: an empty mod in an otherwise empty tree
	                                                                 -- mod is co-indexed with rhd, rest is elided,
	                                                                 LassySmall4/wiki-7064/wiki-7064.p.28.s.3.xml :)

	   else if ($node[@rel=("app","det")])
	      then if ($node/../node[@rel=("hd","mod") and (@pt or @cat)]) (: gapping with an app (or a det)! :)
	            then local:internal_head_position_with_gapping($node/..)
	            else local:external_head_position($node/..)

	     else if ($node[@rel="top"])
	      then "0"

	     else if ( $node[not(@rel="hd")] )
	       then    local:internal_head_position_with_gapping($node/..)

	     else    'ERROR_NO_EXTERNAL_HEAD'
	   } ;
	*/
	return TODO
}

func internalHeadPosition(node *NodeType, q *Context) int {

	if node.Cat == "pp" {
		// if ($node/node[@rel="hd" and @pt=("bw","n")] )  ( n --> TEMPORARY HACK to fix error where NP is erroneously tagged as PP )
		// then $node/node[@rel="hd"]/@end
		if TEST(node, q, `$node/node[@rel=("obj1","pobj1","se")]`) {
			return internalHeadPosition(firstnode(FIND(node, q, `$node/node[@rel=("obj1","pobj1","se")]`)), q)
		}
		if TEST(node, q, `$node/node[@rel="hd"]`) {
			// if ($node/@cat="mwu")  ( mede [op grond hiervan] )
			//     local:internal_head_position($node/node[@rel="hd"] )
			return internalHeadPosition(firstnode(FIND(node, q, `$node/node[@rel="hd"]`)), q)
		}
		if len(node.Node) > 0 {
			return internalHeadPosition(node.Node[0], q)
		}
		return internalHeadPosition(noNode, q)
	}

	if node.Cat == "mwu" {
		return firstint(FIND(node, q, `$node/node[@rel="mwp" and not(../node/@begin < @begin)]/@end`))
	}

	/*

	     else if ($node[@cat="conj"])
	     then    local:internal_head_position(local:leftmost($node/node[@rel="cnj"]))

	     else if ( $node/node[@rel="predc"] )
	          then if ($node/node[@rel="hd" and @ud:pos="AUX"])
	               then local:internal_head_position($node/node[@rel="predc"])
	               else if ( not($node/node[@rel="hd"]) )      (: cases where copula is missing by accident (ungrammatical, not gapping) :)
	                    then local:internal_head_position($node/node[@rel="predc"])
	                    else local:internal_head_position($node/node[@rel="hd"])

	     else if ($node[node[@rel="vc"] and
	                    node[@rel="hd" and
	                         ( @ud:pos="AUX" or
	                           $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index
	                          )
	                        ]
	                   ]
	             )
	     then    local:internal_head_position($node/node[@rel="vc"])

	     else if ( $node/node[@rel="hd"])
	     then local:internal_head_position($node/node[@rel="hd"][1])

	     else if ( $node/node[@rel="body"])
	     then    local:internal_head_position($node/node[@rel="body"][1])

	     else if ( $node/node[@rel="dp"])
	     then    local:internal_head_position(local:leftmost($node/node[@rel="dp"]))
	          (: sometimes co-indexing leads to du's starting at same position ... :)

	     else if ( $node/node[@rel="nucl"])
	     then    local:internal_head_position($node/node[@rel="nucl"][1])

	     else if ( $node/node[@cat="du"]) (: is this neccesary at all? , only one referring to cat, and causes problems if applied before @rel=dp case... :)
	     then    local:internal_head_position($node/node[@cat ="du"][1])

	     else if ( $node[@word] )
	     then    $node/@end

	     (: distinguish empty nodes due to gapping/RNR from nonlocal dependencies
	        use 'empty head' as string to signal precence of gapping/RNR :)
	     else if ($node[@index and not(@word or @cat)] )
	        then if ($node/ancestor::node[@rel="top"]//node[@rel=("whd","rhd") and @index = $node/@index and (@word or @cat)] )
	             then local:internal_head_position($node/ancestor::node[@rel="top"]//node[@index = $node/@index and (@word or @cat)] )
	             else "empty head"

	     else    'ERROR_NO_INTERNAL_HEAD'
	   };
	*/
	return TODO
}

func internalHeadPositionWithGapping(node *NodeType, q *Context) int {
	return TODO
}

func enhancedDependencies(q *Context) {
	// TODO

	// na elke invoeging van nieuwe node: inspect(q) ??
}

/*
func auxiliary(nodes []*NodeType, q *Context) string {
	if len(nodes) != 1 {
		return "ERROR_AUXILIARY_FUNCTION_TAKES_EXACTLY_ONE_ARG"
	}
	return auxiliary1(nodes[0], q)
}
*/

func auxiliary1(node *NodeType, q *Context) string {

	if node.Pt != "ww" {
		return "ERROR_NO_VERB"
	}
	if node.Rel != "hd" {
		return "verb"
	}

	if TEST(node, q, `$node[not(../node[@rel=("obj1","se","vc")]) and
			        (: ud documentation suggests 1 cop per lg, van Eynde suggests much more, compromise: the traditional ones :)
			        @lemma=("zijn","lijken","blijken","blijven","schijnen","heten","voorkomen","worden","dunken") and
		                 ( contains(@sc,'copula') or
		                   contains(@sc,'pred')   or
		                   contains(@sc,'cleft')  or
		                   ../node[@rel="predc"]
		                 ) ]`) {
		return "cop"
	}

	if TEST(node, q, `$node[@lemma=("zijn","worden") and
	                  ( @sc="passive"  or
	                  	 ( ../node[@rel="vc"] and
	                        ( ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="obj1"]/@index or
	                          ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="cnj"]/node[@rel="obj1"]/@index or
	                          ../node[@rel="vc" and not(@pt or @cat)]/@index =
	                              $indexnodes[@rel="vc" and node[@rel="obj1"]/@index = $node/../node[@rel="su"]/@index]/@index
	                         or not(../node[@rel="su"])
	                         )
	                     )
	                  ) ]`) {
		return "aux:pass"
	}

	if TEST(node, q, `
	  (: krijgen passive with iobj control :)
	            $node[@lemma="krijgen" and
	  	              ( ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="obj2"]/@index or
	                    ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="cnj"]/node[@rel="obj2"]/@index
	                  )]`) {
		return "aux:pass"
	}

	if TEST(node, q, `
	  (: alpino has no principled distinction between AUX and VERB, should be TAME verbs semantically, we follow ENGLISH :)
	          $node[not(../node[@rel="predc"]) and  (: hij heeft als opdracht stammen uit elkaar te houden  :)
	                 ( starts-with(@sc,'aux') or
	                   ( ../node[@rel="vc"  and
	                              ( @cat=("ppart","inf","ti") or
	                                ( @cat="conj" and node[@rel="cnj" and @cat=("ppart","inf","ti")] ) or
	                                ( @index and not(@pt or @cat))  (: dangling aux in gapped coordination :)
	                              )
	                            ]   and
	                     @lemma=("blijken","hebben","hoeven","kunnen","moeten","mogen","zijn","zullen")
	                   )
	                 )
	               ]`) {
		return "aux"
	}

	/*
	  else if ($node[@pt="ww"] )
	  then "verb"
	*/
	return "verb"

	/*
	  else "ERROR_NO_VERB"
	*/

}

////////////////////////////////////////////////////////////////

func leftmost(nodes []interface{}) *NodeType {
	switch len(nodes) {
	case 0:
		return noNode
	case 1:
		return nodes[0].(*NodeType)
	}
	sort.Slice(nodes, func(i, j int) bool {
		// solve cases where begin is identical (hij is en blijft omstreden)??
		ii := nodes[i].(*NodeType)
		jj := nodes[j].(*NodeType)
		if ii.Begin != jj.Begin {
			return ii.Begin < jj.Begin // ints
		}
		if ii.End != jj.End {
			return ii.End < jj.End // ints
		}
		return ii.Index < jj.Index // ints
	})
	return nodes[0].(*NodeType)
}

func firstnode(nodes []interface{}) *NodeType {
	if len(nodes) > 0 {
		return nodes[0].(*NodeType)
	}
	return noNode
}

func firstint(ii []interface{}) int {
	if len(ii) > 0 {
		return ii[0].(int)
	}
	return ERROR_NO_VALUE
}
