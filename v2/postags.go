package alud

import (
	"fmt"
)

func addPosTags(q *context) {
	for _, node := range q.ptnodes {
		node.udPos = universalPosTags(node, q)
	}
}

func universalPosTags(gonode *nodeType, q *context) string {
	node := gonode.node
	pt := node.Attr("pt")
	rel := node.Attr("rel")
	parent := node.Parent()
	prel := parent.Attr("rel")
	pcat := parent.Attr("cat")

	if pt == "let" {
		if rel == "--" {
			if parent.Attr("pt") != "let" || parent.Attr("begin") < node.Attr("begin") {
				return "PUNCT"
			}
		}
		return "SYM"
	}
	if pt == "adj" {
		if rel == "det" {
			return "DET"
		}
		if rel == "hd" && pcat == "pp" {
			// vol vertrouwen
			return "ADP"
		}
		if rel == "crd" {
			// respectievelijk
			return "CCONJ"
		}
		if rel == "cmp" { // inclusief hijzelf
			return "SCONJ"
		}
		return "ADJ" // exceptions forced by 2.4 validation
	}
	if pt == "bw" {
		if rel == "crd" {
			return "CCONJ"
		}
		if prel == "det" { // zo min mogelijk, genoeg geld om een ijsje te kopen
			return "DET"
		}
		return "ADV"
	}
	if pt == "lid" {
		return "DET"
	}
	if pt == "n" {
		if node.Attr("ntype") == "eigen" {
			return "PROPN"
		}
		return "NOUN"
	}
	if pt == "spec" {
		spectype := node.Attr("spectype")
		if spectype == "deeleigen" {
			return "PROPN"
		}
		if rel == "crd" { // resp. 33 en 44 %
			return "CCONJ" // exception needed to avoid validation errors
		}
		if rel == "det" || prel == "det" { // 49% (symb), zes- a zevenduizend (afgebr), the (vreemd) zijn/haar (enof)
			return "DET" // exception needed to avoid validation errors
		}
		if spectype == "symb" { // 49%
			return "SYM"
		}
		if spectype == "afk" && rel == "hd" && pcat == "ap" { // incl. Rwanda
			return "ADJ"
		}
		return "X" // vreemd afgebr meta enof
	}
	if pt == "tsw" {
		if rel == "mod" {
			return "ADV"
		}
		return "INTJ"
	}
	if pt == "tw" {
		if node.Attr("numtype") == "rang" {
			return "ADJ"
		}
		if rel == "hd" && prel == "mod" && (pcat == "advp" || pcat == "ap") { // zoveel + obcomp: zoveel mogelijk, zoveel te wensen over dat ...
			return "ADV"
		}
		return "NUM"
	}
	if pt == "vz" {
		return "ADP" // v2: do not use PART for SVPs and complementizers
	}
	if pt == "vnw" {
		vwtype := node.Attr("vwtype")
		if rel == "det" && vwtype != "bez" {
			return "DET"
		}
		if rel == "hd" && pcat == "detp" && vwtype != "bez" { // niet veel meer dan
			// added != bez to account for 'al zijn boeken' GB 03/11/22
			return "DET"
		}
		if rel == "hd" && (prel == "mod" || prel == "rhd") { // heel wat fleuriger, hoe meer ik over deze oorlog hoor,
			return "ADV"
		}
		if rel == "mod" && prel == "det" { // [detp/det vnw/al deze] stripreeksen] --> al wordt advmod
			return "ADV"
		}

		pdtype := node.Attr("pdtype")
		if pdtype == "adv-pron" {
			if rel == "pobj1" {
				return "PRON"
			}
			return "ADV"
		}
		if (rel == "mod" || (rel == "hd" && prel == "mod")) && pdtype == "grad" {
			// veel minder
			return "ADV"
		}
		return "PRON"
	}
	if pt == "vg" {
		if rel == "rhd" && node.Attr("lemma") == "zoals" { //
			return "ADV"
		}
		if node.Attr("conjtype") == "neven" {
			return "CCONJ" // V2: CONJ ==> CCONJ
		}
		return "SCONJ"
	}
	if pt == "ww" {
		aux, err := auxiliary1(node, q)
		if err != nil {
			panic(fmt.Sprintf("No pos found for %s:%s - %v", number(gonode.end), node.Attr("word"), err))
		}
		if aux == "verb" {
			return "VERB"
		}
		return "AUX" // v2: cop and aux:pass --> AUX  (already in place in v1?)
	}
	if pt == "na" { // only in automatic parser output if something went wrong, added for robustness
		return "X"
	}
	panic(fmt.Sprintf("No pos found for %s:%s", number(gonode.end), node.Attr("word")))
}
