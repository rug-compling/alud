// +build ignore

package main

import (
	//	"github.com/kr/pretty"
	"github.com/pebbe/util"

	"encoding/xml"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const (
	ERROR_NO_VALUE = -1000 * (iota + 1)
	ERROR_NO_HEAD_FOUND
	ERROR_NO_EXTERNAL_HEAD
	ERROR_NO_INTERNAL_HEAD_IN_GAPPED_CONSTITUENT
	ERROR_NO_INTERNAL_HEAD
	UNDERSCORE
	TODO
	empty_head
)

var (
	x = util.CheckErr
)

type Context struct {
	alpino        *Alpino_ds
	filename      string
	sentence      string
	sentid        string
	debugs        []string
	warnings      []string
	depth         int
	allnodes      []*NodeType
	ptnodes       []*NodeType
	varallnodes   []interface{}
	varindexnodes []interface{}
	varptnodes    []interface{}
	varroot       []interface{}
}

type Alpino_ds struct {
	XMLName  xml.Name  `xml:"alpino_ds"`
	Node     *NodeType `xml:"node,omitempty"`
	Sentence *SentType `xml:"sentence,omitempty"`
}

type SentType struct {
	Sent   string `xml:",chardata"`
	SentId string `xml:"sentid,attr,omitempty"`
}

type NodeType struct {
	Begin    int         `xml:"begin,attr"`
	Cat      string      `xml:"cat,attr,omitempty"`
	Conjtype string      `xml:"conjtype,attr,omitempty"`
	End      int         `xml:"end,attr"`
	Genus    string      `xml:"genus,attr,omitempty"`
	Getal    string      `xml:"getal,attr,omitempty"`
	Graad    string      `xml:"graad,attr,omitempty"`
	Id       int         `xml:"id,attr,omitempty"`
	Index    int         `xml:"index,attr,omitempty"`
	Lemma    string      `xml:"lemma,attr,omitempty"`
	Lwtype   string      `xml:"lwtype,attr,omitempty"`
	Naamval  string      `xml:"naamval,attr,omitempty"`
	Ntype    string      `xml:"ntype,attr,omitempty"`
	Numtype  string      `xml:"numtype,attr,omitempty"`
	Pdtype   string      `xml:"pdtype,attr,omitempty"`
	Persoon  string      `xml:"persoon,attr,omitempty"`
	Postag   string      `xml:"postag,attr,omitempty"`
	Pt       string      `xml:"pt,attr,omitempty"`
	Pvagr    string      `xml:"pvagr,attr,omitempty"`
	Pvtijd   string      `xml:"pvtijd,attr,omitempty"`
	Rel      string      `xml:"rel,attr,omitempty"`
	Sc       string      `xml:"sc,attr,omitempty"`
	Spectype string      `xml:"spectype,attr,omitempty"`
	Vwtype   string      `xml:"vwtype,attr,omitempty"`
	Word     string      `xml:"word,attr,omitempty"`
	Wvorm    string      `xml:"wvorm,attr,omitempty"`
	Node     []*NodeType `xml:"node"`
	parent   *NodeType

	// als je hier iets aan toevoegt, dan ook toevoegen in emptyheads-in.go in functie reconstructEmptyHead
	udAbbr           string
	udCase           string
	udCopiedFrom     int
	udDefinite       string
	udDegree         string
	udEHeadPosition  int
	udERelation      string
	udEnhanced       string
	udFirstWordBegin int
	udForeign        string
	udGender         string
	udHeadPosition   int
	udNumber         string
	udPerson         string
	udPos            string
	udPoss           string
	udPronType       string
	udReflex         string
	udRelation       string
	udTense          string
	udVerbForm       string

	axParent            []interface{}
	axAncestors         []interface{}
	axChildren          []interface{}
	axDescendants       []interface{}
	axDescendantsOrSelf []interface{}
}

var (
	noNode = &NodeType{
		Begin:               -1000,
		End:                 -1000,
		udCopiedFrom:        -1000,
		Id:                  -1,
		Node:                []*NodeType{},
		axParent:            []interface{}{},
		axAncestors:         []interface{}{},
		axChildren:          []interface{}{},
		axDescendants:       []interface{}{},
		axDescendantsOrSelf: []interface{}{},
		udHeadPosition:      ERROR_NO_EXTERNAL_HEAD,
		udEHeadPosition:     ERROR_NO_EXTERNAL_HEAD,
	}
)

func init() {
	noNode.parent = noNode
}

func main() {

	b, err := ioutil.ReadFile("file.xml")
	x(err)

	var alpino Alpino_ds
	x(xml.Unmarshal(b, &alpino))

	var walk func(*NodeType)
	walk = func(node *NodeType) {
		node.Begin *= 1000
		node.End *= 1000
		if node.Node == nil {
			node.Node = make([]*NodeType, 0)
		} else {
			for _, n := range node.Node {
				walk(n)
			}
		}
	}
	walk(alpino.Node)

	q := &Context{
		alpino:   &alpino,
		varroot:  []interface{}{alpino.Node},
		warnings: []string{},
	}

	inspect(q)

	node := alpino.Node

	fmt.Println("Test 1")
	for _, n := range FIND(q, `$node/node/node[1]`) {
		fmt.Println("id=", n.(*NodeType).Id)
	}

	fmt.Println("Test 2")
	for _, n := range FIND(q, `($node/node/node)[1]`) {
		fmt.Println("id=", n.(*NodeType).Id)
	}

}

func inspect(q *Context) {
	allnodes := make([]*NodeType, 0)
	varallnodes := make([]interface{}, 0)
	ptnodes := make([]*NodeType, 0)
	varindexnodes := make([]interface{}, 0)

	var walk func(*NodeType)
	walk = func(node *NodeType) {

		// bug in Alpino: missing pt
		if node.Word != "" && node.Pt == "" {
			node.Pt = strings.ToLower(strings.Split(node.Postag, "(")[0])
			if node.Pt == "" {
				node.Pt = "na"
			}
		}

		allnodes = append(allnodes, node)
		varallnodes = append(varallnodes, node)
		if node.Pt != "" {
			ptnodes = append(ptnodes, node)
		}
		if node.Index > 0 {
			varindexnodes = append(varindexnodes, node)
		}
		for _, n := range node.Node {
			n.parent = node
			n.axParent = []interface{}{node}
			walk(n)
		}
		node.axChildren = make([]interface{}, 0)
		node.axDescendants = make([]interface{}, 0)
		node.axDescendantsOrSelf = make([]interface{}, 1)
		node.axDescendantsOrSelf[0] = node
		for _, n := range node.Node {
			node.axChildren = append(node.axChildren, n)
			node.axDescendants = append(node.axDescendants, n)
			node.axDescendants = append(node.axDescendants, n.axDescendants...)
			node.axDescendantsOrSelf = append(node.axDescendantsOrSelf, n.axDescendantsOrSelf...) // niet n
		}
	}
	walk(q.alpino.Node)
	q.alpino.Node.parent = noNode
	q.alpino.Node.axParent = []interface{}{}

	for _, node := range allnodes {
		node.axAncestors = make([]interface{}, 0)
		if node != q.alpino.Node {
			node.axAncestors = append(node.axAncestors, node.parent)
			node.axAncestors = append(node.axAncestors, node.parent.axAncestors...)
			if node.axAncestors[len(node.axAncestors)-1] != q.alpino.Node {
				// zou niet mogelijk moeten zijn
				panic("Missing ancestors in " + q.filename)
			}
		}
	}

	sort.Slice(ptnodes, func(i, j int) bool {
		return ptnodes[i].End < ptnodes[j].End
	})
	varptnodes := make([]interface{}, len(ptnodes))
	for i, node := range ptnodes {
		varptnodes[i] = node
	}

	q.allnodes = allnodes
	q.varallnodes = varallnodes
	q.varindexnodes = varindexnodes
	q.ptnodes = ptnodes
	q.varptnodes = varptnodes

}

func internalHeadPosition(node []interface{}, q *Context) int {
	return ERROR_NO_INTERNAL_HEAD
}
