package main

import (
	//	"github.com/kr/pretty"

	"fmt"
	//"sort"
	"strings"
)

type Doer interface {
	Do(subdoc []interface{}, q *Context) []interface{}
}

type Parent interface {
	Children() []interface{}
	Descendants() []interface{}
	DescendantsOrSelf() []interface{}
}

const (
	cmp__lt = iota
	cmp__gt
	collect__ancestors__node
	collect__attributes__begin
	collect__attributes__end
	collect__attributes__cat
	collect__attributes__index
	collect__attributes__lemma
	collect__attributes__naamval
	collect__attributes__pt
	collect__attributes__rel
	collect__attributes__sc
	collect__attributes__spectype
	collect__attributes__ud_3apos
	collect__attributes__vwtype
	collect__attributes__word
	collect__child__node
	collect__descendant__node
	collect__descendant__or__self__node
	collect__descendant__or__self__type__node
	collect__parent__type__node
	collect__self__all__node
	collect__self__node
	elem__number__1
	elem__number__1000
	elem__string__ADJ
	elem__string__ADJ__ADV__ADP__VERB__SCONJ__INTJ
	elem__string__ADJ__X__ADV
	elem__string__ADP
	elem__string__ADP__ADV__ADJ__DET__PRON__CCONJ__NOUN__VERB__INT_2e_2e_2e
	elem__string__ADV
	elem__string__ADV__ADP
	elem__string__ADV__ADP__VERB__CCONJ
	elem__string__AUX
	elem__string__AUX__ADP
	elem__string__CCONJ
	elem__string__DET__PROPN__NOUN__ADJ__PRON__ADV__X
	elem__string__NOUN__PRON
	elem__string__NOUN__PRON__PROPN__X__PUNCT__SYM
	elem__string__NOUN__PRON__PROPN__X__PUNCT__SYM__INTJ
	elem__string__NOUN__PROPN
	elem__string__NOUN__PROPN__VERB
	elem__string__NUM
	elem__string__PRON
	elem__string__PRON__PROPN
	elem__string__PROPN
	elem__string__PUNCT
	elem__string__PUNCT__SYM
	elem__string__PUNCT__SYM__X
	elem__string__PUNCT__SYM__X__CONJ__NOUN__ADP__ADV__DET__PROPN___2e_2e_2e
	elem__string__PUNCT__SYM__X__CONJ__NOUN__PROPN__NUM__ADP__ADV___2e_2e_2e
	elem__string__SCONJ
	elem__string__VERB
	elem__string__aan
	elem__string__advp
	elem__string__advp__ap
	elem__string__advp__ap__mwu__conj
	elem__string__aux
	elem__string__ap
	elem__string__app
	elem__string__bez
	elem__string__blijken__hebben__hoeven__kunnen__moeten__moge_2e_2e_2e
	elem__string__body
	elem__string__bw
	elem__string__cleft
	elem__string__cmp
	elem__string__cnj
	elem__string__cnj__dp__mwp
	elem__string__conj
	elem__string__copula
	elem__string__cp
	elem__string__cp__sv1__smain__ppres__ppart__inf__ti__oti__du__w_2e_2e_2e
	elem__string__cp__sv1__smain__ssub__ppres__ppart__ti__oti__inf_2e_2e_2e
	elem__string__cp__whrel__whq__whsub
	elem__string__deeleigen
	elem__string__det
	elem__string__detp
	elem__string__dooien__gieten__hagelen__miezeren__misten__mo_2e_2e_2e
	elem__string__door
	elem__string__dp
	elem__string__du
	elem__string__du__smain
	elem__string__gen
	elem__string__het
	elem__string__hd
	elem__string__hd__body
	elem__string__hd__ld
	elem__string__hd__ld__vc
	elem__string__hd__mod
	elem__string__hd__nucl__body
	elem__string__hd__su
	elem__string__hd__su__obj1
	elem__string__hd__su__obj1__pc__predc__body
	elem__string__hd__su__obj1__vc
	elem__string__krijgen
	elem__string__me
	elem__string__minmin
	elem__string__mod
	elem__string__mod__pc__ld
	elem__string__mwp
	elem__string__mwu
	elem__string__mwu__np__pp__ap__detp__smain
	elem__string__np
	elem__string__np__ap
	elem__string__np__conj
	elem__string__np__pp
	elem__string__nucl
	elem__string__obj1
	elem__string__obj1__se
	elem__string__obj1__se__me
	elem__string__obj1__se__vc
	elem__string__obj1__vc__se__me
	elem__string__obj1__pobj1__se
	elem__string__obj1__pobj1__se__me
	elem__string__obj2
	elem__string__ontbreken
	elem__string__oti
	elem__string__passive
	elem__string__pc
	elem__string__pc__ld
	elem__string__pobj1
	elem__string__pp
	elem__string__pp__detp__advp
	elem__string__pp__np
	elem__string__pp__np__conj__mwu
	elem__string__ppart
	elem__string__ppart__inf__ti
	elem__string__pred
	elem__string__predc
	elem__string__rhd__whd
	elem__string__smain
	elem__string__su
	elem__string__su__obj1
	elem__string__su__obj1__predc__vc
	elem__string__sup
	elem__string__sv1__smain__ssub__inf__ppres__ppart__oti__ap__ad_2e_2e_2e
	elem__string__top
	elem__string__ti
	elem__string__vc
	elem__string__vc__predc
	elem__string__whd__rhd
	elem__string__whsub__ssub__ti__cp__oti
	elem__string__zijn__lijken__blijken__blijven__schijnen__het_2e_2e_2e
	elem__string__zijn__worden
	equal__is
	function__contains__2__args
	function__count__1__args
	// function__deep__equal__2__args
	function__ends__with__2__args
	function__not__1__args
	function__starts__with__2__args
	plus__plus
	plus__minus
	variable__allnodes
	variable__gap
	variable__head
	variable__indexnodes
	variable__node
	variable__tmp
	variable__subj
)

var (
	TRUE  = []interface{}{"true"}
	FALSE = []interface{}{}

	elemstrings = map[int]string{
		elem__string__ADJ: "ADJ",
		elem__string__ADJ__ADV__ADP__VERB__SCONJ__INTJ: "ADJ ADV ADP VERB SCONJ INTJ",
		elem__string__ADJ__X__ADV:                      "ADJ X ADV",
		elem__string__ADP:                              "ADP",

		elem__string__ADP__ADV__ADJ__DET__PRON__CCONJ__NOUN__VERB__INT_2e_2e_2e: "ADP ADV ADJ DET PRON CCONJ NOUN VERB INTJ",

		elem__string__ADV:                                    "ADV",
		elem__string__ADV__ADP:                               "ADV ADP",
		elem__string__ADV__ADP__VERB__CCONJ:                  "ADV ADP VERB CCONJ",
		elem__string__AUX:                                    "AUX",
		elem__string__AUX__ADP:                               "AUX ADP",
		elem__string__CCONJ:                                  "CCONJ",
		elem__string__DET__PROPN__NOUN__ADJ__PRON__ADV__X:    "DET PROPN NOUN ADJ PRON ADV X",
		elem__string__NOUN__PRON:                             "NOUN PRON",
		elem__string__NOUN__PRON__PROPN__X__PUNCT__SYM:       "NOUN PRON PROPN X PUNCT SYM",
		elem__string__NOUN__PRON__PROPN__X__PUNCT__SYM__INTJ: "NOUN PRON PROPN X PUNCT SYM INTJ",
		elem__string__NOUN__PROPN:                            "NOUN PROPN",
		elem__string__NOUN__PROPN__VERB:                      "NOUN PROPN VERB",
		elem__string__NUM:                                    "NUM",
		elem__string__PRON:                                   "PRON",
		elem__string__PRON__PROPN:                            "PRON PROPN",
		elem__string__PROPN:                                  "PROPN",
		elem__string__PUNCT:                                  "PUNCT",
		elem__string__PUNCT__SYM:                             "PUNCT SYM",
		elem__string__PUNCT__SYM__X:                          "PUNCT SYM X",

		elem__string__PUNCT__SYM__X__CONJ__NOUN__ADP__ADV__DET__PROPN___2e_2e_2e: "PUNCT SYM X CONJ NOUN ADP ADV DET PROPN NUM PRON",
		elem__string__PUNCT__SYM__X__CONJ__NOUN__PROPN__NUM__ADP__ADV___2e_2e_2e: "PUNCT SYM X CONJ NOUN PROPN NUM ADP ADV DET PRON",

		elem__string__SCONJ:               "SCONJ",
		elem__string__VERB:                "VERB",
		elem__string__aan:                 "aan",
		elem__string__advp:                "advp",
		elem__string__advp__ap:            "advp ap",
		elem__string__advp__ap__mwu__conj: "advp ap mwu conj",
		elem__string__aux:                 "aux",
		elem__string__ap:                  "ap",
		elem__string__app:                 "app",
		elem__string__bez:                 "bez",

		elem__string__blijken__hebben__hoeven__kunnen__moeten__moge_2e_2e_2e: "blijken hebben hoeven kunnen moeten mogen zijn zullen",

		elem__string__body:         "body",
		elem__string__bw:           "bw",
		elem__string__cleft:        "cleft",
		elem__string__cmp:          "cmp",
		elem__string__cnj:          "cnj",
		elem__string__cnj__dp__mwp: "cnj dp mwp",
		elem__string__conj:         "conj",
		elem__string__copula:       "copula",
		elem__string__cp:           "cp",

		elem__string__cp__sv1__smain__ppres__ppart__inf__ti__oti__du__w_2e_2e_2e: "cp sv1 smain ppres ppart inf ti oti du whq",
		elem__string__cp__sv1__smain__ssub__ppres__ppart__ti__oti__inf_2e_2e_2e:  "cp sv1 smain ssub ppres ppart ti oti inf du whq whrel rel",

		elem__string__cp__whrel__whq__whsub: "cp whrel whq whsub",
		elem__string__deeleigen:             "deeleigen",
		elem__string__det:                   "det",
		elem__string__detp:                  "detp",

		elem__string__dooien__gieten__hagelen__miezeren__misten__mo_2e_2e_2e: "dooien gieten hagelen miezeren misten motregenen onweren plenzen regenen sneeuwen stormen stortregenen ijzelen vriezen weerlichten winteren zomeren",

		elem__string__door:                          "door",
		elem__string__dp:                            "dp",
		elem__string__du:                            "ud",
		elem__string__du__smain:                     "du smain",
		elem__string__gen:                           "gen",
		elem__string__hd:                            "hd",
		elem__string__hd__body:                      "hd body",
		elem__string__hd__ld:                        "hd ld",
		elem__string__hd__ld__vc:                    "hd ld vc",
		elem__string__hd__mod:                       "hd mod",
		elem__string__hd__nucl__body:                "hd nucl body",
		elem__string__hd__su:                        "hd su",
		elem__string__hd__su__obj1:                  "hd su obj1",
		elem__string__hd__su__obj1__pc__predc__body: "hd su obj1 pc predc body",
		elem__string__hd__su__obj1__vc:              "hd su obj1 vc",
		elem__string__het:                           "het",
		elem__string__krijgen:                       "krijgen",
		elem__string__me:                            "me",
		elem__string__minmin:                        "--",
		elem__string__mod:                           "mod",
		elem__string__mod__pc__ld:                   "mod pc ld",
		elem__string__mwp:                           "mwp",
		elem__string__mwu:                           "mwu",
		elem__string__mwu__np__pp__ap__detp__smain:  "mwu np pp ap detp smain",
		elem__string__np:                            "np",
		elem__string__np__ap:                        "np ap",
		elem__string__np__conj:                      "np conj",
		elem__string__np__pp:                        "np pp",
		elem__string__nucl:                          "nucl",
		elem__string__obj1:                          "obj1",
		elem__string__obj1__se:                      "obj1 se",
		elem__string__obj1__se__me:                  "obj1 se me",
		elem__string__obj1__se__vc:                  "obj1 se vc",
		elem__string__obj1__vc__se__me:              "obj1 vc se me",
		elem__string__obj1__pobj1__se:               "obj1 pobj1 se",
		elem__string__obj1__pobj1__se__me:           "obj1 pobj1 se me",
		elem__string__obj2:                          "obj2",
		elem__string__ontbreken:                     "ontbreken",
		elem__string__oti:                           "oti",
		elem__string__passive:                       "passive",
		elem__string__pc:                            "pc",
		elem__string__pc__ld:                        "pc ld",
		elem__string__pobj1:                         "pobj1",
		elem__string__pp:                            "pp",
		elem__string__pp__detp__advp:                "pp detp advp",
		elem__string__pp__np:                        "pp np",
		elem__string__pp__np__conj__mwu:             "pp np conj mwu",
		elem__string__ppart:                         "ppart",
		elem__string__ppart__inf__ti:                "ppart inf ti",
		elem__string__pred:                          "pred",
		elem__string__predc:                         "predc",
		elem__string__rhd__whd:                      "rhd whd",
		elem__string__smain:                         "smain",
		elem__string__su:                            "su",
		elem__string__su__obj1:                      "su obj1",
		elem__string__su__obj1__predc__vc:           "su obj1 predc vc",
		elem__string__sup:                           "sup",

		elem__string__sv1__smain__ssub__inf__ppres__ppart__oti__ap__ad_2e_2e_2e: "sv1 smain ssub inf ppres ppart oti ap advp",

		elem__string__top:                      "top",
		elem__string__ti:                       "ti",
		elem__string__vc:                       "vc",
		elem__string__vc__predc:                "vc predc",
		elem__string__whd__rhd:                 "whd rhd",
		elem__string__whsub__ssub__ti__cp__oti: "whsub ssub ti cp oti",

		elem__string__zijn__lijken__blijken__blijven__schijnen__het_2e_2e_2e: "zijn lijken blijken blijven schijnen heten voorkomen worden dunken",

		elem__string__zijn__worden: "zijn worden",
	}

	elements = map[int][]interface{}{}
)

func init() {
	for key, val := range elemstrings {
		words := strings.Fields(val)
		elements[key] = make([]interface{}, len(words))
		for i, w := range words {
			elements[key][i] = w
		}
	}
}

type And struct {
	arg1 Doer
	arg2 Doer
}

func (d *And) Do(subdoc []interface{}, q *Context) []interface{} {
	for _, a := range []Doer{d.arg1, d.arg2} {
		if r := a.Do(subdoc, q); len(r) == 0 {
			return FALSE
		}
	}
	return TRUE
}

type Arg struct {
	arg1 Doer
	arg2 Doer
}

func (d *Arg) Do(subdoc []interface{}, q *Context) []interface{} {
	result := []interface{}{}
	for i, a := range []Doer{d.arg1, d.arg2} {
		if i == 0 || a != nil {
			result = append(result, flatten(a.Do(subdoc, q)))
		}
	}
	return result
}

type Cmp struct {
	ARG  int
	arg1 Doer
	arg2 Doer
}

func (d *Cmp) Do(subdoc []interface{}, q *Context) []interface{} {
	arg1 := d.arg1.Do(subdoc, q)
	arg2 := d.arg2.Do(subdoc, q)
	switch d.ARG {
	case cmp__lt: // <
		for _, a1 := range arg1 {
			for _, a2 := range arg2 {
				switch a1t := a1.(type) {
				case int:
					if a1t < a2.(int) {
						return TRUE
					}
				case string:
					if a1t < a2.(string) {
						return TRUE
					}
				default:
					panic(fmt.Sprintf("Missing case for type %T in %s", a1, q.filename))
				}
			}
		}
		return FALSE
	case cmp__gt: // >
		for _, a1 := range arg1 {
			for _, a2 := range arg2 {
				switch a1t := a1.(type) {
				case int:
					if a1t > a2.(int) {
						return TRUE
					}
				case string:
					if a1t > a2.(string) {
						return TRUE
					}
				default:
					panic(fmt.Sprintf("Missing case for type %T in %s", a1, q.filename))
				}
			}
		}
		return FALSE
	default:
		panic("Missing case in " + q.filename)
	}
}

type Collect struct {
	ARG  int
	arg1 Doer
	arg2 Doer
}

func (d *Collect) Do(subdoc []interface{}, q *Context) []interface{} {

	result1 := make([]interface{}, 0)
	for _, r := range d.arg1.Do(subdoc, q) {
		switch d.ARG {
		case collect__ancestors__node:
			result1 = r.(*NodeType).axAncestors
		case collect__attributes__begin:
			if i := r.(*NodeType).Begin; i >= 0 {
				result1 = append(result1, i)
			}
		case collect__attributes__cat:
			if i := r.(*NodeType).Cat; i != "" {
				result1 = append(result1, i)
			}
		case collect__attributes__end:
			if i := r.(*NodeType).End; i >= 0 {
				result1 = append(result1, i)
			}
		case collect__attributes__index:
			if i := r.(*NodeType).Index; i > 0 {
				result1 = append(result1, i)
			}
		case collect__attributes__lemma:
			if i := r.(*NodeType).Lemma; i != "" {
				result1 = append(result1, i)
			}
		case collect__attributes__naamval:
			if i := r.(*NodeType).Naamval; i != "" {
				result1 = append(result1, i)
			}
		case collect__attributes__pt:
			if i := r.(*NodeType).Pt; i != "" {
				result1 = append(result1, i)
			}
		case collect__attributes__rel:
			if i := r.(*NodeType).Rel; i != "" {
				result1 = append(result1, i)
			}
		case collect__attributes__sc:
			if i := r.(*NodeType).Sc; i != "" {
				result1 = append(result1, i)
			}
		case collect__attributes__spectype:
			if i := r.(*NodeType).Spectype; i != "" {
				result1 = append(result1, i)
			}
		case collect__attributes__ud_3apos:
			if i := r.(*NodeType).udPos; i != "" {
				result1 = append(result1, i)
			}
		case collect__attributes__vwtype:
			if i := r.(*NodeType).Vwtype; i != "" {
				result1 = append(result1, i)
			}
		case collect__attributes__word:
			if i := r.(*NodeType).Word; i != "" {
				result1 = append(result1, i)
			}
		case collect__child__node:
			result1 = append(result1, r.(Parent).Children()...)
		case collect__descendant__node:
			result1 = append(result1, r.(Parent).Descendants()...)
		case collect__descendant__or__self__type__node, collect__descendant__or__self__node:
			result1 = append(result1, r.(Parent).DescendantsOrSelf()...)
		case collect__parent__type__node:
			result1 = append(result1, r.(*NodeType).axParent...)
		case collect__self__all__node, collect__self__node:
			result1 = append(result1, r) // TODO: correct?
		default:
			panic("Missing case in " + q.filename)
		}
	}
	if d.arg2 == nil {
		return result1
	}

	result2 := make([]interface{}, 0)
	for _, r := range result1 {
		if len(d.arg2.Do([]interface{}{r}, q)) > 0 {
			result2 = append(result2, r)
		}
	}
	return result2
}

type Elem struct {
	ARG  int
	arg1 Doer
}

func (d *Elem) Do(subdoc []interface{}, q *Context) []interface{} {

	var r []interface{}
	var ok bool

	switch d.ARG {
	case elem__number__1:
		r = []interface{}{1}
	case elem__number__1000:
		r = []interface{}{1000}
	default:
		if r, ok = elements[d.ARG]; !ok {
			panic("Missing case in " + q.filename)
		}
	}

	/*
		TODO: arg1 negeren: klopt dat? zie bijvoorbeeld: foo[@a + 10 == @b]

		Waarvoor dient arg1 dan?

		SORT
		  COLLECT  'child' 'name' 'node' foo
		    NODE
		    PREDICATE
		      EQUAL =
		        PLUS +
		          COLLECT  'attributes' 'name' 'node' a
		            NODE
		          ELEM Object is a number : 1
		            COLLECT  'attributes' 'name' 'node' a
		              NODE
		        COLLECT  'attributes' 'name' 'node' b
		          NODE
	*/

	//	if d.arg1 == nil {
	return r
	//	}

	/*
		result := make([]interface{}, 0)
		for _, a := range d.arg1.Do(subdoc, q) {
			for _, s := range r {
				if a == s {
					result = append(result, a)
				}
			}
		}
		return result
	*/
}

type Equal struct {
	ARG  int
	arg1 Doer
	arg2 Doer
}

func (d *Equal) Do(subdoc []interface{}, q *Context) []interface{} {
	switch d.ARG {
	case equal__is:
		a1 := d.arg1.Do(subdoc, q)
		a2 := d.arg2.Do(subdoc, q)
		for _, aa1 := range a1 {
			for _, aa2 := range a2 {
				if aa1 == aa2 {
					return TRUE
				}
			}
		}
		return FALSE
	default:
		panic("Missing case in " + q.filename)
	}
}

type Filter struct {
	arg1 Doer
	arg2 Doer
}

func (d *Filter) Do(subdoc []interface{}, q *Context) []interface{} {
	result := make([]interface{}, 0)
	for _, r := range d.arg1.Do(subdoc, q) {
		if len(d.arg2.Do([]interface{}{r}, q)) > 0 {
			result = append(result, r)
		}
	}
	return result
}

type Function struct {
	ARG  int
	arg1 Doer
}

func (d *Function) Do(subdoc []interface{}, q *Context) []interface{} {

	r := d.arg1.Do(subdoc, q)

	switch d.ARG {
	case function__contains__2__args:
		for _, s1 := range r[1].([]interface{}) {
			for _, s0 := range r[0].([]interface{}) {
				if strings.Contains(s0.(string), s1.(string)) {
					return TRUE
				}
			}
		}
		return FALSE
	case function__count__1__args:
		return []interface{}{len(r[0].([]interface{}))}
		/*
			case function__deep__equal__2__args:
				// Nodes zijn pointers. Als de pointers niet gelijk zijn zijn de nodes niet gelijk,
				// want elke node heeft een ander ID. Het is dus niet nodig om elkaars children te vergelijken.
				r0 := r[0].([]interface{})
				r1 := r[1].([]interface{})
				if len(r0) != len(r1) {
					return FALSE
				}
				for i := range r0 {
					if r0[i] != r1[i] {
						return FALSE
					}
				}
				return TRUE
		*/
	case function__ends__with__2__args:
		for _, s1 := range r[1].([]interface{}) {
			for _, s0 := range r[0].([]interface{}) {
				if strings.HasSuffix(s0.(string), s1.(string)) {
					return TRUE
				}
			}
		}
		return FALSE
	case function__not__1__args:
		if len(r[0].([]interface{})) == 0 {
			return TRUE
		}
		return FALSE
	case function__starts__with__2__args:
		for _, s1 := range r[1].([]interface{}) {
			for _, s0 := range r[0].([]interface{}) {
				if strings.HasPrefix(s0.(string), s1.(string)) {
					return TRUE
				}
			}
		}
		return FALSE
	default:
		panic("Missing case in " + q.filename)
	}
}

type Node struct {
}

func (d *Node) Do(subdoc []interface{}, q *Context) []interface{} {
	return subdoc
}

type Or struct {
	arg1 Doer
	arg2 Doer
}

func (d *Or) Do(subdoc []interface{}, q *Context) []interface{} {
	for _, a := range []Doer{d.arg1, d.arg2} {
		if r := a.Do(subdoc, q); len(r) > 0 {
			return TRUE
		}
	}
	return FALSE
}

type Plus struct {
	ARG  int
	arg1 Doer
	arg2 Doer
}

func (d *Plus) Do(subdoc []interface{}, q *Context) []interface{} {
	switch d.ARG {
	case plus__plus:
		result := []interface{}{}
		a1 := d.arg1.Do(subdoc, q)
		a2 := d.arg2.Do(subdoc, q)
		for _, aa1 := range a1 {
			for _, aa2 := range a2 {
				result = append(result, aa1.(int)+aa2.(int))
			}
		}
		return result
	case plus__minus:
		result := []interface{}{}
		a1 := d.arg1.Do(subdoc, q)
		a2 := d.arg2.Do(subdoc, q)
		for _, aa1 := range a1 {
			for _, aa2 := range a2 {
				result = append(result, aa1.(int)-aa2.(int))
			}
		}
		return result
	default:
		panic("Missing case in " + q.filename)
	}
}

type Predicate struct {
	arg1 Doer
	arg2 Doer
}

func (d *Predicate) Do(subdoc []interface{}, q *Context) []interface{} {
	result := d.arg1.Do(subdoc, q)
	if d.arg2 == nil {
		return result
	}
	i := d.arg2.Do(subdoc, q)[0].(int) - 1
	if i >= 0 && i < len(result) {
		return []interface{}{result[i]}
	}
	return FALSE
}

type Root struct {
}

func (d *Root) Do(subdoc []interface{}, q *Context) []interface{} {
	return []interface{}{q.varroot}
}

type Sort struct {
	arg1 Doer
}

func (d *Sort) Do(subdoc []interface{}, q *Context) []interface{} {
	result := d.arg1.Do(subdoc, q)
	if len(result) < 2 {
		return result
	}
	switch result[0].(type) {
	case *NodeType:
		/*
			sort.Slice(result, func(i, j int) bool {
				ii := result[i].(*NodeType)
				jj := result[j].(*NodeType)
				if ii.Begin != jj.Begin {
					return ii.Begin < jj.Begin
				}
				if ii.End != jj.End {
					return ii.End < jj.End
				}
				if ii.Index != jj.Index {
					return ii.Index < jj.Index
				}
				return ii.Id < jj.Id
			})
			for i := 1; i < len(result); i++ {
				if result[i].(*NodeType).Id == result[i-1].(*NodeType).Id {
					result = append(result[:i], result[i+1:]...)
					i--
				}
			}
		*/
		// alleen ontdubbelen
		for i := 0; i < len(result); i++ {
			for j := i + 1; j < len(result); j++ {
				if result[i].(*NodeType) == result[j].(*NodeType) {
					result = append(result[:j], result[j+1:]...)
					j--
				}
			}
		}
	case string:
		/*
			sort.Slice(result, func(i, j int) bool {
				return result[i].(string) < result[j].(string)
			})
			for i := 1; i < len(result); i++ {
				if result[i].(string) == result[i-1].(string) {
					result = append(result[:i], result[i+1:]...)
					i--
				}
			}
		*/
		// alleen ontdubbelen
		for i := 0; i < len(result); i++ {
			for j := i + 1; j < len(result); j++ {
				if result[i].(string) == result[j].(string) {
					result = append(result[:j], result[j+1:]...)
					j--
				}
			}
		}
	case int:
		// alleen ontdubbelen
		for i := 0; i < len(result); i++ {
			for j := i + 1; j < len(result); j++ {
				if result[i].(int) == result[j].(int) {
					result = append(result[:j], result[j+1:]...)
					j--
				}
			}
		}
	default:
		panic(fmt.Sprintf("Missing case for type %T in %s", result[0], q.filename))
	}

	return result
}

/*
type Union struct {
	arg1 Doer
	arg2 Doer
}

func (d *Union) Do(subdoc []interface{}, q *Context) []interface{} {
	result := d.arg1.Do(subdoc, q)
	for _, i := range d.arg2.Do(subdoc, q) {
		found := false
		for _, j := range result {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			result = append(result, i)
		}
	}
	return result
}
*/

type Variable struct {
	ARG int
}

func (d *Variable) Do(subdoc []interface{}, q *Context) []interface{} {
	switch d.ARG {
	case variable__node:
		return q.varnode
	case variable__allnodes:
		return q.varallnodes
	case variable__gap:
		return q.vargap
	case variable__head:
		return q.varhead
	case variable__indexnodes:
		return q.varindexnodes
	case variable__tmp:
		return q.vartmp
	case variable__subj:
		return q.varsubj
	default:
		panic("Missing case in " + q.filename)
	}
}

type XPath struct {
	arg1 Doer
}

func (d *XPath) Do(q *Context) []interface{} {
	return d.arg1.Do([]interface{}{}, q)
}

////////////////////////////////////////////////////////////////

func list(i interface{}) []interface{} {
	switch ii := i.(type) {
	case []interface{}:
		return ii
	case []*NodeType:
		doc := []interface{}{}
		for _, n := range ii {
			doc = append(doc, n)
		}
		return doc
	default:
		return []interface{}{ii}
	}
}

func Test(node interface{}, q *Context, xpath *XPath) bool {
	q.varnode = list(node)
	return len(xpath.Do(q)) > 0
}

func Find(node interface{}, q *Context, xpath *XPath) []interface{} {
	q.varnode = list(node)
	return xpath.Do(q)
}

func SUTest(subj interface{}, q *Context, xpath *XPath) bool {
	q.varsubj = list(subj)
	return len(xpath.Do(q)) > 0
}

func SUFind(subj interface{}, q *Context, xpath *XPath) []interface{} {
	q.varsubj = list(subj)
	return xpath.Do(q)
}

func HGTest(head, gap interface{}, q *Context, xpath *XPath) bool {
	q.varhead = list(head)
	q.vargap = list(gap)
	return len(xpath.Do(q)) > 0
}

func HGFind(head, gap interface{}, q *Context, xpath *XPath) []interface{} {
	q.varhead = list(head)
	q.vargap = list(gap)
	return xpath.Do(q)
}

func flatten(aa []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, a := range aa {
		switch t := a.(type) {
		case []interface{}:
			result = append(result, flatten(t)...)
		default:
			result = append(result, a)
		}
	}
	return result
}

func (a *Alpino_ds) Children() []interface{} {
	return []interface{}{a.Node}
}

func (a *Alpino_ds) DescendantsOrSelf() []interface{} {
	r := []interface{}{a}
	r = append(r, a.Node.axDescendantsOrSelf...)
	return r
}

func (n *NodeType) Children() []interface{} {
	return n.axChildren
}

func (n *NodeType) Descendants() []interface{} {
	return n.axDescendants
}

func (n *NodeType) DescendantsOrSelf() []interface{} {
	return n.axDescendantsOrSelf
}
