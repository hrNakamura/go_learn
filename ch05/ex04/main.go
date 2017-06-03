package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, img := range n.Attr {
			if img.Key == "src" {
				links = append(links, img.Val)
			}
		}
	}
	if n.Type == html.ElementNode && n.Data == "script" {
		for _, script := range n.Attr {
			if script.Key == "src" {
				links = append(links, script.Val)
			}
		}
	}
	if n.Type == html.ElementNode && n.Data == "link" {
		isCSS := false
		var vals []string
		for _, link := range n.Attr {
			if link.Key == "type" && link.Val == "text/css" {
				isCSS = true
			}
			if link.Key == "href" {
				vals = append(vals, link.Val)
			}
		}
		if isCSS {
			links = append(links, vals...)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
