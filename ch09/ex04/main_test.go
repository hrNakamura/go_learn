package main

import "testing"

func Benchmark1000Pipeline(b *testing.B) {
	in, out := pipeline(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}

func Benchmark10000Pipeline(b *testing.B) {
	in, out := pipeline(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}

func Benchmark100000Pipeline(b *testing.B) {
	in, out := pipeline(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}
