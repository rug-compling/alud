package main

import (
	"github.com/pebbe/dbxml"
	"github.com/pebbe/util"
	"github.com/rug-compling/alud"

	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	opt_c = flag.Bool("c", false, "don't include comments")
	opt_d = flag.Bool("d", false, "include debug messages in comments")
	opt_e = flag.Bool("e", false, "skip enhanced dependencies")
	opt_f = flag.Bool("f", false, "don't fix punctuation")
	opt_m = flag.Bool("m", false, "don't fix mixplaced heads in coordination")
	opt_p = flag.Bool("p", false, "panic on error (for development)")
	opt_t = flag.Bool("t", false, "don't try to restore detokenized sentence")

	x = util.CheckErr
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
    %s [options] file.dact ...
    find . -name '*.xml' | %s [options]
    find . -name '*.dact' | %s [options]

Options:

    -c : don't include comments
    -d : include debug messages in comments
    -e : skip enhanced dependencies
    -f : don't fix punctuation
    -m : don't fix misplaced heads in coordication
    -p : panic on error
    -t : don't try to restore detokenized sentence

`, p, p, p, p, p)
}

func main() {

	flag.Usage = usage
	flag.Parse()

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
	if *opt_m {
		options |= alud.OPT_NO_FIX_MISPLACED_HEADS
	}
	if *opt_p {
		options |= alud.OPT_PANIC
	}
	if *opt_t {
		options |= alud.OPT_NO_DETOKENIZE
	}

	for _, filename := range filenames {
		if strings.HasSuffix(filename, ".dact") {
			db, err := dbxml.OpenRead(filename)
			x(err)
			docs, err := db.All()
			x(err)
			for docs.Next() {
				doFile([]byte(docs.Content()), docs.Name(), filename, options)
			}
			x(docs.Error())
			db.Close()
			continue
		}

		b, err := ioutil.ReadFile(filename)
		x(err)

		var collection Collection

		if xml.Unmarshal(b, &collection) != nil {
			doFile(b, filename, "", options)
			continue
		}

		dir, err := os.Getwd()
		x(err)
		x(os.Chdir(filepath.Dir(filename)))
		for _, f := range collection.Doc {
			b, err := ioutil.ReadFile(f.Href)
			x(err)
			doFile(b, f.Href, "", options)
		}
		x(os.Chdir(dir))
	}
}

func doFile(doc []byte, filename, archname string, options int) {
	if archname != "" {
		fmt.Println("# archive =", archname)
	}
	result, err := alud.Ud(doc, filename, options)
	if err != nil {
		s := err.Error()
		if i := strings.Index(s, "\n"); i > 0 {
			s = s[:i]
		}
		fmt.Printf("# source = %s\n# error = %s\n\n", filename, s)
		fmt.Fprintf(os.Stderr, "Error in %s: %s: %v\n", archname, filename, err)
	} else {
		fmt.Print(result)
	}
}
