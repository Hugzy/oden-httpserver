package main

import (
	"fmt"
	"os"
	"strings"
)

func buildResponse(r *Request) string {
	var b strings.Builder
	fmt.Println(r.version)
	filename := fmt.Sprintf(".%v", r.url)
	file, err := os.ReadFile(filename)
	if err != nil {
		file, err = os.ReadFile("./notfound.html")
		if err != nil {
			panic(err)
		}
	}

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
