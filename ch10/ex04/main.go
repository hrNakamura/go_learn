package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func contain(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func getPackages(name string) ([]string, error) {
	args := []string{"list", "-f={{.ImportPath}}", name}
	lst, err := exec.Command("go", args...).Output()
	if err != nil {
		return nil, err
	}
	return strings.Fields(string(lst)), nil
}

func getDepends(packages []string) (map[string][]string, error) {
	depList := make(map[string][]string)
	for _, pkg := range packages {
		depList[pkg] = make([]string, 0)
	}

	args := []string{"list", `-f={{.ImportPath}} {{join .Deps " "}}`, "..."}
	lst, err := exec.Command("go", args...).Output()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(bytes.NewReader(lst))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		pkg := fields[0]
		depends := fields[1:]
		for _, dep := range depends {
			if _, ok := depList[dep]; ok && !contain(depList[dep], pkg) {
				depList[dep] = append(depList[dep], pkg)
			}
		}
	}
	return depList, nil
}

func main() {
	pkgs, err := getPackages(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	deps, err := getDepends(pkgs)
	if err != nil {
		log.Fatal(err)
	}
	for key, val := range deps {
		fmt.Printf("%s package depends:\n", key)
		for _, v := range val {
			fmt.Println(v)
		}
		fmt.Println()
	}
}
