package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Elems 要素のスライスを返す
func (s *IntSet) Elems() []int {
	e := []int{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				e = append(e, 64*i+j)
			}
		}
	}
	return e
}

//SymmetricDifferenceWith 非対称集合
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	result := []uint64{}
	for i, tword := range t.words {
		if i >= len(s.words) {
			result = append(result, tword)
		}
		result = append(result, s.words[i]^tword)
	}
	if d := len(s.words) - len(t.words); d > 0 {
		result = append(result, s.words[len(s.words)-d:]...)
	}
	s.words = result
}

//DifferenceWith 差集合
func (s *IntSet) DifferenceWith(t *IntSet) {
	result := []uint64{}
	for i, tword := range t.words {
		if i >= len(s.words) {
			break
		}
		result = append(result, s.words[i]&uint64(^tword))
	}
	if d := len(s.words) - len(t.words); d > 0 {
		result = append(result, s.words[len(s.words)-d:]...)
	}
	s.words = result
}

//IntersectWith 積集合
func (s *IntSet) IntersectWith(t *IntSet) {
	result := []uint64{}
	for i, tword := range t.words {
		if i >= len(s.words) {
			break
		}
		result = append(result, s.words[i]&tword)
	}
	s.words = result
}

// AddAll リストで追加する
func (s *IntSet) AddAll(x ...int) {
	for _, v := range x {
		s.Add(v)
	}
}

// Len 要素数を返す
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for ; word != 0; word &= word - 1 {
			count++
		}
	}
	return count
}

// Remove 要素を削除する
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}
	s.words[word] &= uint64(^(1 << bit))
}

// Clear セットからすべての要素を取り除く
func (s *IntSet) Clear() {
	s.words = []uint64{}
}

// Copy セットのコピーを返す
func (s *IntSet) Copy() *IntSet {
	cp := new(IntSet)
	cp.words = append(cp.words, s.words...)
	return cp
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
