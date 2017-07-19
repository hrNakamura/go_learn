package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn, tz string) {
	defer c.Close()
	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatal(err)
	}
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	port := flag.String("port", "8000", "listen port number")
	timezone := flag.String("tz", "Asia/Tokyo", "set time zone")
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		fmt.Println(conn.LocalAddr())
		go handleConn(conn, *timezone) // handle connections concurrently
	}
	//!-
}
