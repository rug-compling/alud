// +build compile

package main

import (
	"github.com/pebbe/util"

	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Line struct {
	lvl int
	s   string
}

var (
	x = util.CheckErr

	reTest    = regexp.MustCompile("(?s:(?:TEST|FIND)\\(.*?`.*?`\\))")
	reComment = regexp.MustCompile("(?s:\\(:.*?:\\))")
	reRemove  = regexp.MustCompile("[^a-zA-Z0-9]")
	reSet     = regexp.MustCompile("(?s:=\\s*\\(\\s*(\".*?\")\\s*\\))")

	program string

	hash = make(map[string]string)
)

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	x(err)
	program = strings.Replace(string(b), "// +build ignore", "", 1)

	program = reTest.ReplaceAllStringFunc(program, compile)

	fmt.Println("//\n// GENERATED FILE -- DO NOT EDIT\n//\n")

	fmt.Print(program)
}

func compile(s string) string {

	start := strings.Index(s, "`") + 1
	stop := strings.LastIndex(s, "`")
	prefix := strings.Replace(s[:start-1], "TEST", "test", 1)
	prefix = strings.Replace(prefix, "FIND", "find", 1)
	suffix := s[stop+1:]
	s = s[start:stop]

	s2 := reComment.ReplaceAllLiteralString(s, "")

	// [1] -> [first()]
	s2 = strings.Replace(s2, "[1]", "[first()]", -1)

	// =("aap","noot","mies")  ->  ="hashcode"
	s2 = reSet.ReplaceAllStringFunc(s2, func(s string) string {
		mm := reSet.FindStringSubmatch(s)
		h := fmt.Sprintf("%x", md5.Sum([]byte(mm[1])))
		hash[h] = mm[1]
		return `="` + h + `"`
	})

	cmd := exec.Command("testXPath", "--tree", s2)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	x(cmd.Run())
	so := stdout.String()
	se := strings.Replace(stderr.String(), "xmlXPathCompiledEval: evaluation failed\n", "", 1)
	se = strings.Replace(se, "XPath error : Undefined variable\n", "", 1)

	if se != "" {
		fmt.Fprintln(os.Stderr, "ERROR:", s2)
		return prefix + " /* " + s + " */ " + fmt.Sprintf("ERROR(%q)", se) + suffix
	}

	return prefix + " /* " + s + " */ " + parse(so) + suffix
}

func parse(s string) string {

	out := make([]string, 0)

	lines := make([]Line, 0)
	for _, l := range strings.Split(s, "\n") {
		if len(l) == 0 || l[0] != ' ' {
			continue
		}
		lvl := (len(l) - len(strings.TrimLeft(l, " "))) / 2
		lines = append(lines, Line{
			s:   strings.TrimSpace(l),
			lvl: lvl,
		})
	}

	if len(lines) == 0 {
		return "&xPath{\"COMPILE ERROR\"}"
	}

	out = append(out, "&xPath{")
	lvl := 0
	prev := make([]int, 1)
	for _, line := range lines {
		for lvl > line.lvl {
			out = append(out, "},")
			lvl--
			prev = prev[:len(prev)-1]
		}
		if lvl == line.lvl {
			out = append(out, "},")
		} else {
			lvl = line.lvl
		}
		words := strings.Fields(line.s)
		cmd := strings.ToLower(words[0])
		str := "d" + strings.Title(cmd)
		cmd0 := cmd
		if len(prev) <= lvl {
			prev = append(prev, 1)
			cmd = "arg1"
		} else {
			prev[lvl] += 1
			cmd = fmt.Sprint("arg", prev[lvl])
		}
		out = append(out, fmt.Sprintf("%s: &%s{", cmd, str))
		if cmd0 == "variable" {
			out = append(out, fmt.Sprintf("VAR: %s,", words[1]))
		} else if cmd0 == "elem" {
			if words[4] == "string" {
				data, ok := hash[words[6]]
				if !ok {
					data = `"` + words[6] + `"`
				}
				out = append(out, fmt.Sprintf("DATA: []interface{}{%s},", data))
			} else if words[4] == "number" {
				out = append(out, fmt.Sprintf("DATA: []interface{}{%s},", words[6]))
			} else {
				panic("Unknown elem type: " + strings.Join(words, " "))
			}
		} else {
			if len(words) > 1 {
				args := strings.Join(words[1:], " ")
				args = strings.Replace(args, "'", "", -1)
				args = strings.Replace(args, " :", "", -1)
				args = strings.Replace(args, "<", "lt", -1)
				args = strings.Replace(args, ">", "gt", -1)
				args = strings.Replace(args, "(", " ", -1)
				args = strings.Replace(args, ")", "", -1)
				args = strings.Replace(args, "=", "is", -1)
				args = strings.Replace(args, "--", "minmin", -1)
				args = strings.Replace(args, "deep-equal", "deep equal", -1)
				args = strings.Replace(args, "-with", " with", -1)
				args = strings.Replace(args, "-or-self", " or self", -1)
				args = strings.Replace(args, "+", "plus", -1)
				args = strings.Replace(args, " name node", "", -1)
				args = strings.Replace(args, "local:internal_head_position", "local internal head position", -1)
				args = strings.Replace(reRemove.ReplaceAllStringFunc(args, func(s string) string {
					return fmt.Sprintf("_%02x", s[0])
				}), "_20", "__", -1)
				if args == "_2d" {
					args = "minus"
				}
				out = append(out, fmt.Sprintf("ARG: %s__%s,", cmd0, args))
			}
		}
	}
	for lvl > 0 {
		out = append(out, "},")
		lvl--
	}
	out = append(out, "}")

	return strings.Join(out, "\n")
}
