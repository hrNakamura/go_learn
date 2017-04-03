// ch01/ex04 dup2を重複が発見された行を含むファイルの名前を表示するように改造する
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filemap := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filemap)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filemap)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(filemap[line], " "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, filemap map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		files := filemap[input.Text()]
		sort.Strings(files)
		i := sort.SearchStrings(files, f.Name())
		if i >= len(files) {
			filemap[input.Text()] = append(filemap[input.Text()], f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
