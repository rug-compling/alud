package alud

import (
	"bytes"
	"fmt"
	"strings"
)

func conll(q *context, options int) string {

	var buf bytes.Buffer

	if options&OPT_NO_COMMENTS == 0 {
		fmt.Fprintf(&buf, `# source = %s
# sent_id = %s
# text = %s
# auto = %s
`,
			q.filename,
			strings.Replace(q.sentid, "/", "\\", -1), // het teken / is gereserveerd
			q.sentence,
			VersionID())

		if options&OPT_DEBUG != 0 {
			for _, d := range q.debugs {
				fmt.Fprintf(&buf, "# debug = %s\n", d)
			}
		}

		if options&OPT_NO_METADATA == 0 {
			if q.alpino.Metadata != nil && q.alpino.Metadata.Meta != nil {
				for _, m := range q.alpino.Metadata.Meta {
					fmt.Fprintf(&buf, "# meta_%s = %s\n", m.Name, m.Value)
				}
			}
		}
	}

	u := func(s string) string {
		if s == "" {
			return "_"
		}
		return s
	}
	misc := func(node *nodeType) string {
		ss := make([]string, 0, 2)
		if node.udCopiedFrom > 0 {
			ss = append(ss, "CopiedFrom="+number(node.udCopiedFrom))
		}
		if node.udNoSpaceAfter {
			ss = append(ss, "SpaceAfter=No")
		}
		return strings.Join(ss, "|")
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
			u(misc(node)))               // MISC
	}

	fmt.Fprintln(&buf)

	return buf.String()
}

func featuresToString(node *nodeType) string {
	features := make([]string, 0)
	for _, f := range [][2]string{
		{node.udAbbr, "Abbr"},
		{node.udCase, "Case"},
		{node.udDefinite, "Definite"},
		{node.udDegree, "Degree"},
		{node.udForeign, "Foreign"},
		{node.udGender, "Gender"},
		{node.udNumber, "Number"},
		{node.udPerson, "Person"},
		{node.udPoss, "Poss"},
		{node.udPronType, "PronType"},
		{node.udReflex, "Reflex"},
		{node.udTense, "Tense"},
		{node.udVerbForm, "VerbForm"},
	} {
		if f[0] != "" {
			features = append(features, f[1]+"="+f[0])
		}
	}
	return strings.Join(features, "|")
}

func number(n int) string {
	if n < 0 {
		switch n {
		case underscore:
			return "_"
		case empty_head:
			return "empty head"
		case error_no_head:
			panic("No head")
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
