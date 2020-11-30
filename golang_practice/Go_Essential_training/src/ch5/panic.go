package main

import (
	"fmt"
)

func main() {
		vals := []int{1, 2, 3}
		// This will cause a panic
		v := vals[10]
		fmt.Println(v)
}

/*

% go run panic.go 
panic: runtime error: index out of range [10] with length 3

goroutine 1 [running]:
main.main()
	/Users/mayurmore/eclipse-workspace/golang_practice/Go_Essential_training/src/ch5/panic.go:10 +0x1d
exit status 2

*/