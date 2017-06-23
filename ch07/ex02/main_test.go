package main

import (
	"bytes"
	"io"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var b bytes.Buffer
	var w io.Writer
	var cnt *int64
	w, cnt = CountingWriter(&b)
	n, err := w.Write([]byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}

	expect := int64(n)
	if expect != *cnt {
		t.Fatalf("expect=%d, got=%d", expect, *cnt)
	}

	n, err = w.Write([]byte("hello world"))
	expect += int64(n)
	if expect != *cnt {
		t.Fatalf("expect=%d, got=%d", expect, *cnt)
	}
}
