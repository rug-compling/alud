package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/rug-compling/alud/v2"
	"github.com/rug-compling/alud/v2/internal/util"

	"github.com/pebbe/compactcorpus"
	"github.com/pebbe/dbxml"
	"github.com/rug-compling/alpinods"
)

type resultItem struct {
	output string
	errors string
}

type workItem struct {
	doc      []byte
	filename string
	archname string
	sentid   string
	options  int
	chResult chan resultItem
}

type Collection struct {
	XMLName xml.Name `xml:"collection"`
	Doc     []Doc    `xml:"doc"`
}

type Doc struct {
	Href string `xml:"href,attr"`
}

var (
	chWork    chan workItem
	chResults chan chan resultItem
	chDone    chan bool

	opt_a = flag.Bool("a", false, "output in alpino")
	opt_c = flag.Bool("c", false, "don't include comments")
	opt_d = flag.Bool("d", false, "include debug messages in comments")
	opt_e = flag.Bool("e", false, "skip enhanced dependencies")
	opt_f = flag.Bool("f", false, "don't fix punctuation")
	opt_i = flag.String("i", "", "regexp for filtering filename as sentence id")
	opt_m = flag.Bool("m", false, "don't fix mixplaced heads in coordination")
	opt_M = flag.Bool("M", false, "don't copy metadata to comments")
	opt_p = flag.Bool("p", false, "panic on error (for development)")
	opt_r = flag.Int("r", 1, "number of processes")
	opt_s = flag.Bool("s", false, "include short error comments if parse fails")
	opt_t = flag.Bool("t", false, "don't try to restore detokenized sentence")
	opt_v = flag.Bool("v", false, "print version and exit")
	opt_x = flag.Bool("x", false, "include dummy output if parse fails")

	multi = false

	x = util.CheckErr

	reSentID = regexp.MustCompile(`<sentence.*?sentid="(.*?)".*</sentence>`)
)

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

    -a        : output in Alpino, all other options are ignored
    -c        : don't include comments
    -d        : include debug messages in comments
    -e        : skip enhanced dependencies
    -f        : don't fix punctuation
    -i regexp : make sent_id from regexp applied to filename
                example: /skip/this/prefix/(.*).xml
    -m        : don't fix misplaced heads in coordination
    -M        : don't copy metadata to comments
    -p        : panic on error
    -r int    : number of parallel processes (default: 1, max: %d)
    -s        : include short error comments if parse fails
    -t        : don't try to restore detokenized sentence
    -v        : print version and exit
    -x        : include dummy output if parse fails

`, p, p, p, p, p, p, p, max(1, runtime.NumCPU()-2))
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *opt_v {
		fmt.Println("release", alud.Release())
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

	*opt_r = min(*opt_r, runtime.NumCPU()-2)
	if *opt_r > 1 {
		chWork = make(chan workItem, 100)
		chResults = make(chan chan resultItem, 100)
		chDone = make(chan bool)

		// start workers
		for i := 0; i < *opt_r; i++ {
			go func() {
				for work := range chWork {
					output, errors := doFile(work.doc, work.filename, work.archname, work.sentid, work.options)
					work.chResult <- resultItem{output: output, errors: errors}
					close(work.chResult)
				}
			}()
		}

		// start collector
		go func() {
			// print results in correct order
			for chResult := range chResults {
				result := <-chResult
				fmt.Print(result.output)
				fmt.Fprint(os.Stderr, result.errors)
			}
			// signal all is done
			close(chDone)
		}()
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
				startFile([]byte(docs.Content()), docs.Name(), filename, sentid, options)
			}
			x(docs.Error())
			db.Close()
			continue
		}

		if strings.HasSuffix(filename, ".data.dz") || strings.HasSuffix(filename, ".index") {
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
				startFile(data, name, filename, sentid, options)
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
			startFile(b, filename, "", sentid, options)
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
			startFile(b, f.Href, filename, sentid, options)
		}
		x(os.Chdir(dir))
	}

	// wait for go routines to finish
	if *opt_r > 1 {
		close(chResults) // signal collector all jobs are sent
		<-chDone         // wait for collector to finish
	}
}

func startFile(doc []byte, filename, archname, sentid string, options int) {
	if *opt_r < 2 {
		output, errors := doFile(doc, filename, archname, sentid, options)
		fmt.Print(output)
		fmt.Fprint(os.Stderr, errors)
		return
	}

	chResult := make(chan resultItem, 1)

	chWork <- workItem{
		doc:      doc,
		filename: filename,
		archname: archname,
		sentid:   sentid,
		options:  options,
		chResult: chResult,
	}

	chResults <- chResult
}

func doFile(doc []byte, filename, archname, sentid string, options int) (string, string) {
	var output, errors bytes.Buffer

	if *opt_a {
		result, err := alud.UdAlpino(doc, filename, sentid)
		if multi {
			if archname == "" {
				fmt.Fprintf(&output, "<!-- %s -->\n", filename)
			} else {
				fmt.Fprintf(&output, "<!-- %s : %s -->\n", archname, filename)
			}
		}
		fmt.Fprintln(&output, result)
		if multi {
			fmt.Fprintln(&output)
		}
		if err != nil {
			if archname == "" {
				fmt.Fprintf(&errors, "%s\n  error: %v\n", filename, err)
			} else {
				fmt.Fprintf(&errors, "%s : %s\n  error: %v\n", archname, filename, err)
			}
		}
		return output.String(), errors.String()
	}

	if archname != "" && !*opt_c {
		fmt.Fprintln(&output, "# archive =", archname)
	}
	opts := options
	if *opt_s {
		opts |= alud.OPT_DUMMY_OUTPUT
	}
	result, err := alud.Ud(doc, filename, sentid, opts)
	if err == nil {
		fmt.Fprint(&output, result)
	} else {
		s := err.Error()
		if i := strings.Index(s, "\n"); i > 0 {
			s = s[:i]
		}
		m := reSentID.FindSubmatch(doc)

		id2 := ""
		if len(m) == 2 {
			id := string(m[1])
			id2 = "  sentence ID: " + id + "\n"
		}

		if *opt_s {
			for _, line := range strings.SplitAfter(result, "\n") {
				if len(line) > 0 && (line[0] == '#' || line[0] == '\n') {
					fmt.Fprint(&output, line)
				}
			}
		}
		if *opt_x {
			fmt.Fprint(&output, result)
		}
		if archname == "" {
			fmt.Fprintf(&errors, "%s\n%s  error: %v\n", filename, id2, err)
		} else {
			fmt.Fprintf(&errors, "%s: %s\n%s  error: %v\n", archname, filename, id2, err)
		}
	}
	return output.String(), errors.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
