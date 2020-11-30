// fmt.Sprint example

package main

import (
	"fmt"
)

func main() {
	n := 42
	s := fmt.Sprintf("%d", n)

	fmt.Printf("s = %q (type %T)\n", s, s)
}

/*
ch2 % go run sprintf.go 
s = "42" (type string)
*/