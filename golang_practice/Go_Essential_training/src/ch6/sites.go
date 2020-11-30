// Get content type of sites
package main

import (
	"fmt"
	"net/http"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	for _, url := range urls {
		returnType(url)
	}
}

/*
% go run sites.go 
https://golang.org -> text/html; charset=utf-8
https://api.github.com -> application/json; charset=utf-8
https://httpbin.org/xml -> application/xml
*/