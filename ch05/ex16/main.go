package main

import (
	"fmt"
	"os"
)

func join(strs ...string) string {
	var str string
	for _, s := range strs {
		str += s
	}
	return str
}

func main() {
	fmt.Println(join(os.Args[1:]...))
}
