package main

import (
	//	"github.com/kr/pretty"

	"fmt"
	"sort"
	"strings"
)

type IndexType int

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
	collect__attributes__ud_3aERelation
	collect__attributes__ud_3aEHeadPosition
	collect__attributes__ud_3aHeadPosition
	collect__attributes__ud_3aPronType
	collect__attributes__ud_3aRelation
	collect__attributes__ud_3apos
	collect__attributes__vwtype
	collect__attributes__word
	collect__child__node
	collect__descendant__node
	collect__descendant__or__self__node
	collect__descendant__or__self__type__node
	collect__parent__type__node
	collect__parent__node
	collect__self__all__node
	collect__self__node
	equal__is
	function__contains__2__args
	function__count__1__args
	function__ends__with__2__args
	function__first__0__args
	function__last__0__args
	function__local__internal__head__position__1__args
	function__not__1__args
	function__starts__with__2__args
	plus__plus
	plus__minus
)

var (
	TRUE  = []interface{}{true}
	FALSE = []interface{}{}
)

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
			// TODO: waarom flatten?
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
			result1 = append(result1, r.(*NodeType).axAncestors...)
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
		case collect__attributes__ud_3aERelation:
			if i := r.(*NodeType).udERelation; i != "" {
				result1 = append(result1, i)
			}
		case collect__attributes__ud_3aEHeadPosition:
			if i := r.(*NodeType).udEHeadPosition; i > 0 {
				result1 = append(result1, i)
			}
		case collect__attributes__ud_3aHeadPosition:
			if i := r.(*NodeType).udHeadPosition; i > 0 {
				result1 = append(result1, i)
			}
		case collect__attributes__ud_3aPronType:
			if i := r.(*NodeType).udPronType; i != "" {
				result1 = append(result1, i)
			}
		case collect__attributes__ud_3aRelation:
			if i := r.(*NodeType).udRelation; i != "" {
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
		case collect__parent__type__node, collect__parent__node:
			result1 = append(result1, r.(*NodeType).axParent...)
		case collect__self__all__node, collect__self__node:
			result1 = append(result1, r) // TODO: correct?
		default:
			panic("Missing case in " + q.filename)
		}
	}

	// ontdubbelen
	if len(result1) == 0 {
		return result1
	}
	switch result1[0].(type) {
	case *NodeType:
		for i := 0; i < len(result1)-1; i++ {
			for j := i + 1; j < len(result1); j++ {
				if result1[i].(*NodeType) == result1[j].(*NodeType) {
					result1 = append(result1[:j], result1[j+1:]...)
					j--
				}
			}
		}
	case string:
		for i := 0; i < len(result1)-1; i++ {
			for j := i + 1; j < len(result1); j++ {
				if result1[i].(string) == result1[j].(string) {
					result1 = append(result1[:j], result1[j+1:]...)
					j--
				}
			}
		}
	case int:
		for i := 0; i < len(result1)-1; i++ {
			for j := i + 1; j < len(result1); j++ {
				if result1[i].(int) == result1[j].(int) {
					result1 = append(result1[:j], result1[j+1:]...)
					j--
				}
			}
		}
	default:
		panic(fmt.Sprintf("Missing case for type %T in %s", result1[0], q.filename))
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
	DATA []interface{}
	arg1 Doer
}

func (d *Elem) Do(subdoc []interface{}, q *Context) []interface{} {

	r := d.DATA

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
	for i, r := range d.arg1.Do(subdoc, q) {
		r2 := d.arg2.Do([]interface{}{r}, q)
		if len(r2) > 0 {
			if it, ok := r2[0].(IndexType); ok && len(r2) == 1 {
				// arg2 is een index
				ii := int(it)
				if ii == i+1 {
					result = append(result, r)
				}
			} else {
				result = append(result, r)
			}
		}
	}

	if len(result) == 0 {
		return result
	}
	switch result[0].(type) {
	case *NodeType:
		for i := 0; i < len(result)-1; i++ {
			for j := i + 1; j < len(result); j++ {
				if result[i].(*NodeType) == result[j].(*NodeType) {
					result = append(result[:j], result[j+1:]...)
					j--
				}
			}
		}
	case string:
		for i := 0; i < len(result)-1; i++ {
			for j := i + 1; j < len(result); j++ {
				if result[i].(string) == result[j].(string) {
					result = append(result[:j], result[j+1:]...)
					j--
				}
			}
		}
	case int:
		for i := 0; i < len(result)-1; i++ {
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

type Function struct {
	ARG  int
	arg1 Doer
}

func (d *Function) Do(subdoc []interface{}, q *Context) []interface{} {

	var r []interface{}
	if d.arg1 != nil {
		r = d.arg1.Do(subdoc, q)
	}

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
	case function__ends__with__2__args:
		for _, s1 := range r[1].([]interface{}) {
			for _, s0 := range r[0].([]interface{}) {
				if strings.HasSuffix(s0.(string), s1.(string)) {
					return TRUE
				}
			}
		}
		return FALSE
	case function__first__0__args:
		return []interface{}{IndexType(1)}
	case function__last__0__args:
		return []interface{}{IndexType(len(subdoc))}
	case function__local__internal__head__position__1__args:
		return []interface{}{internalHeadPosition(r[0].([]interface{}), q)}
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
	i := int(d.arg2.Do(subdoc, q)[0].(IndexType)) - 1 // TODO: always an index?
	if i >= 0 && i < len(result) {
		return []interface{}{result[i]}
	}
	return FALSE
}

type Root struct {
}

func (d *Root) Do(subdoc []interface{}, q *Context) []interface{} {
	return q.varroot
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
		sort.Slice(result, func(i, j int) bool {
			return result[i].(*NodeType).Id < result[j].(*NodeType).Id
		})
		for i := 1; i < len(result); i++ {
			if result[i].(*NodeType) == result[i-1].(*NodeType) {
				result = append(result[:i], result[i+1:]...)
				i--
			}
		}
	case string:
		sort.Slice(result, func(i, j int) bool {
			return result[i].(string) < result[j].(string)
		})
		for i := 1; i < len(result); i++ {
			if result[i].(string) == result[i-1].(string) {
				result = append(result[:i], result[i+1:]...)
				i--
			}
		}
	case int:
		sort.Slice(result, func(i, j int) bool {
			return result[i].(int) < result[j].(int)
		})
		for i := 1; i < len(result); i++ {
			if result[i].(int) == result[i-1].(int) {
				result = append(result[:i], result[i+1:]...)
				i--
			}
		}
	default:
		panic(fmt.Sprintf("Missing case for type %T in %s", result[0], q.filename))
	}
	return result
}

type Variable struct {
	VAR interface{}
}

func (d *Variable) Do(subdoc []interface{}, q *Context) []interface{} {
	switch t := d.VAR.(type) {
	case []interface{}:
		return t
	case *NodeType:
		return []interface{}{t}
	case []*NodeType:
		ii := make([]interface{}, len(t))
		for i, v := range t {
			ii[i] = v
		}
		return ii
	case int:
		return []interface{}{t}
	case string:
		return []interface{}{t}
	default:
		panic(fmt.Sprintf("Missing case for type %T in %s", t, q.filename))
	}
}

type XPath struct {
	arg1 Doer
}

func (d *XPath) Do(q *Context) []interface{} {
	return d.arg1.Do([]interface{}{}, q)
}

////////////////////////////////////////////////////////////////

func Test(q *Context, xpath *XPath) bool {
	return len(xpath.Do(q)) > 0
}

func Find(q *Context, xpath *XPath) []interface{} {
	return xpath.Do(q)
}

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
	return a.Node.axDescendantsOrSelf
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