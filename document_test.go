package main

/*

dit is alleen maar om een coverage te krijgen

gebruik:

    go test -coverprofile=coverage.out .
    go tool cover -html=coverage.out

*/

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func TestDoc(t *testing.T) {
	b, err := ioutil.ReadFile("collection")
	x(err)

	var collection Collection
	x(xml.Unmarshal(b, &collection))
	for _, f := range collection.Doc {
		b, err := ioutil.ReadFile(f.Href)
		x(err)
		doFile(b, f.Href)
	}
}
