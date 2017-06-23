package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}
	url := os.Args[1]
	id := os.Args[2]

	resp, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		os.Exit(1)
	}
	node := ElementByID(doc, id)
	if node == nil {
		fmt.Printf("id=%v is not found\n", id)
	} else {
		fmt.Printf("%v\n", node)
	}
}

//TODO endElementはいらない(nilでよい)
func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(id, doc, startElement, endElement)
}

func forEachNode(id string, n *html.Node, pre, post func(id string, n *html.Node) bool) *html.Node {
	//TODO 結果を変数に格納する必要なし
	if pre != nil {
		if result := pre(id, n); result == true {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(id, c, pre, post)
		if node != nil {
			return node
		}
	}

	if post != nil {
		if result := post(id, n); result == true {
			return n
		}
	}
	return nil
}

//!-forEachNode
var depth int

//!+startend
func startElement(id string, n *html.Node) bool {
	if n.Type == html.ElementNode {
		for _, p := range n.Attr {
			if p.Key == "id" && p.Val == id {
				return true
			}
		}
		depth++
	}
	return false
}

func endElement(id string, n *html.Node) bool {
	if n.Type == html.ElementNode {
		for _, p := range n.Attr {
			if p.Key == "id" && p.Val == id {
				return true
			}
		}
		depth--
	}
	return false
}

//!-startend
