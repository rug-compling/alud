package alud

import (
	//	"github.com/jbowtie/gokogiri"
	"github.com/jbowtie/gokogiri/xml"
	// "github.com/jbowtie/gokogiri/xpath"
)

type context struct {
	filename string
	sentence string
	sentid   string
	debugs   []string
	depth    int
	doc      *xml.XmlDocument
	root     *xml.ElementNode
	rootnode xml.Node
	ptnodes  []xml.Node
	idxnodes []xml.Node
	allnodes []xml.Node
}

type traceT struct {
	msg    string
	debugs []string
	trace  []traceType
}

type traceType struct {
	s    string
	node xml.Node
	head xml.Node
	gap  xml.Node
	subj xml.Node
}
