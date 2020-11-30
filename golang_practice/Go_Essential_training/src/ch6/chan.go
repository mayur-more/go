// channels
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// This will block
		<-ch
		fmt.Println("Here")
}

/*
% go run chan.go          
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/Users/mayurmore/eclipse-workspace/golang_practice/Go_Essential_training/src/ch6/chan.go:12 +0x51
exit status 2
*/