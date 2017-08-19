package popcount

import (
	"testing"
)

const x uint64 = 123

func benchPC(b *testing.B, size int, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			f(uint64(j))
		}
	}
}

func BenchmarkPCTable10(b *testing.B) {
	benchPC(b, 10, PopCount)
}

func BenchmarkPCTable100(b *testing.B) {
	benchPC(b, 100, PopCount)
}

func BenchmarkPCTable1000(b *testing.B) {
	benchPC(b, 1000, PopCount)
}

func BenchmarkPCLoop10(b *testing.B) {
	benchPC(b, 10, PCLoop)
}

func BenchmarkPCLoop100(b *testing.B) {
	benchPC(b, 100, PCLoop)
}

func BenchmarkPCLoop1000(b *testing.B) {
	benchPC(b, 1000, PCLoop)
}

func BenchmarkPCShift10(b *testing.B) {
	benchPC(b, 10, PCShift)
}

func BenchmarkPCShift100(b *testing.B) {
	benchPC(b, 100, PCShift)
}

func BenchmarkPCShift1000(b *testing.B) {
	benchPC(b, 1000, PCShift)
}

func BenchmarkPCClear10(b *testing.B) {
	benchPC(b, 10, PCClear)
}

func BenchmarkPCClear100(b *testing.B) {
	benchPC(b, 100, PCClear)
}

func BenchmarkPCClear1000(b *testing.B) {
	benchPC(b, 1000, PCClear)
}
