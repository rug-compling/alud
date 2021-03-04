package alud

import (
	"github.com/rug-compling/alpinods"

	"bytes"
	"encoding/xml"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// options can be or'ed as last argument to Ud()
const (
	OPT_DEBUG                  = 1 << iota // include debug messages in comments
	OPT_DUMMY_OUTPUT                       // include dummy output if parse fails
	OPT_NO_COMMENTS                        // don't include comments
	OPT_NO_DETOKENIZE                      // don't try to restore detokenized sentence
	OPT_NO_ENHANCED                        // skip enhanced dependencies
	OPT_NO_FIX_MISPLACED_HEADS             // don't fix misplaced heads in coordination
	OPT_NO_FIX_PUNCT                       // don't fix punctuation
	OPT_NO_METADATA                        // don't copy metadata to comments
	OPT_PANIC                              // panic on error (for development)
)

const (
	underscore = -1000 * (iota + 1)
	empty_head
	error_no_head
	error_no_value
)

var (
	noNode = &nodeType{
		NodeAttributes: alpinods.NodeAttributes{
			Begin: -1000,
			End:   -1000,
			ID:    -1,
		},
		udCopiedFrom:        -1000,
		Node:                []*nodeType{},
		axParent:            []interface{}{},
		axAncestors:         []interface{}{},
		axChildren:          []interface{}{},
		axDescendants:       []interface{}{},
		axDescendantsOrSelf: []interface{}{},
		udHeadPosition:      error_no_head,
		udEHeadPosition:     error_no_head,
	}
)

func init() {
	noNode.parent = noNode
}

// Version ID string
func VersionID() string {
	return "ALUD" + version
}

// Derive Universal Dependencies from parsed sentence in alpino_ds format.
//
// If sentid is "" it is derived from the filename.
func Ud(alpino_doc []byte, filename, sentid string, options int) (conllu string, err error) {
	conllu, _, err = ud(alpino_doc, filename, sentid, options)
	return
}

func ud(alpino_doc []byte, filename, sentid string, options int) (conllu string, q *context, err error) {
	if options&OPT_PANIC == 0 {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("%v", untrace(r))
				if options&OPT_DUMMY_OUTPUT == 0 {
					conllu = ""
				} else {
					conllu = dummyOutput(alpino_doc, filename, sentid, options, err)
				}
			}
		}()
	}
	conllu, q, err = udTry(alpino_doc, filename, sentid, options)
	if err != nil && options&OPT_DUMMY_OUTPUT != 0 {
		conllu = dummyOutput(alpino_doc, filename, sentid, options, err)
	}
	return // geen argumenten i.v.m. recover
}

func udTry(alpino_doc []byte, filename, sentid string, options int) (conllu string, q *context, err error) {

	var alpino alpino_ds
	err = xml.Unmarshal(alpino_doc, &alpino)
	if err != nil {
		return "", nil, err
	}

	if sentid != "" {
		alpino.Sentence.SentId = sentid
	} else if alpino.Sentence.SentId == "" {
		id := filepath.Base(filename)
		if strings.HasSuffix(id, ".xml") {
			id = id[:len(id)-4]
		}
		alpino.Sentence.SentId = id
	}

	var walk func(*nodeType)
	walk = func(node *nodeType) {
		node.Begin *= 1000
		node.End *= 1000
		node.ID *= 1000
		if node.Node == nil {
			node.Node = make([]*nodeType, 0)
		} else {
			for _, n := range node.Node {
				walk(n)
			}
		}
	}
	walk(alpino.Node)

	q = &context{
		alpino:   &alpino,
		filename: filename,
		sentence: alpino.Sentence.Sent,
		sentid:   alpino.Sentence.SentId,
		varroot:  []interface{}{alpino.Node},
		swapped:  [][2]*nodeType{},
	}

	inspect(q)

	if options&OPT_NO_FIX_MISPLACED_HEADS == 0 {
		fixMisplacedHeadsInCoordination(q)
	}
	addPosTags(q)
	addFeatures(q)
	addDependencyRelations(q)
	if options&OPT_NO_ENHANCED == 0 {
		enhancedDependencies(q)
	}
	if options&OPT_NO_FIX_PUNCT == 0 {
		fixpunct(q)
	}
	if options&OPT_NO_DETOKENIZE == 0 {
		untokenize(q)
	}
	check(q, options)
	return conll(q, options), q, nil
}

func inspect(q *context) {
	allnodes := make([]*nodeType, 0)
	varallnodes := make([]interface{}, 0)
	ptnodes := make([]*nodeType, 0)
	varindexnodes := make([]interface{}, 0)

	var walk func(*nodeType)
	walk = func(node *nodeType) {

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

func check(q *context, options int) {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "\t"+strings.TrimSpace(strings.Replace(conll(q, options), "\n", "\n\t", -1)), q))
		}
	}()

	root := -1
	items := make(map[string]int)
	for i, node := range q.ptnodes {
		if node.udHeadPosition == 0 && node.udRelation == "root" {
			if root >= 0 {
				panic("Multiple roots")
			}
			root = i
		}
		if node.udHeadPosition == 0 && node.udRelation != "root" {
			panic(fmt.Sprintf("Not a root: %s %q", number(node.End), node.Word))
		}
		if node.udHeadPosition != 0 && node.udRelation == "root" {
			panic(fmt.Sprintf("Invalid root: %s %q", number(node.End), node.Word))
		}
		items[number(node.End)] = i
	}
	if root < 0 {
		panic("Missing root")
	}

	// dit controleert standaard UD
	for _, node := range q.ptnodes {
		seen := make(map[int]bool)
		p := node.udHeadPosition
		if p < 0 {
			continue
		}
		for {
			if p == 0 {
				break
			}
			if seen[p] {
				panic(fmt.Sprintf("Loop in standard UD for word: %s %q", number(node.End), node.Word))
			}
			seen[p] = true
			i, ok := items[number(p)]
			if ok {
				p = q.ptnodes[i].udHeadPosition
			}
			if !ok || p == node.udHeadPosition {
				panic(fmt.Sprintf("Unreachable word in standard UD: %s %q", number(node.End), node.Word))
			}
		}
	}

	// dit controleert enhanced UD
	if options&OPT_NO_ENHANCED == 0 {
		for _, node := range q.ptnodes {
			found := false
			seen := make(map[string]bool)
			queue := []string{number(node.End)}
			for i := 0; i < len(queue); i++ {
				id := queue[i]
				if id == "0" {
					found = true
					break
				}
				if seen[id] {
					continue
				}
				seen[id] = true
				n := q.ptnodes[items[id]]
				for _, s := range strings.Split(n.udEnhanced, "|") {
					queue = append(queue, strings.Split(s, ":")[0])
				}
			}
			if !found {
				panic(fmt.Sprintf("Unreachable word in enhanced UD: %s %q", number(node.End), node.Word))
			}
		}
	}
}

func dummyOutput(alpino_doc []byte, filename, sentid string, options int, errin error) string {

	var buf bytes.Buffer

	if sentid == "" {
		sentid = filepath.Base(filename)
		if strings.HasSuffix(sentid, ".xml") {
			sentid = sentid[:len(sentid)-4]
		}
	}

	fmt.Fprintf(&buf, `# source = %s
# sent_id = %s
# auto = ALUD2.5.1-alpha009
# error = %s
`, filename, sentid, strings.Split(errin.Error(), "\n")[0])

	var alpino alpino_ds
	err := xml.Unmarshal(alpino_doc, &alpino)
	if err != nil {
		deps := "_"
		if options&OPT_NO_ENHANCED == 0 {
			deps = "0:root"
		}

		fmt.Fprintln(&buf, "# text = Fout\n1\tFout\tfout\tX\t_\t_\t0\troot\t"+deps+"\tError=Yes")
		return buf.String()
	}

	var walk func(*nodeType)
	walk = func(node *nodeType) {
		node.Begin *= 1000
		node.End *= 1000
		node.ID *= 1000
		if node.Node == nil {
			node.Node = make([]*nodeType, 0)
		} else {
			for _, n := range node.Node {
				walk(n)
			}
		}
	}
	walk(alpino.Node)

	q := &context{
		alpino:   &alpino,
		filename: filename,
		sentence: alpino.Sentence.Sent,
		sentid:   alpino.Sentence.SentId,
		varroot:  []interface{}{alpino.Node},
		swapped:  [][2]*nodeType{},
	}

	inspect(q)

	if options&OPT_NO_DETOKENIZE == 0 {
		untokenize(q)
	}

	head := "0"
	deprel := "root"
	deps := "_"

	u := func(s string) string {
		if s == "" {
			return "_"
		}
		return s
	}
	postag := func(s string) string {
		return strings.Join(strings.FieldsFunc(s, func(r rune) bool {
			return r == '(' || r == ')' || r == ','
		}), "|")
	}

	fmt.Fprintf(&buf, "# text = %s\n", q.sentence)
	for i, node := range q.ptnodes {
		if i == 1 {
			head = "1"
			deprel = "dep"
		}
		if i < 2 && options&OPT_NO_ENHANCED == 0 {
			deps = head + ":" + deprel
		}
		misc := "Error=Yes"
		if node.udNoSpaceAfter {
			misc = "SpaceAfter=No|Error=Yes"
		}
		fmt.Fprintf(&buf, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			number(node.End),       // ID
			node.Word,              // FORM
			node.Lemma,             // LEMMA
			"X",                    // UPOS
			u(postag(node.Postag)), // XPOS
			"_",                    // FEATS
			head,                   // HEAD
			deprel,                 // DEPREL
			deps,                   // DEPS
			misc)                   // MISC
	}

	return buf.String()
}
