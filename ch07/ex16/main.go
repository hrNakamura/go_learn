package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gopl.io/ch7/eval"
)

func parseEnv(s string) (eval.Env, error) {
	env := eval.Env{}
	fields := strings.Fields(s)
	for _, str := range fields {
		ops := strings.Split(str, "=")
		if len(ops) != 2 {
			return nil, fmt.Errorf("valiable input error: %s", str)
		}
		val, err := strconv.ParseFloat(ops[1], 64)
		if err != nil {
			return nil, fmt.Errorf("valiable value parse fail: %s", err)
		}
		env[eval.Var(ops[0])] = val
	}
	return env, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		expStr := r.FormValue("exp")
		if expStr == "" {
			http.Error(w, "no expression", http.StatusBadRequest)
			return
		}
		exp, err := eval.Parse(expStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		envStr := r.FormValue("env")
		env, err := parseEnv(envStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "%s\n", envStr)
		fmt.Fprintf(w, "%s=%v\n", expStr, exp.Eval(env))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
	// ?exp=x-2&env=x=3
	// ?exp=pow(x, y)&env=x=3%20y=3
}
