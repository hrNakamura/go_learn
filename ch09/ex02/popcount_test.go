package popcount_test

import (
	"testing"

	popcount "myProject/go_learn/ch09/ex02"
)

func TestPopCount(t *testing.T) {
	for _, test := range []struct {
		bytes    uint64
		expected int
	}{
		{0x00, 0},
		{0xFF, 8},
		{0x1234567890ABCDEF, 32},
	} {
		if got := popcount.PopCount(test.bytes); got != test.expected || popcount.Inverse != true {
			t.Errorf("value = 0x%X, got = %d, expected = %d\n", test.bytes, got, test.expected)
		}
	}
}
