package intset

import (
	"bytes"
	"fmt"
)

const bitSize = 32 << (^uint(0) >> 63)

type WordIntSet struct {
	words []uint
}

func (s *WordIntSet) Has(x int) bool {
	word, bit := x/bitSize, uint(x%bitSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *WordIntSet) Add(x int) {
	word, bit := x/bitSize, uint(x%bitSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *WordIntSet) UnionWith(t IntSet) {
	for _, v := range t.Ints() {
		s.Add(v)
	}
}

func (s *WordIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *WordIntSet) Ints() []int {
	ints := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				ints = append(ints, bitSize*i+j)
			}
		}
	}
	return ints
}

//!-string
