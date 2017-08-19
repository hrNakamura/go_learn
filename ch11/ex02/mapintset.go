package intset

import (
	"bytes"
	"fmt"
	"sort"
)

type MapIntSet struct {
	mapset map[int]bool
}

func NewMapIntSet() *MapIntSet {
	return &MapIntSet{map[int]bool{}}
}

func (s *MapIntSet) Has(x int) bool {
	return s.mapset[x]
}

func (s *MapIntSet) Add(x int) {
	s.mapset[x] = true
}

func (s *MapIntSet) UnionWith(t IntSet) {
	for _, v := range t.Ints() {
		s.Add(v)
	}
}

func (s *MapIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, v := range s.Ints() {
		if i != 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *MapIntSet) Ints() []int {
	ints := make([]int, 0)
	for key := range s.mapset {
		ints = append(ints, key)
	}
	sort.Ints(ints)
	return ints
}
