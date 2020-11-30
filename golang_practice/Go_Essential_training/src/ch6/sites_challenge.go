// Get content type of sites (using channels)
package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("%s -> error: %s", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> %s", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	// Create response channel
	ch := make(chan string)
	for _, url := range urls {
		go returnType(url, ch)
	}

	for range urls { // Run number of URLs times
		out := <-ch
		fmt.Println(out)
	}
}

/*
 % go run sites_challenge.go 
https://golang.org -> error: Get "https://golang.org": net/http: TLS handshake timeout
https://httpbin.org/xml -> error: Get "https://httpbin.org/xml": net/http: TLS handshake timeout
https://api.github.com -> application/json; charset=utf-8
*/