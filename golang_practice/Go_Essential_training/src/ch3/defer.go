package main

import (
	"fmt"
)

func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func worker() {
	defer cleanup("A")

	fmt.Println("worker")
}

func main() {
	worker()
}

/*
 % go run defer.go 
worker
Cleaning up A
*/