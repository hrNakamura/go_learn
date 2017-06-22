package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

func printStartElement(n *html.Node) {
	fmt.Printf("%*s<%s", depth*2, "", n.Data)
	for _, a := range n.Attr {
		fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
	}
	if n.FirstChild == nil {
		fmt.Println("/>")
	} else {
		fmt.Println(">")
	}
}

func printEndElement(n *html.Node) {
	if n.FirstChild != nil {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func printStartComment(n *html.Node) {
	fmt.Printf("<!--%s", n.Data)
}

func printEndComment(n *html.Node) {
	fmt.Println("-->")
}

func printTextNode(n *html.Node) {
	str := strings.TrimSpace(n.Data)
	if str != "" {
		fmt.Printf("%*s%s", depth*2, "", str)
	}
}

//!+startend
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		printStartElement(n)
		depth++
	}
	if n.Type == html.TextNode {
		printTextNode(n)
	}
	if n.Type == html.CommentNode {
		printStartComment(n)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		printEndElement(n)
	}
	if n.Type == html.CommentNode {
		depth--
		printEndComment(n)
	}
}

//!-startend
