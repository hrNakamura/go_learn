package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	elems := make(map[string]int)
	outline(elems, doc)
	fmt.Println(elems)
}

func outline(elem map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		elem[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(elem, c)
	}
}
