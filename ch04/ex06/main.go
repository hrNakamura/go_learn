package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	a := []byte("あいう　　ef  　g")
	fmt.Printf("%q\n", string(a))
	a = remSpace(a)
	fmt.Printf("%q\n", string(a))
}

func remSpace(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		if utf8.RuneStart(s[i]) {
			if r, size := utf8.DecodeRune(s[i:]); unicode.IsSpace(r) {
				if r2, size2 := utf8.DecodeRune(s[i+size:]); unicode.IsSpace(r2) {
					copy(s[i+size:], s[i+size+size2:])
					d := size
					if size != 1 {
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
