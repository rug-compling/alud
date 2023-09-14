package main

import (
	"github.com/rug-compling/alud/v2"
	"github.com/rug-compling/alud/v2/internal/util"

	"github.com/pebbe/compactcorpus"
	"github.com/pebbe/dbxml"
	"github.com/rug-compling/alpinods"

	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	opt_a = flag.Bool("a", false, "output in alpino")
	opt_c = flag.Bool("c", false, "don't include comments")
	opt_d = flag.Bool("d", false, "include debug messages in comments")
	opt_e = flag.Bool("e", false, "skip enhanced dependencies")
	opt_f = flag.Bool("f", false, "don't fix punctuation")
	opt_i = flag.String("i", "", "regexp for filtering filename as sentence id")
	opt_m = flag.Bool("m", false, "don't fix mixplaced heads in coordination")
	opt_M = flag.Bool("M", false, "don't copy metadata to comments")
	opt_p = flag.Bool("p", false, "panic on error (for development)")
	opt_s = flag.Bool("s", false, "include short error comments if parse fails")
	opt_t = flag.Bool("t", false, "don't try to restore detokenized sentence")
	opt_v = flag.Bool("v", false, "print version and exit")
	opt_x = flag.Bool("x", false, "include dummy output if parse fails")

	multi = false

	x = util.CheckErr

	reSentID = regexp.MustCompile(`<sentence.*?sentid="(.*?)".*</sentence>`)
)

type Collection struct {
	XMLName xml.Name `xml:"collection"`
	Doc     []Doc    `xml:"doc"`
}

type Doc struct {
	Href string `xml:"href,attr"`
}

func usage() {
	p := filepath.Base(os.Args[0])
	fmt.Printf(`
Usage, examples:

    %s [options] file.xml ...
    %s [options] collection.xml ...
    %s [options] file.data.dz ...
    %s [options] file.dact ...
    find . -name '*.xml' | %s [options]
    find . -name '*.data.dz' | %s [options]
    find . -name '*.dact' | %s [options]

Options:

    -a : output in Alpino, all other options are ignored
    -c : don't include comments
    -d : include debug messages in comments
    -e : skip enhanced dependencies
    -f : don't fix punctuation
    -i regexp : make sent_id from regexp applied to filename
                example: /skip/this/prefix/(.*).xml
    -m : don't fix misplaced heads in coordination
    -M : don't copy metadata to comments
    -p : panic on error
    -s : include short error comments if parse fails
    -t : don't try to restore detokenized sentence
    -v : print version and exit
    -x : include dummy output if parse fails

`, p, p, p, p, p, p, p)
}

func main() {

	flag.Usage = usage
	flag.Parse()

	if *opt_v {
		fmt.Println(alud.VersionID())
		fmt.Println("alpino_ds.dtd version", alpinods.DtdVersion)
		major, minor, patch := dbxml.Version()
		fmt.Printf("DbXML %d.%d.%d\n", major, minor, patch)
		return
	}

	filenames := []string{}
	filenames = append(filenames, flag.Args()...)
	if !util.IsTerminal(os.Stdin) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			filenames = append(filenames, scanner.Text())
		}
		x(scanner.Err())
	}

	if len(filenames) == 0 {
		usage()
		return
	}

	if len(filenames) > 1 {
		multi = true
	}

	var reID *regexp.Regexp
	var options int
	if *opt_c {
		options |= alud.OPT_NO_COMMENTS
	}
	if *opt_d {
		options |= alud.OPT_DEBUG
	}
	if *opt_e {
		options |= alud.OPT_NO_ENHANCED
	}
	if *opt_f {
		options |= alud.OPT_NO_FIX_PUNCT
	}
	if *opt_i != "" {
		var err error
		reID, err = regexp.Compile(*opt_i)
		x(err)
	}
	if *opt_m {
		options |= alud.OPT_NO_FIX_MISPLACED_HEADS
	}
	if *opt_M {
		options |= alud.OPT_NO_METADATA
	}
	if *opt_p {
		options |= alud.OPT_PANIC
	}
	if *opt_t {
		options |= alud.OPT_NO_DETOKENIZE
	}
	if *opt_x {
		options |= alud.OPT_DUMMY_OUTPUT
	}

	for _, filename := range filenames {
		if strings.HasSuffix(filename, ".dact") {
			multi = true
			db, err := dbxml.OpenRead(filename)
			x(err)
			docs, err := db.All()
			x(err)

			for docs.Next() {
				var sentid string
				if *opt_i != "" {
					sentid = reID.FindStringSubmatch(docs.Name())[1]
				}
				doFile([]byte(docs.Content()), docs.Name(), filename, sentid, options)
			}
			x(docs.Error())
			db.Close()
			continue
		}

		if strings.HasSuffix(filename, ".data.dz") {
			multi = true
			corpus, err := compactcorpus.Open(filename)
			x(err)
			r, err := corpus.NewRange()
			x(err)
			for r.HasNext() {
				name, data := r.Next()
				var sentid string
				if *opt_i != "" {
					sentid = reID.FindStringSubmatch(name)[1]
				}
				doFile(data, name, filename, sentid, options)
			}
			continue
		}

		b, err := ioutil.ReadFile(filename)
		x(err)

		var collection Collection

		if xml.Unmarshal(b, &collection) != nil {
			multi = false
			var sentid string
			if *opt_i != "" {
				sentid = reID.FindStringSubmatch(filename)[1]
			}
			doFile(b, filename, "", sentid, options)
			continue
		}

		multi = true
		dir, err := os.Getwd()
		x(err)
		x(os.Chdir(filepath.Dir(filename)))
		for _, f := range collection.Doc {
			b, err := ioutil.ReadFile(f.Href)
			x(err)
			var sentid string
			if *opt_i != "" {
				sentid = reID.FindStringSubmatch(f.Href)[1]
			}
			doFile(b, f.Href, "", sentid, options)
		}
		x(os.Chdir(dir))
	}
}

func doFile(doc []byte, filename, archname, sentid string, options int) {
	if *opt_a {
		result, err := alud.UdAlpino(doc, filename, sentid)
		if multi {
			if archname == "" {
				fmt.Printf("<!-- %s -->\n", filename)
			} else {
				fmt.Printf("<!-- %s : %s -->\n", archname, filename)
			}
		}
		fmt.Println(result)
		if multi {
			fmt.Println()
		}
		if err != nil {
			if archname == "" {
				fmt.Fprintf(os.Stderr, "%s\n  error: %v\n", filename, err)
			} else {
				fmt.Fprintf(os.Stderr, "%s : %s\n  error: %v\n", archname, filename, err)
			}
		}
		return
	}

	if archname != "" {
		fmt.Println("# archive =", archname)
	}
	result, err := alud.Ud(doc, filename, sentid, options)
	if err == nil {
		fmt.Print(result)
	} else {
		s := err.Error()
		if i := strings.Index(s, "\n"); i > 0 {
			s = s[:i]
		}
		m := reSentID.FindSubmatch(doc)

		id1 := ""
		id2 := ""
		if len(m) == 2 {
			id := string(m[1])
			id1 = "# sent_id = " + id + "\n"
			id2 = "  sentence ID: " + id + "\n"
		}

		if *opt_s {
			fmt.Printf("# source = %s\n%s# error = %s\n# auto = %s\n\n", filename, id1, s, alud.VersionID())
		}
		if *opt_x {
			fmt.Println(result)
		}
		if archname == "" {
			fmt.Fprintf(os.Stderr, "%s\n%s  error: %v\n", filename, id2, err)
		} else {
			fmt.Fprintf(os.Stderr, "%s: %s\n%s  error: %v\n", archname, filename, id2, err)
		}
	}
}
