package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		in   string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"a,b,c", ":", 1},
	}
	for _, test := range tests {
		words := strings.Split(test.in, test.sep)
		if got, want := len(words), test.want; got != want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.in, test.sep, got, want)
		}
	}
}
