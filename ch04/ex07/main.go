//　UTF‐8でエンコードされた文字列を表す[]byteスライスをそのスライス内で逆順にする
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := []byte("あbcえe")
	fmt.Println(string(a))
	fmt.Println(a)
	reverse(a)
	fmt.Println(string(a))
}

func reverse(s []byte) {
	for i := len(s) - 1; i >= 0; {
		_, size := utf8.DecodeRune(s[0:])
		for j := size - 1; j >= 0; j-- {
			move(s, j, i)
			fmt.Println(s)
			i--
		}
	}
}

//stで指定した要素をendの位置まで右に移動する
func move(s []byte, st, end int) {
	for i := st; i < end; i++ {
		s[i+1], s[i] = s[i], s[i+1]
	}
}
