package main

import (
	"myProject/go_learn/ch10/ex02/archiver"
	_ "myProject/go_learn/ch10/ex02/archiver/tar"
	_ "myProject/go_learn/ch10/ex02/archiver/zip"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	archiver.Open(os.Args[1], os.Stdout)
}
