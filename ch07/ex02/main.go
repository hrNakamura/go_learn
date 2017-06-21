package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

type countingWriter struct {
	writer io.Writer
	count  *int64
}

func main() {
	var b bytes.Buffer
	var w io.Writer
	var cnt *int64
	w, cnt = CountingWriter(&b)
	n, err := w.Write([]byte("hello world"))
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(n)
	fmt.Println(*cnt)
}

func (c countingWriter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	if err == nil {
		*c.count += int64(n)
	}
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var count int64
	c := countingWriter{w, &count}
	return c, c.count
}
