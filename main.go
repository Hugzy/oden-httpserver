package main

import (
	"bufio"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":8000")

	defer ln.Close()

	conn, _ := ln.Accept()

	defer conn.Close()

	reader := bufio.NewReader(conn)

	parse(reader)

	response := "HTTP/1.1 200 OK\r\n\r\nhello world"
	conn.Write([]byte(response))
}
