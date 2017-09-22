package cycle

import (
	"reflect"
	"unsafe"
)

func cycle(x reflect.Value, seen map[pType]bool) bool {
	//TODO これでは単に同じフィールドを指すだけで循環だと判定する
	if x.CanAddr() {
		p := pType{unsafe.Pointer(x.UnsafeAddr()), x.Type()}
		if seen[p] {
			return true
		}
		seen[p] = true
	}
	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return cycle(x.Elem(), seen)
	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if cycle(x.Index(i), seen) {
				return true
			}
		}
		return false
	case reflect.Struct:
		for i := 0; i < x.NumField(); i++ {
			if cycle(x.Field(i), seen) {
				return true
			}
		}
		return false
	case reflect.Map:
		for _, k := range x.MapKeys() {
			if cycle(x.MapIndex(k), seen) || cycle(k, seen) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func Cycle(x interface{}) bool {
	seen := make(map[pType]bool)
	return cycle(reflect.ValueOf(x), seen)
}

type pType struct {
	p unsafe.Pointer
	t reflect.Type
}
