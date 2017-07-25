package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func scan(r io.Reader, input chan<- string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		input <- scanner.Text()
	}
}

//!+
func handleConn(c net.Conn) {
	defer c.Close()
	inputs := make(chan string)
	go scan(c, inputs)
	d := 10 * time.Second
	timer := time.NewTimer(d)
	for {
		select {
		case in := <-inputs:
			timer.Reset(d)
			go echo(c, in, 1*time.Second)
		case <-timer.C:
			return
		}
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
