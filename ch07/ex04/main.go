package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type myNewReader struct {
	s string
}

func (r myNewReader) Read(p []byte) (int, error) {
	n := copy(p, []byte(r.s))
	return n, nil
}

func MyNewReader(s string) io.Reader {
	return &myNewReader{s}
}

func main() {

	doc, err := html.Parse(MyNewReader(os.Args[1]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
