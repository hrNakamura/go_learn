package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

type LimitedReader struct {
	R io.Reader
	N int64
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) >= l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

func main() {
	reader := bytes.NewReader([]byte("abcdefghijklmnopqrstuvwxyz"))
	data, err := ioutil.ReadAll(LimitReader(reader, 10)) // read 10 characters
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	data, err = ioutil.ReadAll(LimitReader(reader, 10)) // read 10 characters
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	data, err = ioutil.ReadAll(LimitReader(reader, 10)) // read 10 characters
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
