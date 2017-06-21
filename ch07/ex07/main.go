package main

import (
	"flag"
	"fmt"

	"./tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	//tempconvのString()から゜を取り除いたためヘルプはｃ表記となる
	flag.Parse()
	fmt.Println(*temp)
}
