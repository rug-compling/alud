package main

import (
	"fmt"
	"os"
	"sort"
)

func left(nodes []interface{}) []interface{} {
	n := leftmost(nodes)
	if n == noNode {
		return []interface{}{}
	}
	return []interface{}{n}
}

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
		return ii.Id > jj.Id // ints, omgekeerd
	})
	return nodes[0].(*NodeType)
}

func first(nodes []interface{}) []interface{} {
	if len(nodes) > 1 {
		return nodes[:1]
	}
	return nodes
}

func firstnode(nodes []interface{}) *NodeType {
	//return leftmost(nodes)

	if len(nodes) > 0 {
		return nodes[0].(*NodeType)
	}
	return noNode
}

func last(nodes []interface{}) []interface{} {
	if len(nodes) > 0 {
		return nodes[len(nodes)-1:]
	}
	return []interface{}{}
}

func lastnode(nodes []interface{}) *NodeType {
	if len(nodes) > 0 {
		return nodes[len(nodes)-1].(*NodeType)
	}
	return noNode
}

func firstint(ii []interface{}) int {
	if len(ii) > 0 {
		return ii[0].(int)
	}
	return ERROR_NO_VALUE
}

func depthCheck(q *Context, s string) bool {
	q.depth++
	if q.depth < 1000 {
		return false
	}
	q.warnings = append(q.warnings, "Recursion depth limit for "+s)
	fmt.Fprintln(os.Stderr, "WARNING: Recursion depth limit for %s in %s", s, q.filename)
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
