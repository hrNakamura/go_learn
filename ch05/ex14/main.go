package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

func printCource(s string) []string {
	fmt.Println(s)
	return prereqs[s]
}

//!+main
func main() {
	if len(os.Args) != 2 {
		return
	}
	course := os.Args[1]
	_, ok := prereqs[course]
	if !ok {
		fmt.Printf("%s not found", course)
		return
	}
	breadthFirst(printCource, []string{os.Args[1]})
}
