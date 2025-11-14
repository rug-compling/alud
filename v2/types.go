package alud

import (
	"github.com/rug-compling/alpinods"

	"encoding/xml"
)

type context struct {
	alpino        *alpino_ds
	filename      string
	sentence      string
	sentid        string
	debugs        []string
	depth         int
	allnodes      []*nodeType
	ptnodes       []*nodeType
	varallnodes   []interface{}
	varindexnodes []interface{}
	varptnodes    []interface{}
	varroot       []interface{}
	swapped       [][2]*nodeType
}

type traceT struct {
	msg    string
	debugs []string
	trace  []traceType
}

type traceType struct {
	s    string
	node *nodeType
	head *nodeType
	gap  *nodeType
	subj *nodeType
}

type alpino_ds struct {
	XMLName  xml.Name           `xml:"alpino_ds"`
	Version  string             `xml:"version,attr,omitempty"`
	Metadata *alpinods.Metadata `xml:"metadata,omitempty"`
	Parser   *alpinods.Parser   `xml:"parser,omitempty"`
	Node     *nodeType          `xml:"node,omitempty"`
	Sentence *alpinods.Sentence `xml:"sentence,omitempty"`
	Comments *alpinods.Comments `xml:"comments,omitempty"`
	Root     []*deprelType      `xml:"root,omitempty"`
	Conllu   *alpinods.Conllu   `xml:"conllu,omitempty"`
}

type nodeType struct {
	alpinods.NodeAttributes
	Data []*alpinods.Data `xml:"data,omitempty"`
	Ud   *alpinods.Ud     `xml:"ud,omitempty"`
	Node []*nodeType      `xml:"node"`

	parent *nodeType

	// als je hier iets aan toevoegt, dan ook toevoegen in emptyheads-in.go in functie reconstructEmptyHead
	// DEP:DTD:1.8
	udAbbr           string
	udCase           string
	udCopiedFrom     int
	udDefinite       string
	udDegree         string
	udEHeadPosition  int
	udERelation      string
	udEnhanced       string
	udExtPos         string
	udFirstWordBegin int
	udForeign        string
	udGender         string
	udHeadPosition   int
	udMood           string
	udNoSpaceAfter   bool
	udNumber         string
	udOldState       *nodeType
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

type deprelType struct {
	XMLName xml.Name

	RecursionLimit string `xml:"recursion_limit,attr,omitempty"`
	recursion      []string

	Ud    string `xml:"ud,attr,omitempty"`
	ID    string `xml:"id,attr,omitempty"`
	Form  string `xml:"form,attr,omitempty"`
	Lemma string `xml:"lemma,attr,omitempty"`
	Upos  string `xml:"upos,attr,omitempty"`
	alpinods.Feats
	Head      string `xml:"head,attr,omitempty"`
	Deprel    string `xml:"deprel,attr,omitempty"`
	DeprelAux string `xml:"deprel_aux,attr,omitempty"`

	alpinods.DeprelAttributes

	Deps []*deprelType `xml:",omitempty"`
}

type nodeDepType struct {
	node *nodeType
	dep  *alpinods.Dep
}
