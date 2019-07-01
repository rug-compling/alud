package main

import (
	"github.com/pebbe/util"

	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	x = util.CheckErr
)

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
