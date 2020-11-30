// Calculate the mean of two numbers
package main

import (
	"fmt"
)

func main() {
	var x int
	var y int

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	x = 1
	y = 2

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	var mean int
	mean = (x + y) / 2
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
//
//---------------- OUTPUT ----------------------
//
//src % go run mean.go
//x=0, type of int
//y=0, type of int
//x=1, type of int
//y=2, type of int
//result: 1, type of int