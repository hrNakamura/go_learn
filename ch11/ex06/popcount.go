//!+
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PCLoop is PopCount using forLoop
func PCLoop(x uint64) int {
	var ret int
	var i uint
	for ; i < 8; i++ {
		ret += int(pc[byte(x>>(i*8))])
	}
	return ret
}

// PCShift is PopCount using bit shift only
func PCShift(x uint64) int {
	var ret int
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			ret++
		}
		x = x >> 1
	}
	return ret
}

//PCClear is PopCount using first bit clear
func PCClear(x uint64) int {
	var ret int
	for ; x != 0; x = x & (x - 1) {
		ret++
	}
	return ret
}

//!-
