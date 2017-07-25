package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type RootSize struct {
	root string
	size int64
}

//!+
func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	rootSize := make(chan RootSize)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, rootSize)
	}
	go func() {
		n.Wait()
		close(rootSize)
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles = make(map[string]int64)
	var nbytes = make(map[string]int64)
loop:
	for {
		select {
		case rs, ok := <-rootSize:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles[rs.root]++
			nbytes[rs.root] += rs.size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes) // final totals
	//!+
	// ...select loop...
}

//!-

func printDiskUsage(nfiles, nbytes map[string]int64) {
	var tFiles int64
	var tSizes float64
	for key := range nfiles {
		fmt.Printf("%s: %d files  %.1f GB\n", key, nfiles[key], float64(nbytes[key])/1e9)
		tFiles += nfiles[key]
		tSizes += float64(nbytes[key])
	}
	fmt.Printf("total: %d files  %.1f GB\n", tFiles, tSizes/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(dir, root string, n *sync.WaitGroup, rootSize chan<- RootSize) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, root, n, rootSize)
		} else {
			rootSize <- RootSize{root, entry.Size()}
		}
	}
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
