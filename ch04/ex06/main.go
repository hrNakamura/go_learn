package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	a := []byte("abcd　　ef  g")
	fmt.Println(string(a))
	a = remSpace(a)
	fmt.Println(string(a))
}

func remSpace(s []byte) []byte {
	for i := 0; i < utf8.RuneCount(s); i++ {
		if utf8.RuneStart(s[i]) {
			if r, size := utf8.DecodeRune(s[i:]); unicode.IsSpace(r) {
				if r2, _ := utf8.DecodeRune(s[i+size:]); unicode.IsSpace(r2) {
					copy(s[i+size:], s[i+size*2:])
					d := size
					if size > 1 {
						copy(s[i:], s[i+(size-1):])
						s[i] = byte(32) //全角スペースの場合は半角スペースに置換
						d += (size - 1)
					}
					return remSpace(s[:len(s)-d])
				}
			}
		}
	}
	return s
}
