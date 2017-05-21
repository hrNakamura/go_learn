//
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", preComma(os.Args[i]))
	}
}

func preComma(s string) string {
	strs := strings.Split(s, ".")
	if len(strs) > 2 {
		panic("invalid strings")
	}
	if len(strs) == 1 {
		return comma(strs[0])
	}
	return comma(strs[0]) + "." + strs[1]
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-
