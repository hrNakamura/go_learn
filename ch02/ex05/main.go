package main

import (
	"fmt"
	"os"
	"strconv"

	"time"

	"./popcount"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	s := os.Args[1]
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		os.Exit(2)
	}
	start := time.Now()
	c := popcount.PopCount(v)
	fmt.Printf("PopCount(Normal) of %d = %d(%0.3fsec.)\n", v, c, time.Since(start).Seconds())
	start = time.Now()
	c = popcount.PCLoop(v)
	fmt.Printf("PopCount(Loop) of %d = %d(%0.3fsec.)\n", v, c, time.Since(start).Seconds())
	start = time.Now()
	c = popcount.PCShift(v)
	fmt.Printf("PopCount(Shift Only) of %d = %d(%0.3fsec.)\n", v, c, time.Since(start).Seconds())
	start = time.Now()
	c = popcount.PCClear(v)
	fmt.Printf("PopCount(Clear) of %d = %d(%0.3fsec.)\n", v, c, time.Since(start).Seconds())
}
