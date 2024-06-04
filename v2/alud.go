package alud

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jbowtie/gokogiri"
	"github.com/jbowtie/gokogiri/xml"
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
				// TODO err = fmt.Errorf("%v", untrace(r))
				if options&OPT_DUMMY_OUTPUT == 0 {
					conllu = ""
				} else {
					// TODO				conllu = dummyOutput(alpino_doc, filename, sentid, options, err)
				}
			}
		}()
	}
	conllu, q, err = udTry(alpino_doc, filename, sentid, options)
	if err != nil && options&OPT_DUMMY_OUTPUT != 0 {
		// TODO  	conllu = dummyOutput(alpino_doc, filename, sentid, options, err)
	}
	return // geen argumenten i.v.m. recover
}

func udTry(alpino_doc []byte, filename, sentid string, options int) (conllu string, q *context, err error) {
	if sentid == "" && filename != "" {
		id := filepath.Base(filename)
		if strings.HasSuffix(id, ".xml") {
			id = id[:len(id)-4]
		}
		sentid = id
	}

	q = &context{
		filename: filename,
		sentid:   sentid,
	}

	q.doc, err = gokogiri.ParseXml(alpino_doc)
	if err != nil {
		return
	}
	defer q.doc.Free()

	q.root = q.doc.Root()
	if sentence, err := q.root.Search("./sentence"); err == nil {
		if len(sentence) > 0 {
			q.sentence = sentence[0].Content()
			sentid = sentence[0].Attr("sentid")
			if sentid != "" {
				q.sentid = sentid
			}
		}
	}

	roots, err := q.root.Search("./node")
	if err != nil {
		return
	}
	if len(roots) != 1 {
		err = fmt.Errorf("len(roots) == %d", len(roots))
		return
	}
	q.rootnode = roots[0]

	nodes, err := q.root.Search(`.//node`)
	if err != nil {
		return
	}
	q.allnodes = make([]*nodeType, len(nodes))
	q.idxnodes = make([]*nodeType, 0)
	q.ptnodes = make([]*nodeType, 0)
	q.varallnodes = make([]xml.Node, len(nodes))
	q.varidxnodes = make([]xml.Node, 0)
	q.varptnodes = make([]xml.Node, 0)
	for i, node := range nodes {
		begin, _ := strconv.Atoi(node.Attr("begin"))
		end, _ := strconv.Atoi(node.Attr("end"))
		id, _ := strconv.Atoi(node.Attr("id"))
		node.SetAttr("begin", fmt.Sprintf("%04d000", begin))
		node.SetAttr("end", fmt.Sprintf("%04d000", end))
		node.SetAttr("id", fmt.Sprintf("%04d000", id))
		q.allnodes[i] = &nodeType{
			node:  node,
			begin: begin * 1000,
			end:   end * 1000,
			id:    id * 1000,
		}
		q.varallnodes[i] = node
		pt := node.Attr("pt")
		if pt != "" {
			q.ptnodes = append(q.ptnodes, q.allnodes[i])
			q.varptnodes = append(q.varptnodes, node)
		}
		if node.Attr("index") != "" && (pt != "" || node.Attr("cat") != "") {
			q.idxnodes = append(q.idxnodes, q.allnodes[i])
			q.varidxnodes = append(q.varidxnodes, node)
		}
	}

	if options&OPT_NO_FIX_MISPLACED_HEADS == 0 {
		// TODO fixMisplacedHeadsInCoordination(q)
	}
	addPosTags(q)
	// TODO addFeatures(q)
	// TODO addDependencyRelations(q)
	if options&OPT_NO_ENHANCED == 0 {
		// TODO enhancedDependencies(q)
	}
	if options&OPT_NO_FIX_PUNCT == 0 {
		// TODO fixpunct(q, options&OPT_NO_ENHANCED == 0)
	}
	if options&OPT_NO_DETOKENIZE == 0 {
		// TODO untokenize(q)
	}
	// TODO check(q, options)
	return conll(q, options), q, nil
}

func check(q *context, options int) {
	defer func() {
		if r := recover(); r != nil {
			// TODO
			// panic(trace(r, "\t"+strings.TrimSpace(strings.Replace(conll(q, options), "\n", "\n\t", -1)), q))
			panic("PANIC")
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
			panic(fmt.Sprintf("Not a root: %s %q", number(node.end), node.node.Attr("word")))
		}
		if node.udHeadPosition != 0 && node.udRelation == "root" {
			panic(fmt.Sprintf("Invalid root: %s %q", number(node.end), node.node.Attr("word")))
		}
		items[number(node.end)] = i
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
				panic(fmt.Sprintf("Loop in standard UD for word: %s %q", number(node.end), node.node.Attr("word")))
			}
			seen[p] = true
			i, ok := items[number(p)]
			if ok {
				p = q.ptnodes[i].udHeadPosition
			}
			if !ok || p == node.udHeadPosition {
				panic(fmt.Sprintf("Unreachable word in standard UD: %s %q", number(node.end), node.node.Attr("word")))
			}
		}
	}

	// dit controleert enhanced UD
	if options&OPT_NO_ENHANCED == 0 {
		for _, node := range q.ptnodes {
			found := false
			seen := make(map[string]bool)
			queue := []string{number(node.end)}
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
				panic(fmt.Sprintf("Unreachable word in enhanced UD: %s %q", number(node.end), node.node.Attr("word")))
			}
		}
	}
}

/*
func dummyOutput(alpino_doc []byte, filename, sentid string, options int, errin error) string {
	var alpino alpino_ds
	err := xml.Unmarshal(alpino_doc, &alpino)

	if sentid == "" && err == nil {
		sentid = alpino.Sentence.SentId
	}
	if sentid == "" {
		sentid = filepath.Base(filename)
		if strings.HasSuffix(sentid, ".xml") {
			sentid = sentid[:len(sentid)-4]
		}
	}

	var buf bytes.Buffer

	fmt.Fprintf(&buf, `# source = %s
# sent_id = %s
# error = %s
# auto = %s
`, filename, sentid, strings.Split(errin.Error(), "\n")[0], VersionID())

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

*/
