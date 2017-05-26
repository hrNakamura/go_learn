package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 3 {
		fmt.Printf("%s , %s anagram=%v", os.Args[1], os.Args[2], anagram(os.Args[1], os.Args[2]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, r := range s1 {
		c1 := strings.Count(s1, string(r))
		c2 := strings.Count(s2, string(r))
		if c1 != c2 {
			return false
		}
	}
	return true
}

//!-
