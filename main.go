package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		fmt.Println("goobye")
		os.Exit(0)
	}()

	fmt.Println("Starting webserver")
	fmt.Println("Webserver listening on port :8000")

	for {
		ln, _ := net.Listen("tcp", ":8000")

		conn, _ := ln.Accept()

		reader := bufio.NewReader(conn)

		request, err := parse(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("request", request)

		response := buildResponse(request)

		conn.Write([]byte(response))

		ln.Close()
		conn.Close()
	}
}
