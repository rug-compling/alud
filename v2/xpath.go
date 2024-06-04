package alud

import (
	"fmt"
	"unsafe"

	"github.com/jbowtie/gokogiri/xml"
	pu "github.com/pebbe/util"
)

var x = pu.CheckErr

func test(q *context, expr string, args ...interface{}) bool {
	return len(find(q, expr, args...)) == 0
}

func find(q *context, expr string, args ...interface{}) []xml.Node {
	if len(args) == 0 {
		res, err := q.root.Search(expr)
		x(err)
		if err != nil {
			panic(err)
		}
		return res
	}
	vars := make(map[string]interface{})
	for i := 0; i < len(args); i += 2 {
		s := args[i].(string)
		switch t := args[i+1].(type) {
		case []xml.Node:
			vars[s] = xml.Nodeset(t).ToPointers()
		case xml.Node:
			vars[s] = []unsafe.Pointer{t.NodePtr()}
		default:
			panic(fmt.Errorf("Unknown type %T", t))
		}
	}
	res, err := q.root.SearchWithVariables(expr, &varScope{vars: vars})
	x(err)
	if err != nil {
		panic(err)
	}
	return res
}
