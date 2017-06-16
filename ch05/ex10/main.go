package main

import (
	"fmt"
	"sort"
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
	var keys []int
	topo := topoSort(prereqs)
	for key, _ := range topo {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for key := range keys {
		fmt.Printf("%v:\t%v\n", key, topo[key])
	}

}

func topoSort(m map[string][]string) map[int]string {
	var count int
	order := map[int]string{}
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[count] = item
				count++
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	// sort.Strings(keys)
	visitAll(keys)
	return order
}
