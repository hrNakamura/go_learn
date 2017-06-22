package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	nodes := make([]*html.Node, 0)
	tags := make(map[string]bool, len(name))
	for _, tag := range name {
		tags[tag] = true
	}
	pre := func(n *html.Node) {
		if n.Type != html.ElementNode {
			return
		}
		_, ok := tags[n.Data]
		if ok {
			nodes = append(nodes, n)
		}
		return
	}
	forEachNode(doc, pre, nil)

	return nodes
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

func main() {
	if len(os.Args) < 3 {
		log.Fatal("input url tags")
	}
	url := os.Args[1]
	tags := os.Args[2:]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	nodes := ElementByTagName(doc, tags...)

	fmt.Println(len(nodes))
}
