// Go strings
package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// uint8 is a byte
}


/*
ch2 % go run strings.go 
The colour of magic
19
book[0] = 84 (type uint8)
*/