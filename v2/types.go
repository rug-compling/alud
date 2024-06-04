package alud

import (
	//	"github.com/jbowtie/gokogiri"
	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
)

type context struct {
	filename    string
	sentence    string
	sentid      string
	debugs      []string
	depth       int
	doc         *xml.XmlDocument
	root        *xml.ElementNode
	rootnode    xml.Node
	ptnodes     []*nodeType
	idxnodes    []*nodeType
	allnodes    []*nodeType
	varptnodes  []xml.Node
	varidxnodes []xml.Node
	varallnodes []xml.Node
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

type nodeType struct {
	node  xml.Node
	begin int
	end   int
	id    int

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
}

type varScope struct {
	vars map[string]interface{}
}

func (v *varScope) ResolveVariable(local, namespace string) interface{} {
	return v.vars[local]
}

func (v *varScope) IsFunctionRegistered(a, b string) bool {
	return false
}

func (v *varScope) ResolveFunction(a, b string) xpath.XPathFunction {
	return nil
}
