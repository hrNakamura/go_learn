package popcount

import "sync"

// pc[i] is the population count of i.
var pcInitOnce sync.Once
var pc [256]byte
var Inverse bool //1度だけ呼び出されたことを確認するための変数

func initPC() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	Inverse = !Inverse
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	pcInitOnce.Do(initPC)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
