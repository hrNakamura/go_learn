package equal

import (
	"math"
	"reflect"
	"unsafe"
)

const epsilon = 1e-9

func equalNumber(x, y float64) bool {
	if x == y {
		return true
	}
	d := math.Abs(x - y)
	b := math.Max(x, y) * epsilon
	// fmt.Printf("x=%g, y=%g, diff=%.18f, epsilon=%.18f, border=%.18f\n", x, y, d, epsilon, b)
	if d <= epsilon || d <= b {
		return true
	}
	return false
}

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}
	switch x.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return equalNumber(float64(x.Int()), float64(y.Int()))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return equalNumber(float64(x.Uint()), float64(y.Uint()))

	case reflect.Float32, reflect.Float64:
		return equalNumber(float64(x.Float()), float64(y.Float()))

	case reflect.Complex64, reflect.Complex128:
		return equalNumber(float64(real(x.Complex())), float64(real(y.Complex()))) && equalNumber(float64(imag(x.Complex())), float64(imag(y.Complex())))
	}
	panic("unreachable")
}

func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}
