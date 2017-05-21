package main

import (
	"fmt"
)

const (
	// _ = 1e+3 * iota
	KB = 1000.0
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Printf("KB=%v\n", KB)
	fmt.Printf("MB=%v\n", MB)
	fmt.Printf("GB=%v\n", GB)
	fmt.Printf("TB=%v\n", TB)
	fmt.Printf("PB=%v\n", PB)
	fmt.Printf("EB=%v\n", EB)
	fmt.Printf("ZB=%v\n", ZB)
	fmt.Printf("YB=%v\n", YB)
}
