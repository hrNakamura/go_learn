package main

import (
	"fmt"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
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

//!-table

//!+main
func main() {
	for i, v := range topoSort(prereqs) {
		fmt.Printf("%v:\t%v\n", i, v)
	}

}

//TODO 戻り値は[]stringでよい
func topoSort(m map[string][]string) []string {
	order := []string{}
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	// var keys []string
	for key := range m {
		// keys = append(keys, key)
		visitAll([]string{key})
	}

	// sort.Strings(keys)
	// visitAll(keys)
	return order
}
