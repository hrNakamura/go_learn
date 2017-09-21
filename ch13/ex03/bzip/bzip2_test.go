// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bzip_test

import (
	"bytes"
	"compress/bzip2" // reader
	"io"
	"sync"
	"testing"

	"github.com/hrNakamura/go_learn/ch13/ex03/bzip"
)

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := bzip.NewWriter(&compressed)

	// Write a repetitive message in a million pieces,
	// compressing one copy but not the other.
	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Check the size of the compressed stream.
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// Decompress and compare with original.
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decompression yielded a different message")
	}
}

func TestSyncCompress(t *testing.T) {
	var compressed1, compressed2, uncompressed bytes.Buffer
	w1 := bzip.NewWriter(&compressed1)
	w2 := bzip.NewWriter(&compressed2)

	tee := io.MultiWriter(w1, &uncompressed)

	wg := &sync.WaitGroup{}
	writeFunc := func(w io.Writer) {
		defer wg.Done()
		for i := 0; i < 1000000; i++ {
			io.WriteString(w, "hello")
		}
	}
	wg.Add(2)
	go writeFunc(tee)
	go writeFunc(w2)
	wg.Wait()

	if err := w1.Close(); err != nil {
		t.Fatal(err)
	}
	if err := w2.Close(); err != nil {
		t.Fatal(err)
	}

	// Check the size of the compressed stream.
	if got, want := compressed1.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// Decompress and compare with original.
	var decompressed1 bytes.Buffer
	io.Copy(&decompressed1, bzip2.NewReader(&compressed1))
	if !bytes.Equal(uncompressed.Bytes(), decompressed1.Bytes()) {
		t.Errorf("decompression yielded a different message, %v,%v", len(uncompressed.Bytes()),len(decompressed1.Bytes()))
	}

	// Check the size of the compressed stream.
	if got, want := compressed2.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// Decompress and compare with original.
	var decompressed2 bytes.Buffer
	io.Copy(&decompressed2, bzip2.NewReader(&compressed2))
	if !bytes.Equal(uncompressed.Bytes(), decompressed2.Bytes()) {
		t.Error("decompression yielded a different message")
	}
}
