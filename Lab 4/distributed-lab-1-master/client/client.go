package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, error := reader.ReadString('\n')
		if error != nil {
			break
		}
		fmt.Println("")
		fmt.Println(msg)
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	stdin := bufio.NewReader(os.Stdin)
	msg, error := stdin.ReadString('\n')
	if error != nil {
		return
	}
	fmt.Fprintln(conn, msg)
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
	conn, _ := net.Dial("tcp", *addrPtr)
	go read(conn)
	for {
		write(conn)
	}
}
