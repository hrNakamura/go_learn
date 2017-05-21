package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	var sha int
	flag.IntVar(&sha, "sha", 256, "SHA mode")
	flag.Parse()

	sc := bufio.NewScanner(os.Stdin)
	var str string
	if sc.Scan() {
		str = sc.Text()
	}
	switch sha {
	case 384:
		fmt.Printf("SHA=%v\n", sha)
		s := sha512.Sum384([]byte(str))
		fmt.Printf("input=%s\n%x\n", str, s)
	case 512:
		fmt.Printf("SHA=%v\n", sha)
		s := sha512.Sum512([]byte(str))
		fmt.Printf("input=%s\n%x\n", str, s)
	default:
		fmt.Printf("SHA=%v\n", 256)
		s := sha256.Sum256([]byte(str))
		fmt.Printf("input=%s\n%x\n", str, s)
	}
}
