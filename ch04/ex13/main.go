package main

import "net/http"
import "log"

import "fmt"

func main() {
	qURL := "http://www.omdbapi.com/?t=rogue+one&apikey=dff4d11&"
	resp, err := http.Get(qURL)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer resp.Body.Close()
	// fmt.Println(resp.Status)
	fmt.Print(resp.Body)
}
