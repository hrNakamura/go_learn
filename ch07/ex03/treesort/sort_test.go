// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort

import (
	"math/rand"
	"sort"
	"testing"

	"gopl.io/ch4/treesort"
)

func TestString(t *testing.T) {

	d := &tree{}
	if d.String() != "0" {
		t.Error(d.String())
	}
	d = add(d, 3)
	if d.String() != "0, 3" {
		t.Error(d.String())
	}
	d = add(d, 1)
	d = add(d, 4)
	if d.String() != "0, 1, 3, 4" {
		t.Error(d.String())
	}
}

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}
