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
	nvalue := popcount.PopCount(v)
	fmt.Printf("PopCount(Normal) of %d = %d(%0.3fsec.)\n", v, nvalue, time.Since(start).Seconds())
	start = time.Now()
	lvalue := popcount.PCLoop(v)
	fmt.Printf("PopCount(Loop) of %d = %d(%0.3fsec.)\n", v, lvalue, time.Since(start).Seconds())
}
