package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

const SPACE = " "

type Method int

type Request struct {
	fields  map[string]*Field
	url     string
	version string
	method  Method
}

type Field struct {
	name          string
	key           []string
	isSingleValue bool
}

const (
	GET Method = iota
	HEAD
	POST
	PUT
	DELETE
	CONNECT
	OPTIONS
	TRACE
	NIL
)

var methodMap = map[Method]string{
	GET:     "GET",
	HEAD:    "HEAD",
	POST:    "POST",
	PUT:     "PUT",
	DELETE:  "DELETE",
	CONNECT: "CONNECT",
	OPTIONS: "OPTIONS",
	TRACE:   "TRACE",
}

func getMethod(s string) (Method, error) {
	switch s {
	case "GET":
		return GET, nil
	case "HEAD":
		return HEAD, nil
	case "POST":
		return POST, nil
	case "PUT":
		return PUT, nil
	case "DELETE":
		return DELETE, nil
	case "CONNECT":
		return CONNECT, nil
	case "OPTIONS":
		return OPTIONS, nil
	case "TRACE":
		return TRACE, nil
	}
	err := fmt.Sprint("http method was not one of the allowed ones:", s)
	return NIL, errors.New(err)
}

func (m Method) String() string {
	return methodMap[m]
}

func (f *Field) String() string {
	var builder strings.Builder
	builder.WriteString(f.name)
	keys := fmt.Sprintf("%s", f.key)
	builder.WriteString(keys)
	return builder.String()
}

func (r *Request) parseRequestLine(s string) error {
	s = strings.Trim(s, "\n\r")
	split := strings.Split(s, SPACE)
	method, err := getMethod(split[0])
	if err != nil {
		return errors.New("cannot parse request")
	}
	r.method = method
	r.url = split[1]
	r.version = split[2]

	return nil
}

func (f *Field) parseField(s string) {
	s = strings.Trim(s, "\n\r")
	nameKey := strings.SplitN(s, ":", 2)
	name := nameKey[0]
	secondPart := strings.TrimSpace(nameKey[1])
	keys := strings.Split(secondPart, ",")
	f.name = name
	for _, key := range keys {
		key = strings.Trim(key, " ")
		f.key = append(f.key, key)
	}
}

func parse(reader *bufio.Reader) (*Request, error) {
	r := &Request{}
	r.fields = make(map[string]*Field)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("could not read request line")
	}
	r.parseRequestLine(requestLine)

	for {
		line, err := reader.ReadString('\n')
		if line == "\r\n" || line == "\n" {
			break
		}
		if err != nil {
			// return nil, errors.New("could not read header")
			break
		}

		f := &Field{}
		f.parseField(line)
		field := r.fields[f.name]

		if field == nil {
			r.fields[f.name] = f
		} else {
			r.fields[f.name].key = append(r.fields[f.name].key, f.key...)
		}
	}

	return r, nil
}
