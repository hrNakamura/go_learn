// 練習問題4.8 Unicode分類に従った種類ごとのカウント
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int) // counts of Unicode characters
	letterCounts := make(map[rune]int)
	difitCounts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		if unicode.IsLetter(r) {
			letterCounts[r]++
		}
		if unicode.IsDigit(r) {
			difitCounts[r]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\nletter\tcount\n")
	for i, n := range letterCounts {
		if i > 0 {
			fmt.Printf("%q\t%d\n", i, n)
		}
	}
	fmt.Print("\ndigit\tcount\n")
	for i, n := range difitCounts {
		if i > 0 {
			fmt.Printf("%q\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
