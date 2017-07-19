package main

import (
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 1 {
		os.Exit(1)
	}
	var conns []net.Conn
	for _, arg := range os.Args[1:] {
		sp := strings.Split(arg, "=")
		if len(sp) != 2 {
			log.Fatalf("%s no match format", arg)
		}
		conn, err := net.Dial("tcp", sp[1])
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		conns = append(conns, conn)
		os.Stdout.WriteString(sp[0])
		os.Stdout.WriteString("\t\t")
	}
	os.Stdout.WriteString("\n")
	printWall(conns)
}

func printWall(conns []net.Conn) {
	for {
		for i, conn := range conns {
			b := make([]byte, 1024)
			n, err := conn.Read(b)
			if err != nil {
				log.Fatal(err)
			}
			if n > 0 {
				os.Stdout.WriteString("\r")
				for j := 0; j < i; j++ {
					os.Stdout.WriteString("\t\t")
				}
				os.Stdout.Write(b)
			}
		}
		os.Stdout.WriteString("\n")
		time.Sleep(1 * time.Second)
	}
}
