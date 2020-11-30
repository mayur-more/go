// Makign HTTP calls
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// GET request
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

/*
% go run httpc.go        
{
  "args": {}, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Host": "httpbin.org", 
    "User-Agent": "Go-http-client/2.0", 
    "X-Amzn-Trace-Id": "Root=1-5fc4719c-0fe5af1d5f9656c16a21a71e"
  }, 
  "origin": "27.57.254.37", 
  "url": "https://httpbin.org/get"
}

*/