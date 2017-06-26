package main

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestMyNewReader(t *testing.T) {
	s := "Hello myNewReader"
	buffer := &bytes.Buffer{}
	n, err := buffer.ReadFrom(MyNewReader(s))
	if err != nil {
		t.Fatal(err)
	}
	if n != int64(len(s)) {
		t.Fatalf("expect length=%d, get Length=%d", len(s), n)
	}
	if buffer.String() != s {
		t.Fatalf("expect %s, get %s", s, buffer.String())
	}

	f, err := ioutil.ReadFile("./index.html")
	if err != nil {
		t.Fatal(err)
	}
	str := string(f)

	node, err := html.Parse(MyNewReader(str))
	if err != nil {
		t.Fatal(err)
	}

	node2, err := html.Parse(strings.NewReader(str))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(node, node2) {
		t.Fail()
	}
}
