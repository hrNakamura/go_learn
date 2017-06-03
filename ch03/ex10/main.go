//練習問題　3.10　再帰呼び出しを行わない
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	//TODO Byte.Bufferを使うこと
	b := []byte(s)
	n := len(b)
	if n <= 3 {
		return s
	}
	m := n % 3
	str := make([]byte, m) //TODO 空スライスとする、長さを0としないと正常な結果とならない
	if m != 0 {
		str = append(str, b[:m]...)
		str = append(str, byte(','))
	}

	for i := m; i < n; i += 3 {
		if i != m {
			str = append(str, byte(','))
		}
		str = append(str, b[i:i+3]...)
	}

	return string(str)
}

//!-
