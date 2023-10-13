package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)
}

func main() {
	// defines a pointer variable to find the command line input
	/*msgP := flag.String("msg", "Default message", "The message you want to send")
	flag.Parse()*/

	// this is an easier way to achieve the same thing as above
	stdin := bufio.NewReader(os.Stdin)

	conn, _ := net.Dial("tcp", "127.0.0.1:8030")

	for {
		fmt.Printf("Enter text -> ")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, msg)
		read(conn)
	}
}
