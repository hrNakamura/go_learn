package main

import (
	"bufio"
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	retVal := s
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		//TODO Replace関数で置換する(空白もそのままコピーする)
		if strings.Index(word, "$") == 0 {
			retVal = strings.Replace(retVal, word, f(word[1:]), -1)
		}
	}
	return retVal
}

func main() {
	text := "$foo foo　$あああ $_Foo"
	ex := expand(text, strings.ToUpper)
	fmt.Println(ex)
}
