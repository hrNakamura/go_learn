package main

import (
	"fmt"
	"sort"
)

type Palindrome []rune

func (x Palindrome) Len() int           { return len(x) }
func (x Palindrome) Less(i, j int) bool { return x[i] < x[j] }
func (x Palindrome) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func IsPalindrome(s sort.Interface) bool {
	i := 0
	j := s.Len() - 1
	for i < j {
		if !s.Less(i, j) && !s.Less(j, i) {
			//何もしない
		} else {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	const s = "abcba"
	fmt.Printf("%s is palindrome:%v\n", s, IsPalindrome(Palindrome(s)))
	const ja = "あいういあ"
	fmt.Printf("%s is palindrome:%v\n", ja, IsPalindrome(Palindrome(ja)))
}
