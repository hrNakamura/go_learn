package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func main() {
	const text = "Hello go world \n go language\n excercise"
	var w WordCounter
	w.Write([]byte(text))
	fmt.Println(w)

	var l LineCounter
	l.Write([]byte(text))
	fmt.Println(l)
}

func (c *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	var count int
	for sc.Scan() {
		count++
	}
	*c = WordCounter(count)
	return count, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanLines)
	var count int
	for sc.Scan() {
		count++
	}
	*c = LineCounter(count)
	return count, nil
}
