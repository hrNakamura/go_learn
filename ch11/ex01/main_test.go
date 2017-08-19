package main

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input       string
		wantCounts  map[rune]int
		wantUTFLen  [utf8.UTFMax + 1]int
		wantInvalid int
	}{
		{"abcaa", map[rune]int{'a': 3, 'b': 1, 'c': 1}, [utf8.UTFMax + 1]int{0, 5, 0, 0, 0}, 0},
		{"あいういい", map[rune]int{'あ': 1, 'い': 3, 'う': 1}, [utf8.UTFMax + 1]int{0, 0, 0, 5, 0}, 0},
		{"\U0010FFFF\U0010FFFF", map[rune]int{'\U0010FFFF': 2}, [utf8.UTFMax + 1]int{0, 0, 0, 0, 2}, 0},
		{"", map[rune]int{}, [utf8.UTFMax + 1]int{0, 0, 0, 0, 0}, 0},
	}
	for _, test := range tests {
		counts, utflen, invalid, err := charCount(strings.NewReader(test.input))
		if err != nil {
			t.Error(err)
		}
		if !matchMap(counts, test.wantCounts) || !matchUtfLen(utflen, test.wantUTFLen) || invalid != test.wantInvalid {
			t.Errorf("(got, want) = counts: (%v, %v) utflen: (%v, %v) invalid: (%d, %d)", counts, test.wantCounts, utflen, test.wantUTFLen, invalid, test.wantInvalid)
		}
	}
}

func matchMap(got, want map[rune]int) bool {
	if len(got) != len(want) {
		return false
	}
	for key, val := range got {
		if wVal, ok := want[key]; !ok || val != wVal {
			return false
		}
	}
	return true
}

func matchUtfLen(got, want [utf8.UTFMax + 1]int) bool {
	for i, val := range got {
		if val != want[i] {
			return false
		}
	}
	return true
}
