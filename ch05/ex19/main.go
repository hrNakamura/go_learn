package main

import "fmt"

func returnNonZero() (retval int) {
	defer func() {
		recover()
		retval = 100
	}()
	panic("0")
}

func main() {
	fmt.Println(returnNonZero())
}
