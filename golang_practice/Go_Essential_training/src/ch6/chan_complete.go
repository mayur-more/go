// channels
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// This will block
	/*
		<-ch
		fmt.Println("Here")
	*/

	go func() {
		// Send number of the channel
		ch <- 353
	}()

	// Receive from the channel
	val := <-ch
	fmt.Printf("got %d\n", val)

	fmt.Println("-----")

	// Send multiple
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	fmt.Println("-----")

	// close to signal we're done
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("received %d\n", i)
	}
}

/*
% go run chan_complete.go 
got 353
-----
sending 0
received 0
sending 1
received 1
sending 2
received 2
-----
sending 0
received 0
sending 1
received 1
sending 2
received 2
*/