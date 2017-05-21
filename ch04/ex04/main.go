package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	a = rotate(a, 3)
	fmt.Println(a)
}

func rotate(s []int, n int) []int {
	s = append(s, s[:n]...)
	return s[n:]
}
