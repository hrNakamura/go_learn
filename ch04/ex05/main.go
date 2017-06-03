package main

import (
	"fmt"
)

//TODO stringスライス
func main() {
	a := []string{"aaa", "bbb", "bbb", "bb", "cc", "dd"}
	fmt.Println(a)
	a = remDup(a)
	fmt.Println(a)
}

func remDup(s []string) []string {
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			copy(s[i-1:], s[i:])
			return remDup(s[:len(s)-1])
		}
	}
	return s
}
