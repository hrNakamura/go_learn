package main

import (
	"crypto/sha256"
	"fmt"

	"gopl.io/ch2/popcount"
)

//!+

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	for i := 0; i < len(c1); i++ {
		v := c1[i] ^ c2[i]
		if v != 0 {
			fmt.Printf("%02x(%08b), %02x(%08b), diffCount=%v\n", c1[i], c1[i], c2[i], c2[i], popcount.PopCount(uint64(v)))
		}
	}
}
