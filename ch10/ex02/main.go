package main

import (
	"os"

	"github.com/hrNakamura/go_learn/ch10/ex02/archiver"
	_ "github.com/hrNakamura/go_learn/ch10/ex02/archiver/tar"
	_ "github.com/hrNakamura/go_learn/ch10/ex02/archiver/zip"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	archiver.Open(os.Args[1], os.Stdout)
}
