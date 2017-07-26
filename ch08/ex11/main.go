package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var cancel = make(chan struct{})
var response = make(chan *http.Response)

func doRequest(url string, wg *sync.WaitGroup) error {
	defer wg.Done()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	response <- resp
	return nil
}

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}
	wg := sync.WaitGroup{}
	start := time.Now()
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go func(url string) {
			if err := doRequest(url, &wg); err != nil {
				log.Fatal(err)
			}
		}(url)
	}

	resp := <-response
	defer resp.Body.Close()
	close(cancel)
	fmt.Printf("elapsed:%f sec.\t%s\n", time.Since(start).Seconds(), resp.Request.URL)
}
