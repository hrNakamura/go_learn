// ch01/ex03 非効率なechoとJoinを用いたechoの実行時間の差を計測するs
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//!+
func main() {
	start := time.Now()
	fmt.Println(Join(os.Args[1:]))
	fmt.Printf("%.4fsec. elapsed - strings.join()\n", time.Since(start).Seconds())
	start = time.Now()
	fmt.Println(Append(os.Args[1:]))
	fmt.Printf("%.4fsec. elapsed - string +\n", time.Since(start).Seconds())
}

//Join strings.Join()によって文字列配列を連結する
//a 連結する文字列配列
func Join(a []string) string {
	return strings.Join(a, " ")
}

//Append 演算子+によって文字列配列を連結する
//a 連結する文字列配列
func Append(a []string) string {
	var s, sep string
	for i := 0; i < len(a); i++ {
		s += sep + a[i]
		sep = " "
	}
	return s
}
