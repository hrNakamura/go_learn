package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"gopl.io/ch5/links"
)

var tokens = make(chan struct{}, 20)
var maxDepth int
var seen = make(map[string]bool)
var mutex = sync.Mutex{}

//!+crawl
func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("depth:%d, url:%s\n", depth, url)
	if depth >= maxDepth {
		return
	}
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	for _, link := range list {
		mutex.Lock()
		if !seen[link] {
			seen[link] = true
			mutex.Unlock()
			wg.Add(1)
			go crawl(link, depth+1, wg)
		} else {
			mutex.Unlock()
		}
	}
}

//!-crawl

//!+main
func main() {
	flag.IntVar(&maxDepth, "depth", 3, "max crawl depth")
	flag.Parse()

	wg := sync.WaitGroup{}
	for _, link := range flag.Args() {
		wg.Add(1)
		go crawl(link, 0, &wg)
	}
	wg.Wait()
}
