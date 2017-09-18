package equal

import "testing"

func TestEqual(t *testing.T) {
	tests := []struct {
		x, y interface{}
		want bool
	}{
		{int(1), int(1), true},
		{float64(1.23), float64(1.23), true},
		{1 + 2i, 1 + 2i, true},
		{int(1), int(2), false},
		{int64(1e+9), int64(1e+9 + 1), true},
		{int64(1e+9), int64(1e+9 - 1), true},
		{float64(1.0), float64(1.0 + 1e-10), true},
		{float64(1.0), float64(1.0 - 1e-9), true},
		{float64(1.0), float64(1.0 + 1e-9), false},
		{complex(1, 1e+9), complex(1-1e-9, 1e+9+1), true},
	}
	for _, test := range tests {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%v, %v) = %t", test.x, test.y, !test.want)
		} else {
			t.Logf("Equal(%v, %v) = %t", test.x, test.y, test.want)
		}
	}
}
