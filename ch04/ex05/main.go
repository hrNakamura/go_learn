package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 1, 1, 4, 5}
	fmt.Println(a)
	a = remDup(a)
	fmt.Println(a)
	b := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(b)
	b = remDup(b)
	fmt.Println(b)

}

func remDup(s []int) []int {
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			copy(s[i-1:], s[i:])
			return remDup(s[:len(s)-1])
		}
	}
	return s
}
