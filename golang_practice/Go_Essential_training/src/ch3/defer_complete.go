package main

import (
	"fmt"
)

func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func worker() {
	defer cleanup("A")
	defer cleanup("B")

	fmt.Println("worker")
}

func main() {
	worker()
}

/*
% go run defer_complete.go 
worker
Cleaning up B
Cleaning up A
*/