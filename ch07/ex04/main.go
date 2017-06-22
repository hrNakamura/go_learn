package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

type myNewReader struct {
	s string
}

func (r myNewReader) Read(p []byte) (int, error) {
	n := copy(p, []byte(r.s))
	r.s = r.s[n:]
	if len(r.s) == 0 {
		return n, io.EOF
	}
	return n, nil
}

func MyNewReader(s string) io.Reader {
	return &myNewReader{s}
}

func main() {
	s := "<html><body><p>Hello Go</p></body></html>"
	n, err := html.Parse(MyNewReader(s))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Type=%v\n", n.Type)
	}
}
