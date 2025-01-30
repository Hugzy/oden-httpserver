package main

import (
	"bufio"
	"fmt"
	"runtime/debug"
	"strings"
	"testing"
)

var input = `GET / HTTP/1.1
Host: localhost:8000
Connection: keep-alive
Cache-Control: max-age=0
sec-ch-ua: "Google Chrome";v="129", "Not=A?Brand";v="8", "Chromium";v="129"
sec-ch-ua-mobile: ?0
sec-ch-ua-platform: "Linux"
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
Sec-Fetch-Site: none
Sec-Fetch-Mode: navigate
Sec-Fetch-User: ?1
Sec-Fetch-Dest: document
Sec-Fetch-Dest: document2
Accept-Encoding: gzip, deflate, br, zstd
Accept-Language: en-GB,en;q=0.9,da-DK;q=0.8,da;q=0.7,en-US;q=0.6
Cookie: adminer_sid=9j0fludam9jice8eqhq3n57ffp; adminer_key=5b3c11e67f1de0da2a0893b8adc256d4`

func TestHelloWorld(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(input))
	r, err := parse(reader)
	if err != nil {
		t.Log(string(debug.Stack()))
		t.Error(err)
	}

	if r.url != "/" {
		t.Error("url was not / was", r.url)
	}

	if r.method != GET {
		t.Error("method was not GET was", r.method)
	}

	host := r.fields["Host"].key[0]

	if host != "localhost:8000" {
		t.Error("host was not localhost:8000 was", host)
	}

	for _, v := range r.fields {
		fmt.Println(v)
	}
}
