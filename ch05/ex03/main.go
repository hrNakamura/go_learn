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
	outline(doc)
}

func outline(n *html.Node) {
	if n.Type == html.TextNode && (n.Data == "script" || n.Data == "style") {
		if n.NextSibling != nil {
			outline(n.NextSibling)
		} else {
			return
		}
	}
	if n.Type == html.TextNode {
		fmt.Printf("%v\n", n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(c)
	}
}
