package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch7/eval"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Print("Input Exp:")
	stdin.Scan()
	expStr := stdin.Text()

	fmt.Print("Input Valiables(ex. <var1>=<value> <var2>=<value> ...):")
	stdin.Scan()
	envStr := stdin.Text()

	env := eval.Env{}
	fields := strings.Fields(envStr)
	for _, str := range fields {
		ops := strings.Split(str, "=")
		if len(ops) != 2 {
			fmt.Fprintf(os.Stderr, "valiable input error: %s", str)
			os.Exit(1)
		}
		val, err := strconv.ParseFloat(ops[1], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "valiable value parse fail: %s", err)
			os.Exit(1)
		}
		env[eval.Var(ops[0])] = val
	}

	exp, err := eval.Parse(expStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "expression parse fail; %s", err)
	}
	fmt.Println(exp.Eval(env))
	os.Exit(0)
}
