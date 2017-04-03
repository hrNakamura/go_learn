package main

import "testing"

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arg := []string{"a", "b", "c"}
		Join(arg)
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arg := []string{"a", "b", "c"}
		Append(arg)
	}
}
