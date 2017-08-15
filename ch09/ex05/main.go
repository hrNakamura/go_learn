package main

import (
	"fmt"
	"time"
)

func Pingpong() {
	messageA := make(chan string)
	messageB := make(chan string)
	end := make(chan struct{})
	var count int64
	ticker := time.NewTicker(1 * time.Second)
	var second int64
	go func() {
		messageB <- "Hello"
	LOOP:
		for {
			select {
			case <-ticker.C:
				second++
				fmt.Printf("count=%v, seconds=%v\n", count, second)
				if second >= 10 {
					break LOOP
				}
				count = 0
			case m := <-messageA:
				count++
				messageB <- m
			}
		}
		end <- struct{}{}
	}()

	go func() {
		for {
			m := <-messageB
			messageA <- m
		}
	}()
	<-end
	fmt.Println()
}

func main() {
	Pingpong()
}
