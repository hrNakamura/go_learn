package main

import (
	"fmt"
	"log"
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
	"linear algebra":        {"calculus"},
}

//!-table

//!+main
func main() {
	topo, err := topoSort(prereqs)
	if err != nil {
		log.Fatal(err)
	}
	for i, course := range topo {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
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

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	if checkLoop(order) {
		return nil, fmt.Errorf("loop %v", prereqs)
	}
	return order, nil
}

func checkLoop(s []string) bool {
	for i := 0; i < len(s); i++ {
		key1 := s[i]
		reqs := prereqs[key1]
		for j := i + 1; len(reqs) != 0 && j < len(s); j++ {
			key2 := s[j]
			for k := 0; k < len(reqs); k++ {
				if reqs[k] == key2 {
					return true
				}
			}
		}
	}

	return false
}
