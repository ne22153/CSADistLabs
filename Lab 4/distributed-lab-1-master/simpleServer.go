package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
		fmt.Fprintln(conn, "OK")
	}
}

func main() {
	ln, _ := net.Listen("tcp", ":8030")
	// the for loop means that it will remain active, even after recieving a message
	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}
