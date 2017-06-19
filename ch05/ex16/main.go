package main

import (
	"fmt"
	"os"
)

func join(sep string, strs ...string) string {
	var str string
	for i, s := range strs {
		str += s
		if i != len(strs)-1 {
			str += sep
		}
	}
	return str
}

func main() {
	fmt.Println(join(os.Args[1], os.Args[2:]...))
}
