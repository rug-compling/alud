package alud

import (
	"fmt"
	"strings"
)

func addFeatures(q *context) {
	for _, node := range q.ptnodes {
		switch node.udPos {
		case "NOUN", "PROPN":
			nominalFeatures(node, q)
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
		if node.Rel == "mwp" && node.Spectype != "symb" && node.Spectype != "deeleigen" && node.Ntype != "eigen'" && node.Begin == node.parent.Begin {
			extpos(node, q)
		}
	}
}

func nominalFeatures(node *nodeType, q *context) {
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
		panic(fmt.Sprintf("Irregular gender for %s:%s : %s", number(node.End), node.Word, node.Genus))
	}
	switch node.Getal {
	case "ev":
		node.udNumber = "Sing"
	case "mv":
		node.udNumber = "Plur"
	case "":
		node.udNumber = ""
	default:
		panic(fmt.Sprintf("Irregular number for %s:%s : %s", number(node.End), node.Word, node.Getal))
	}
	switch node.Graad {
	case "dim":
		node.udDegree = ""
	case "basis":
		node.udDegree = ""
	case "":
		node.udDegree = ""
	default:
		panic(fmt.Sprintf("Irregular degree for %s:%s : %s", number(node.End), node.Word, node.Graad))
	}
}

func adjectiveFeatures(node *nodeType, q *context) {
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
		panic(fmt.Sprintf("Irregular degree for %s:%s : %s", number(node.End), node.Word, node.Graad))
	}
}

func pronounFeatures(node *nodeType, q *context) {
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
		panic(fmt.Sprintf("Irregular prontype for %s:%s : %s", number(node.End), node.Word, node.Vwtype))
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
		panic(fmt.Sprintf("Irregular person for %s:%s : %s", number(node.End), node.Word, node.Persoon))
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
		panic(fmt.Sprintf("Irregular case for %s:%s : %s", number(node.End), node.Word, node.Naamval))
	}
}

func verbalFeatures(node *nodeType, q *context) {

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
		panic(fmt.Sprintf("Irregular verbform for %s:%s : %s", number(node.End), node.Word, node.Wvorm))
	}

	switch node.Pvtijd {
	case "verl":
		node.udTense = "Past"
	case "tgw", "conj":
		node.udTense = "Pres"
	case "":
		node.udTense = ""
	default:
		panic(fmt.Sprintf("Irregular tense for %s:%s : %s", number(node.End), node.Word, node.Pvtijd))
	}

	switch node.Pvagr {
	case "ev", "met-t":
		node.udNumber = "Sing"
	case "mv":
		node.udNumber = "Plur"
	case "":
		node.udNumber = ""
	default:
		panic(fmt.Sprintf("Irregular number for %s:%s : %s", number(node.End), node.Word, node.Pvagr))
	}
}

func determinerFeatures(node *nodeType, q *context) {
	switch node.Lwtype {
	case "bep":
		node.udDefinite = "Def"
	case "onbep":
		node.udDefinite = "Ind"
	case "":
		node.udDefinite = ""
	default:
		panic(fmt.Sprintf("Irregular definite for %s:%s : %s", number(node.End), node.Word, node.Lwtype))
	}
}

func specialFeatures(node *nodeType, q *context) {
	switch node.Spectype {
	case "vreemd":
		node.udForeign = "Yes"
	case "afk":
		node.udAbbr = "Yes"
	}
}

func extpos(node *nodeType, q *context) {
	if node.Frame == "adjective(pred(nonadv))" {
		if node.parent.Rel == "hd" && node.Word == "in" { // in plaats van
			node.udExtPos = "ADP"
		} else {
			node.udExtPos = "ADJ"
		}
	}
	if node.Frame == "adjective(no_e(adv))" && node.Word == "dicht" { // frame error
		node.udExtPos = "ADP"
	}
	if node.Frame == "adjective(no_e(locadv))" && node.Word == "dicht" { // dicht bij, frame error
		node.udExtPos = "ADP"
	}
	if node.Frame == "adjective(postn_no_e(dir_locadv))" && node.Word == "links" { // links boven ons
		node.udExtPos = "ADP"
	}
	if node.Frame == "pre_num_adv(pl_indef)" && node.Word == "tegen" { // tegen de vier uur
		node.udExtPos = "DET"
	}
	switch node.Frame {
	case "adverb", "modal_adverb", "modal_adverb(noun_prep)", "intensifier", "loc_adverb", "er_adverb([ter,gelegenheid,van])", "waar_adverb(vanuit)",
		"post_wh_adverb", "sentence_adverb", "pre_num_adv(both)", "pre_wh_adverb", "adjective(pred(adv))", "adjective(pred(adv),pp(naar))",
		"tmp_adverb", "pre_num_adv", "adjective(het_st(adv))", "adjective(het_st(adv),subject_sbar)", "adjective(pred(dir_locadv))", "comp_adverb(als)",
		"adjective(pred(adv),transitive)":
		node.udExtPos = "ADV"
	case "complementizer(inf)":
		node.udExtPos = "SCONJ"
	case "conj(maar)", "conj('tot en met')", "conj(en)", "etc", "complementizer(np)", "conj('laat staan')":
		node.udExtPos = "CCONJ"
	case "adjective(meer)", "adjective(pred(padv))":
		node.udExtPos = "ADJ"
	case "complementizer(te)", "pp", "pp(te)", "pp(van)", "particle([af,aan])", "noun(both,both,both)", "preposition(dichtbij)":
		node.udExtPos = "ADP"
	case "tmp_np":
		node.udExtPos = "PRON" // counter intuitive, but avoids validation errors
	}
	if strings.Contains(node.Frame, "preposition") {
		node.udExtPos = "ADP"
	}
	if strings.Contains(node.Frame, "pp(van)") {
		node.udExtPos = "ADP"
	}
	if strings.Contains(node.Frame, "fixed_part") {
		node.udExtPos = "ADP"
	}
	if strings.Contains(node.Frame, "proper_name") { // n.b. also ensure rel is flat
		node.udExtPos = "PROPN"
	}
	if strings.Contains(node.Frame, "number") { // n.b. also ensure rel is flat
		node.udExtPos = "PRON" // value should be NUM but not allowed
	}
	if strings.Contains(node.Frame, "pronoun") {
		node.udExtPos = "PRON"
	}
	if node.parent.Rel == "cmp" || node.parent.Rel == "dlink" { // dat wil zeggen
		node.udExtPos = "SCONJ"
	}
}
