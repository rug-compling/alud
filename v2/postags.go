package alud

import (
	"fmt"
)

func addPosTags(q *context) {
	for _, node := range q.ptnodes {
		node.udPos = universalPosTags(node, q)
	}
}

func universalPosTags(node *nodeType, q *context) string {
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
		if rel == "cmp" { // inclusief hijzelf
			return "SCONJ"
		}
		return "ADJ" // exceptions forced by 2.4 validation
	}
	if pt == "bw" {
		// case is obsolete, see discussion with Leiden/Gent july 2025
		// if rel == "crd" {
	        // 		return "CCONJ"
		// } 
		// added rel=hd to avoid matching with 'al' in 'al zijn films' GB 5-11-24,
		// see also https://github.com/UniversalDependencies/docs/issues/1059#issuecomment-2453407662
		if node.parent.Rel == "det" && rel == "hd" { // zo min mogelijk, genoeg geld om een ijsje te kopen
			return "DET"
		}
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
		if node.Rel == "crd" { //resp. 33 en 44 %
			return "CCONJ" // exception needed to avoid validation errors
		}
		if node.Rel == "det" || node.parent.Rel == "det" { //49% (symb), zes- a zevenduizend (afgebr), the (vreemd) zijn/haar (enof)
			return "DET" // exception needed to avoid validation errors
		}
		if node.Spectype == "symb" { // 49%
			return "SYM"
		}
		if node.Spectype == "afk" && node.Rel == "hd" && node.parent.Cat == "ap" { // incl. Rwanda
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
		if node.Numtype == "rang" {
			return "ADJ"
		}
		if rel == "hd" && node.parent.Rel == "mod" && (node.parent.Cat == "advp" || node.parent.Cat == "ap") { // zoveel + obcomp: zoveel mogelijk, zoveel te wensen over dat ...
			return "ADV"
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
		if rel == "hd" && node.parent.Cat == "detp" && node.Vwtype != "bez" { // niet veel meer dan
			// added != bez to account for 'al zijn boeken' GB 03/11/22
			return "DET"
		}
		if rel == "hd" && (node.parent.Rel == "mod" || node.parent.Rel == "rhd") { // heel wat fleuriger, hoe meer ik over deze oorlog hoor,
			return "ADV"
		}
		if rel == "mod" && node.parent.Rel == "det" { // [detp/det vnw/al deze] stripreeksen] --> al wordt advmod
			return "ADV"
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
		if rel == "rhd" && node.Lemma == "zoals" { //
			return "ADV"
		}
		if node.Conjtype == "neven" {
			return "CCONJ" // V2: CONJ ==> CCONJ
		}
		return "SCONJ"
	}
	if pt == "ww" {
		aux, err := auxiliary1(node, q)
		if err != nil {
			panic(fmt.Sprintf("No pos found for %s:%s - %v", number(node.End), node.Word, err))
		}
		if aux == "verb" {
			return "VERB"
		}
		return "AUX" // v2: cop and aux:pass --> AUX  (already in place in v1?)
	}
	if pt == "na" { // only in automatic parser output if something went wrong, added for robustness
		return "X"
	}
	panic(fmt.Sprintf("No pos found for %s:%s", number(node.End), node.Word))
}
