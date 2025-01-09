package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":8000")

	defer ln.Close()

	conn, _ := ln.Accept()

	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		// get line, output
		line, _ := reader.ReadString('\n')
		fmt.Print(string(line))
		if line == "\r\n" || line == "\n" {
			break
		}
	}

	response := "HTTP/1.1 200 OK\r\n\r\nhello world"
	conn.Write([]byte(response))
}
