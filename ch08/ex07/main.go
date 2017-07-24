package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"gopl.io/ch5/links"
)

var host string

func saveCopyPage(current string) error {
	url, err := url.Parse(current)
	if err != nil {
		return err
	}
	if host == "" {
		host = url.Host
	}
	if host != url.Host {
		return nil
	}
	dir := url.Host
	var filename string
	if filepath.Ext(url.Path) == "" {
		dir = filepath.Join(dir, url.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(dir, filepath.Dir(url.Path))
		filename = filepath.Join(dir, url.Path)
	}
	fmt.Printf("%s save %s\n", current, filename)
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}
	resp, err := http.Get(current)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

//!-breadthFirst
var tokens = make(chan struct{}, 20)

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	err := saveCopyPage(url)
	if err != nil {
		log.Print(err)
	}
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

//!+main
func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
