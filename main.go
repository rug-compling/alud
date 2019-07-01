package main

import (
	//	"github.com/kr/pretty"
	"github.com/pebbe/util"

	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	for _, filename := range os.Args[1:] {
		doc, err := ioutil.ReadFile(filename)
		x(err)

		doDoc(doc, filename)
	}

	if !util.IsTerminal(os.Stdin) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			filename := scanner.Text()
			doc, err := ioutil.ReadFile(filename)
			x(err)

			result, err := doDoc(doc, filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error in %s: %v\n", filename, err)
			} else {
				fmt.Print(result)
			}
		}
	}

}
