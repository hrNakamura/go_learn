package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

type myNewReader struct {
	s string
	i int
}

func (r *myNewReader) Read(p []byte) (n int, err error) {
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(p, r.s[r.i:])
	r.i += n
	return
}

func MyNewReader(s string) io.Reader {
	return &myNewReader{s, 0}
}

func main() {

	p := "<html><body><p>hello go</p></body></html>"
	n, err := html.Parse(MyNewReader(p))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Type=%v\n", n.Type)
	}
}
