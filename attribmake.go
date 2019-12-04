// +build attrib

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	attribs := strings.Fields(string(data))

	fmt.Print(`//
// THIS IS A GENERATED FILE. DO NOT EDIT.
//

package alud

type fullNode struct {
`)

	skip := map[string]bool{
		"begin":  true,
		"end":    true,
		"id":     true,
		"index":  true,
		"lemma":  true,
		"postag": true,
		"pt":     true,
		"rel":    true,
		"word":   true,
		"cat":    true,
	}

	for _, attrib := range attribs {
		if !skip[attrib] {
			up := upper(attrib)
			fmt.Printf("\t%-12s string `xml:\"%s,attr,omitempty\"`\n", up, attrib)
		}
	}
	fmt.Print("}\n")
}

func upper(s string) string {
	return strings.Replace(strings.Title(strings.Replace(strings.Replace(s, "_", " ", -1), "-", " ", -1)), " ", "", -1)
}
