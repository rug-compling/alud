package main

import (
	"github.com/pebbe/util"

	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	x = util.CheckErr
)

type Collection struct {
	XMLName xml.Name `xml:"collection"`
	Doc     []Doc    `xml:"doc"`
}

type Doc struct {
	Href string `xml:"href,attr"`
}

func main() {

	filenames := []string{}
	filenames = append(filenames, os.Args[1:]...)
	if !util.IsTerminal(os.Stdin) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			filenames = append(filenames, scanner.Text())
		}
		x(scanner.Err())
	}

	for _, filename := range filenames {
		b, err := ioutil.ReadFile(filename)
		x(err)

		var collection Collection

		if xml.Unmarshal(b, &collection) != nil {
			doFile(b, filename)
			continue
		}

		dir, err := os.Getwd()
		x(err)
		x(os.Chdir(filepath.Dir(filename)))
		for _, f := range collection.Doc {
			b, err := ioutil.ReadFile(f.Href)
			x(err)
			doFile(b, f.Href)
		}
		x(os.Chdir(dir))
	}
}

func doFile(doc []byte, filename string) {
	result, err := doDoc(doc, filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in %s: %v\n", filename, err)
	} else {
		fmt.Print(result)
	}
}
