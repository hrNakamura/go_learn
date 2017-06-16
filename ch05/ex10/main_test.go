package main

import (
	"sort"
	"testing"
)

func TestTopSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		r1 := topoSort(prereqs)
		if !EqMap(r1) {
			t.Fatal("top count not topological order")
		}
	}
}

func EqMap(m map[int]string) bool {
	var keys []int
	for key := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for i := 0; i < len(keys); i++ {
		v1 := m[keys[i]]
		reqs := prereqs[v1]
		for j := i + 1; len(reqs) != 0 && j < len(keys); j++ {
			v2 := m[keys[j]]
			for k := 0; k < len(reqs); k++ {
				if reqs[k] == v2 {
					return false
				}
			}
		}
	}

	return true
}
