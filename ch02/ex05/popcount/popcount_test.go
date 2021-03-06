package popcount

import (
	"testing"
)

const x uint64 = 123

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

func BenchmarkPCLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PCLoop(x)
	}
}

func BenchmarkPCShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PCShift(x)
	}
}

func BenchmarkPCClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PCClear(x)
	}
}
