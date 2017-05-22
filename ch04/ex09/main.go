// 練習問題4.9 単語出現頻度のカウント
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // 単語のカウント
	var tCount int                 //全単語数

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts[scanner.Text()]++
		tCount++
	}
	fmt.Printf("count\tfreq\t\tword\n")
	for c, n := range counts {
		fmt.Printf("%d\t%f\t%q\n", n, float64(n)/float64(tCount), c)
	}
}

//!-
