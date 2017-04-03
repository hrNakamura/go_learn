// ch01/ex02 コマンドライン引数のインデックスとその値を1行ごとに出力する
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "
	for i := 1; i < len(os.Args); i++ {
		s = sep + os.Args[i]
		fmt.Printf("%d%s\n", i, s)
	}
}
