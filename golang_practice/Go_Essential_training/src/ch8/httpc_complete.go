// Makign HTTP calls
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Job is a job description
type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	// GET request
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	fmt.Println("----")

	// POST request
	job := &Job{
		User:   "Saitama",
		Action: "punch",
		Count:  1,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("error: can't encode job - %s", err)
	}

	resp, err = http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
/*
% go run httpc_complete.go 
{
  "args": {}, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Host": "httpbin.org", 
    "User-Agent": "Go-http-client/2.0", 
    "X-Amzn-Trace-Id": "Root=1-5fc471d7-51316ff707b36e502ac10b80"
  }, 
  "origin": "27.57.254.37", 
  "url": "https://httpbin.org/get"
}
----
{
  "args": {}, 
  "data": "{\"user\":\"Saitama\",\"action\":\"punch\",\"count\":1}\n", 
  "files": {}, 
  "form": {}, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Content-Length": "46", 
    "Content-Type": "application/json", 
    "Host": "httpbin.org", 
    "User-Agent": "Go-http-client/2.0", 
    "X-Amzn-Trace-Id": "Root=1-5fc471d8-2634d1d06daf5b4e488ac56d"
  }, 
  "json": {
    "action": "punch", 
    "count": 1, 
    "user": "Saitama"
  }, 
  "origin": "27.57.254.37", 
  "url": "https://httpbin.org/post"
}
*/