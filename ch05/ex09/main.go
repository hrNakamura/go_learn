package main

import (
	"bufio"
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	var retVal string
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if strings.Index(word, "$") == 0 {
			retVal += f(word[1:])
			retVal += " "
		} else {
			retVal += word
			retVal += " "
		}
	}
	return retVal
}

func main() {
	text := "$foo foo"
	ex := expand(text, strings.ToUpper)
	fmt.Println(ex)
}
