package main

import "fmt"

func returnNonZero() (retval int) {
	//TODO recoverの条件チェックを入れる
	defer func() {
		if p := recover(); p != nil {
			retval = 100
		}
	}()
	panic()
}

func main() {
	fmt.Println(returnNonZero())
}
