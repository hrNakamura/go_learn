package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const remoteRoot = "C:\\Users\\Admin\\Documents\\_src\\go\\src\\myProject\\go_learn"

type ftpConn struct {
	conn         net.Conn
	prevCmd      string
	dir          string
	sep          string
	err          error
	username     string
	password     string
	addr         string
	pasvListener net.Listener
}

func newFTPConn(c net.Conn) *ftpConn {
	separator := "/"
	if runtime.GOOS == "windows" {
		separator = "\\"
	}
	return &ftpConn{conn: c, dir: remoteRoot, sep: separator}
}

func (c *ftpConn) println(s ...interface{}) {
	s = append(s, "\r\n")
	_, c.err = fmt.Fprint(c.conn, s...)
}

func (c *ftpConn) user(cmds []string) {
	if len(cmds) < 2 {
		c.println("501 Syntax error in parameters or arguments.")
		return
	}
	c.username = cmds[1]
	c.println("331 User name okay, need password.")
}

func (c *ftpConn) pass(cmds []string) {
	if len(cmds) < 2 {
		c.println("501 Syntax error in parameters or arguments.")
		return
	}
	c.password = cmds[1]
	c.println("230 logged in, proceed.")
}

func (c *ftpConn) port(cmds []string) {
	if len(cmds) < 2 {
		c.println("501 Syntax error in parameters or arguments.")
		return
	}
	ips := strings.Split(cmds[1], ",")
	if len(ips) != 6 {
		c.println("501 Syntax error in parameters or arguments.")
		return
	}
	p1, err := strconv.Atoi(ips[4])
	if err != nil {
		c.println("501 Syntax error in parameters or arguments.")
		return
	}
	p2, err := strconv.Atoi(ips[5])
	if err != nil {
		c.println("501 Syntax error in parameters or arguments.")
		return
	}
	c.addr = fmt.Sprintf("%s.%s.%s.%s:%d", ips[0], ips[1], ips[2], ips[3], p1*256+p2)
	c.println("200 PORT Command okay.")
}

func (c *ftpConn) pasv() {
	var err error
	c.pasvListener, err = net.Listen("tcp4", "")
	if err != nil {
		c.println("451 aborted. Local error in processing.")
		return
	}
	_, port, err := net.SplitHostPort(c.pasvListener.Addr().String())
	if err != nil {
		c.println("451 aborted. Local error in processing.")
		c.pasvListener.Close()
		return
	}
	fmt.Println(port)
	host, _, err := net.SplitHostPort(c.conn.LocalAddr().String())
	if err != nil {
		c.println("451 aborted. Local error in processing.")
		c.pasvListener.Close()
		return
	}
	ipAdder, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		c.println("451 aborted. Local error in processing.")
		c.pasvListener.Close()
		return
	}
	ips := ipAdder.IP.To4()
	pVal, err := strconv.Atoi(port)
	if err != nil {
		c.println("451 aborted. Local error in processing.")
		c.pasvListener.Close()
		return
	}
	adder := fmt.Sprintf("%d,%d,%d,%d", ips[0], ips[1], ips[2], ips[3]) + fmt.Sprintf(",%d,%d", pVal/256, pVal%256)
	c.println(fmt.Sprintf("227 Entering Passive Mode (%s).", adder))
}

func (c *ftpConn) syst() {
	var osType string
	switch runtime.GOOS {
	case "windows":
		osType = "Windows NT"
	default:
		osType = "UNIX"
	}
	c.println(fmt.Sprintf("215 %s", osType))
}

func (c *ftpConn) pwd() {
	c.println(fmt.Sprintf("257 \"%s\" is current directory", c.dir))
}

func (c *ftpConn) list(cmds []string) {
	var target string
	var err error
	switch len(cmds) {
	case 1:
		target = "."
	case 2:
		target = cmds[1]
	default:
		c.println("501 Syntax error in parameters or arguments.")
		return
	}
	conn, err := c.createDataConn()
	if err != nil {
		fmt.Println(err)
		c.println("425 Can't open data connection.")
		return
	}
	defer conn.Close()
	c.println("150 Opening ASCII mode data connection")

	absPath := filepath.Join(c.dir, target)
	file, err := os.Open(absPath)
	if err != nil {
		fmt.Println(err)
		c.println("550 Requested action not taken.")
		return
	}
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		c.println("450 Requested file action not taken.")
		return
	}
	if stat.IsDir() {
		files, err := file.Readdirnames(0)
		if err != nil {
			fmt.Println(err)
			c.println("550 Requested action not taken.")
			return
		}
		for _, f := range files {
			_, err = fmt.Fprintf(conn, "%s\r\n", f)
			if err != nil {
				c.println("426 Connection closed; transfer aborted.")
				return
			}
		}
	} else {
		_, err = fmt.Fprintf(conn, "%s\r\n", target)
		if err != nil {
			c.println("426 Connection closed; transfer aborted.")
			return
		}
	}
	c.println("226 Closing data connection.")
}

func (c *ftpConn) createDataConn() (conn io.ReadWriteCloser, err error) {
	switch c.prevCmd {
	case "PORT":
		conn, err = net.Dial("tcp", c.addr)
	case "PASV":
		conn, err = c.pasvListener.Accept()
	default:
		return nil, fmt.Errorf("previuos command not Connection: %s", c.prevCmd)
	}
	return
}

func (c *ftpConn) cd(dst string) {
	// var sep string
	// if runtime.GOOS == "windows" {
	// 	sep = "\\"
	// } else {
	// 	sep = "/"
	// }
	// dirs := strings.Split(dst, sep)
	// var dstFull string
	// for i, path := range dirs {
	// }
}

func (c *ftpConn) handleConn() {
	defer c.conn.Close()
	c.println("220 Ready.")
	s := bufio.NewScanner(c.conn)
LOOP:
	for s.Scan() {
		fmt.Println(s.Text())
		cmds := strings.Fields(s.Text())
		if len(cmds) == 0 {
			continue
		}
		switch strings.ToUpper(cmds[0]) {
		case "USER":
			c.user(cmds)
		case "PASS":
			c.pass(cmds)
		case "PORT":
			c.port(cmds)
		case "SYST":
			c.syst()
		case "PWD":
			c.pwd()
		case "LIST":
			c.list(cmds)
		case "PASV":
			c.pasv()
		case "QUIT":
			c.println("221 Goodbye.")
			break LOOP
		default:
			c.println("502 Command not implemented.")
		}
		c.prevCmd = strings.ToUpper(cmds[0])
	}
	fmt.Printf("Close FTP Server:%s", c.conn.LocalAddr().String())
}

func main() {
	port := flag.Int("p", 8000, "FTP port")
	flag.Parse()
	listener, err := net.Listen("tcp4", fmt.Sprintf(":%d", *port))
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
