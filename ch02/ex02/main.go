package main

import (
	"fmt"
	"os"
	"strconv"

	"strings"

	"bufio"

	"./lengthconv"
	"./weightconv"
	"gopl.io/ch2/tempconv"
)

func main() {
	var args []string
	if len(os.Args) == 1 {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		args = strings.Split(s.Text(), " ")
	} else {
		args = os.Args[1:]
	}
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		p := weightconv.Pond(t)
		g := weightconv.Kilogram(t)
		ft := lengthconv.Feet(t)
		m := lengthconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", p, weightconv.PToKg(p), g, weightconv.KgToP(g))
		fmt.Printf("%s = %s, %s = %s\n\n", ft, lengthconv.FtToM(ft), m, lengthconv.MToFt(m))
	}
}

//!-
