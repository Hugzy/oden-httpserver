package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func buildResponse(r *Request) string {
	var b strings.Builder
	fmt.Println(r.version)
	file, err := os.ReadFile("./main.html")
	check(err)
	b.WriteString("HTTP/1.1 ")
	b.WriteString("200 OK")
	b.WriteString("\r\n")
	b.WriteString("\r\n")
	// b.WriteString("hello world")
	b.WriteString(string(file))

	// response := "HTTP/1.1 200 OK\r\n\r\nhello world"
	fmt.Println(b.String())
	return b.String()
}
