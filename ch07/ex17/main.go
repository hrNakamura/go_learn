package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack [][]string // stack of element names and values
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			attr := make([]string, 0)
			attr = append(attr, tok.Name.Local)
			for _, val := range tok.Attr {
				if val.Name.Local == "class" || val.Name.Local == "id" {
					attr = append(attr, val.Value)
				}
			}
			stack = append(stack, attr) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				for _, attr := range stack {
					fmt.Printf("%s ", strings.Join(attr, " "))
				}
				fmt.Printf(": %s\n", tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x [][]string, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		for _, n := range x[0] {
			if n == y[0] {
				y = y[1:]
				break
			}
		}
		x = x[1:]
	}
	return false
}
