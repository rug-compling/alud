package main

import (
	"fmt"
	"os"
	"sort"
)

// meest linkse node
func nLeft(nodes []interface{}) *NodeType {
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
		//if ii.End != jj.End {
		return ii.End < jj.End // ints
		//}
		//return ii.Id > jj.Id // ints, omgekeerd
	})
	return nodes[0].(*NodeType)
}

// meest linkse node als []interface{}, met lengte 0 of 1
func ifLeft(nodes []interface{}) []interface{} {
	n := nLeft(nodes)
	if n == noNode {
		return []interface{}{}
	}
	return []interface{}{n}
}

// eerste interface{} als []interface{}, met lengte 0 of 1
func if1(nodes []interface{}) []interface{} {
	if len(nodes) > 1 {
		return nodes[:1]
	}
	return nodes
}

// laatste interface{} als []interface{}, met lengte 0 of 1
func ifZ(nodes []interface{}) []interface{} {
	if len(nodes) > 0 {
		return nodes[len(nodes)-1:]
	}
	return []interface{}{}
}

// eerste node
func n1(nodes []interface{}) *NodeType {
	//return nLeft(nodes)

	if len(nodes) > 0 {
		return nodes[0].(*NodeType)
	}
	return noNode
}

// laatste node
func nZ(nodes []interface{}) *NodeType {
	if len(nodes) > 0 {
		return nodes[len(nodes)-1].(*NodeType)
	}
	return noNode
}

// eerste int
func i1(ii []interface{}) int {
	if len(ii) > 0 {
		return ii[0].(int)
	}
	return ERROR_NO_VALUE
}

// laatste int
func iZ(ii []interface{}) int {
	if l := len(ii); l > 0 {
		return ii[l-1].(int)
	}
	return ERROR_NO_VALUE
}

func depthCheck(q *Context, s string) bool {
	q.depth++
	if q.depth < 1000 {
		return false
	}
	q.warnings = append(q.warnings, "Recursion depth limit for "+s)
	fmt.Fprintf(os.Stderr, "WARNING: Recursion depth limit for %s in %s", s, q.filename)
	return true
}

/*
func dump(alpino *Alpino_ds) {
	b, err := xml.MarshalIndent(alpino, "", "  ")
	x(err)
	s := strings.Replace(string(b), "000", "", -1)
	fmt.Println("<?xml version=\"1.0\"?>\n" + s)
}
*/
