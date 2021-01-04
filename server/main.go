package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, _ := net.Listen("tcp", ":8080")

	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("New message recieved:", message)
		newMessage := strings.ToUpper(message)
		conn.Write([]byte(newMessage + "\n"))
	}
}
