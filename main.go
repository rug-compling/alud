package main

import (
	//	"github.com/kr/pretty"
	"github.com/pebbe/util"

	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	ERROR_NO_VALUE = -1000 * (iota + 1)
	ERROR_NO_EXTERNAL_HEAD
	ERROR_NO_INTERNAL_HEAD
	TODO
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
	vargap        []interface{}
	varhead       []interface{}
	varnode       []interface{}
	varroot       []interface{}
	varsubj       []interface{}
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
	// Aform        string      `xml:"aform,attr,omitempty"`
	Begin int `xml:"begin,attr"`
	// Buiging      string      `xml:"buiging,attr,omitempty"`
	// Case         string      `xml:"case,attr,omitempty"`
	Cat string `xml:"cat,attr,omitempty"`
	// Comparative  string      `xml:"comparative,attr,omitempty"`
	Conjtype string `xml:"conjtype,attr,omitempty"`
	// Def          string      `xml:"def,attr,omitempty"`
	// Dial         string      `xml:"dial,attr,omitempty"`
	// Dscmanual    string      `xml:"dscmanual,attr,omitempty"`
	// Dscsense     string      `xml:"dscsense,attr,omitempty"`
	End int `xml:"end,attr"`
	// Frame        string      `xml:"frame,attr,omitempty"`
	// Gen          string      `xml:"gen,attr,omitempty"`
	Genus string `xml:"genus,attr,omitempty"`
	Getal string `xml:"getal,attr,omitempty"`
	// GetalN       string      `xml:"getal-n,attr,omitempty"`
	Graad string `xml:"graad,attr,omitempty"`
	Id    int    `xml:"id,attr,omitempty"`
	// Iets         string      `xml:"iets,attr,omitempty"`
	Index int `xml:"index,attr,omitempty"`
	// Infl         string      `xml:"infl,attr,omitempty"`
	// Lcat         string      `xml:"lcat,attr,omitempty"`
	Lemma  string `xml:"lemma,attr,omitempty"`
	Lwtype string `xml:"lwtype,attr,omitempty"`
	// MwuRoot      string      `xml:"mwu_root,attr,omitempty"`
	// MwuSense     string      `xml:"mwu_sense,attr,omitempty"`
	Naamval string `xml:"naamval,attr,omitempty"`
	// Neclass      string      `xml:"neclass,attr,omitempty"`
	// Npagr        string      `xml:"npagr,attr,omitempty"`
	Ntype string `xml:"ntype,attr,omitempty"`
	// Num          string      `xml:"num,attr,omitempty"`
	Numtype string `xml:"numtype,attr,omitempty"`
	// OtherId string `xml:"other_id,attr,omitempty"`
	// Pb           string      `xml:"pb,attr,omitempty"`
	Pdtype string `xml:"pdtype,attr,omitempty"`
	// Per          string      `xml:"per,attr,omitempty"`
	// Personalized string      `xml:"personalized,attr,omitempty"`
	Persoon string `xml:"persoon,attr,omitempty"`
	// Pos          string      `xml:"pos,attr,omitempty"`
	// Positie      string      `xml:"positie,attr,omitempty"`
	Postag string `xml:"postag,attr,omitempty"`
	// Pron         string      `xml:"pron,attr,omitempty"`
	Pt     string `xml:"pt,attr,omitempty"`
	Pvagr  string `xml:"pvagr,attr,omitempty"`
	Pvtijd string `xml:"pvtijd,attr,omitempty"`
	// Refl         string      `xml:"refl,attr,omitempty"`
	Rel string `xml:"rel,attr,omitempty"`
	// Rnum         string      `xml:"rnum,attr,omitempty"`
	// Root         string      `xml:"root,attr,omitempty"`
	Sc string `xml:"sc,attr,omitempty"`
	// Sense        string      `xml:"sense,attr,omitempty"`
	// Special      string      `xml:"special,attr,omitempty"`
	Spectype string `xml:"spectype,attr,omitempty"`
	// Status       string      `xml:"status,attr,omitempty"`
	// Stype        string      `xml:"stype,attr,omitempty"`
	// Tense        string      `xml:"tense,attr,omitempty"`
	// Vform        string      `xml:"vform,attr,omitempty"`
	Vwtype string `xml:"vwtype,attr,omitempty"`
	// Vztype       string      `xml:"vztype,attr,omitempty"`
	// Wh           string      `xml:"wh,attr,omitempty"`
	// Wk           string      `xml:"wk,attr,omitempty"`
	Word                string      `xml:"word,attr,omitempty"`
	Wvorm               string      `xml:"wvorm,attr,omitempty"`
	Node                []*NodeType `xml:"node"`
	parent              *NodeType
	udAbbr              string
	udCase              string
	udCopiedFrom        string
	udDefinite          string
	udDegree            string
	udEnhanced          string
	udForeign           string
	udGender            string
	udHeadPosition      int
	udNumber            string
	udPerson            string
	udPos               string
	udPoss              string
	udPronType          string
	udReflex            string
	udRelation          string
	udTense             string
	udVerbForm          string
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
		Id:                  -1,
		Node:                []*NodeType{},
		axParent:            []interface{}{},
		axAncestors:         []interface{}{},
		axChildren:          []interface{}{},
		axDescendants:       []interface{}{},
		axDescendantsOrSelf: []interface{}{},
		udHeadPosition:      ERROR_NO_EXTERNAL_HEAD,
	}

	x = util.CheckErr
)

func init() {
	noNode.parent = noNode
}

func main() {

	for _, filename := range os.Args[1:] {
		doc, err := ioutil.ReadFile(filename)
		x(err)

		doDoc(doc, filename)
	}

	if !util.IsTerminal(os.Stdin) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			filename := scanner.Text()
			doc, err := ioutil.ReadFile(filename)
			x(err)

			doDoc(doc, filename)
		}
	}

}

func doDoc(doc []byte, filename string) {

	var alpino Alpino_ds
	err := xml.Unmarshal(doc, &alpino)
	if err != nil {
		fmt.Println(filename, err)
		return
	}

	if alpino.Sentence.SentId == "" {
		id := filepath.Base(filename)
		if strings.HasSuffix(id, ".xml") {
			id = id[:len(id)-4]
		}
		alpino.Sentence.SentId = id
	}

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
		filename: filename,
		sentence: alpino.Sentence.Sent,
		sentid:   alpino.Sentence.SentId,
		varroot:  []interface{}{alpino.Node},
		warnings: []string{},
	}

	inspect(q)

	fixMisplacedHeadsInCoordination(q)
	addPosTags(q)
	addFeatures(q)
	addDependencyRelations(q)
	enhancedDependencies(q)
	fixpunct(q)

	result := conll(q)
	fmt.Print(result)

}

func inspect(q *Context) {
	allnodes := make([]*NodeType, 0)
	varallnodes := make([]interface{}, 0)
	ptnodes := make([]*NodeType, 0)
	varindexnodes := make([]interface{}, 0)

	var walk func(*NodeType)
	walk = func(node *NodeType) {
		allnodes = append(allnodes, node)
		varallnodes = append(varallnodes, node)
		if node.Pt != "" || node.Word != "" { // bug in Alpino: missing pt
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
			node.axAncestors = append(node.axAncestors, node.parent.axAncestors...)
			node.axAncestors = append(node.axAncestors, node.parent)
			if node.axAncestors[0] != q.alpino.Node {
				// zou niet mogelijk moeten zijn
				panic("Missing ancestors in " + q.filename)
			}
		}
	}

	sort.Slice(ptnodes, func(i, j int) bool {
		return ptnodes[i].Begin < ptnodes[j].Begin
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

func number(n int) string {
	if n < 0 {
		switch n {
		case ERROR_NO_VALUE:
			return "ERROR_NO_VALUE"
		case ERROR_NO_EXTERNAL_HEAD:
			return "ERROR_NO_EXTERNAL_HEAD"
		case ERROR_NO_INTERNAL_HEAD:
			return "ERROR_NO_INTERNAL_HEAD"
		case TODO:
			return "TODO"
		default:
			panic("Missing case")
		}
	}
	i, j := n/1000, n%1000
	if j == 0 {
		return fmt.Sprint(i)
	}
	return fmt.Sprintf("%d.%d", i, j)
}

func conll(q *Context) string {

	var buf bytes.Buffer

	sort.Slice(q.ptnodes, func(i, j int) bool {
		return q.ptnodes[i].Begin < q.ptnodes[j].Begin
	})

	// TODO: ontdubbelen

	fmt.Fprintf(&buf, `# source = %s
# sent_id = %s
# text = %s
`,
		q.filename,
		q.sentid,
		q.sentence)

	for i, d := range q.debugs {
		fmt.Fprintf(&buf, "# debug%d = %s\n", i+1, d)
	}
	for i, w := range q.warnings {
		fmt.Fprintf(&buf, "# warning%d = %s\n", i+1, w)
	}

	u := func(s string) string {
		if s == "" {
			return "_"
		}
		return s
	}
	uc := func(s string) string {
		if s == "" {
			return "_"
		}
		return "CopiedFrom=" + s
	}

	postag := func(s string) string {
		return strings.Join(strings.FieldsFunc(s, func(r rune) bool {
			return r == '(' || r == ')' || r == ','
		}), "|")
	}

	for _, node := range q.ptnodes {
		fmt.Fprintf(&buf, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			number(node.End),            // ID
			node.Word,                   // FORM
			node.Lemma,                  // LEMMA
			u(node.udPos),               // UPOS
			u(postag(node.Postag)),      // XPOS
			u(featuresToString(node)),   // FEATS
			number(node.udHeadPosition), // HEAD
			u(node.udRelation),          // DEPREL
			u(node.udEnhanced),          // DEPS
			uc(node.udCopiedFrom))       // MISC
	}

	fmt.Fprintln(&buf)

	return buf.String()
}

func featuresToString(node *NodeType) string {
	features := make([]string, 0)
	for _, f := range [][2]string{
		[2]string{node.udAbbr, "Abbr"},
		[2]string{node.udCase, "Case"},
		[2]string{node.udDefinite, "Definite"},
		[2]string{node.udDegree, "Degree"},
		[2]string{node.udForeign, "Foreign"},
		[2]string{node.udGender, "Gender"},
		[2]string{node.udNumber, "Number"},
		[2]string{node.udPerson, "Person"},
		[2]string{node.udPoss, "Poss"},
		[2]string{node.udPronType, "PronType"},
		[2]string{node.udReflex, "Reflex"},
		[2]string{node.udTense, "Tense"},
		[2]string{node.udVerbForm, "VerbForm"},
	} {
		if f[0] != "" {
			features = append(features, f[1]+"="+f[0])
		}
	}
	return strings.Join(features, "|")
}

func dump(alpino *Alpino_ds) {
	b, err := xml.MarshalIndent(alpino, "", "  ")
	x(err)
	s := strings.Replace(string(b), "000", "", -1)
	fmt.Println("<?xml version=\"1.0\"?>\n" + s)
}
