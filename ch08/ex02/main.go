package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"runtime"
	"strconv"
	"strings"
)

type ftpConn struct {
	conn net.Conn
	dir  string
}

func newFTPConn(c net.Conn) *ftpConn {
	var root string
	if runtime.GOOS == "windows" {
		root = "C:\\"
	} else {
		root = "/"
	}
	return &ftpConn{c, root}
}

func (c *ftpConn) println(s ...interface{}) {
	fmt.Fprintln(c.conn, s...)
}

func (c *ftpConn) pwd() {
	c.println(c.dir)
}

func (c *ftpConn) ls() error {
	files, err := ioutil.ReadDir(c.dir)
	if err != nil {
		return err
	}
	var fInfo []string
	for _, f := range files {
		fInfo = append(fInfo, fmt.Sprintf("%s\t%d\t%s", f.Mode(), f.Size(), f.Name()))
	}
	for _, str := range fInfo {
		c.println(str)
	}
	return nil
}

func (c *ftpConn) cd(dst string) {
	var sep string
	if runtime.GOOS == "windows" {
		sep = "\\"
	} else {
		sep = "/"
	}
	dirs := strings.Split(dst, sep)
	var dstFull string
	for i, path := range dirs {
	}
}

func (c *ftpConn) handleConn() {
	defer c.conn.Close()
	c.println(fmt.Sprintf("Connect FTP Server:%s", c.conn.LocalAddr().String()))
	s := bufio.NewScanner(c.conn)
LOOP:
	for s.Scan() {
		fields := strings.Fields(s.Text())
		if len(fields) == 0 {
			continue
		}
		switch fields[0] {
		case "ls":
			if err := c.ls(); err != nil {
				log.Print(err)
				break LOOP
			}
		case "bye":
			break LOOP
		default:
			c.println(fmt.Sprintf("undefined command:%s\n", fields[0]))
		}
	}
	c.println(fmt.Sprintf("Close FTP Server:%s", c.conn.LocalAddr().String()))
}

func main() {
	port := flag.Int("p", 8000, "FTP port")
	flag.Parse()
	adder := "localhost:" + strconv.Itoa(*port)
	listener, err := net.Listen("tcp", adder)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go newFTPConn(conn).handleConn()
	}
}
