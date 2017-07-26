package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
type client struct {
	Output chan<- string // an outgoing message channel
	Name   string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.Output <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.Output <- "Member:"
			for c := range clients {
				cli.Output <- c.Name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Output)
		}
	}
}

//!-broadcaster

const timeout = 300 * time.Second

//!+handleConn
func handleConn(conn net.Conn) {
	out := make(chan string) // outgoing client messages
	in := make(chan string)
	err := make(chan error)
	go clientWriter(conn, out)
	go clientReader(conn, in, err)

	out <- "your name?:"
	timer := time.NewTimer(30 * time.Second)
	var who string
	select {
	case name := <-in:
		who = name
	case <-timer.C:
		return
	}
	timer.Reset(timeout)
	cli := client{out, who}
	out <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli
	go func() {
		<-timer.C
		conn.Close()
	}()
LOOP:
	for {
		select {
		case msg := <-in:
			timer.Reset(timeout)
			messages <- who + ": " + msg
		case <-timer.C:
			conn.Close()
			break LOOP
		case <-err:
			break LOOP
		}
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func clientReader(conn net.Conn, ch chan<- string, err chan<- error) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
	err <- input.Err()
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
