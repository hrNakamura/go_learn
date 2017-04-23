package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	var c tempconv.Celsius = 24.5
	fmt.Printf("Celcius %s = %s\n", c.String(), tempconv.CtoK(c).String())
	c = tempconv.BoilingC
	fmt.Printf("Boiling %s %s %s\n", c.String(), tempconv.CToF(c).String(), tempconv.CtoK(c).String())
}
