// Calculate the mean of two numbers
package main

import (
	"fmt"
)

func main() {
	var x float64
	var y float64

	x = 1
	y = 2

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	var mean float64
	mean = (x + y) / 2
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
//
//---------------- OUTPUT ----------------------
//
//src % go run mean_1.go
//# command-line-arguments
//./mean_1.go:19:9: invalid operation: x + y (mismatched types float64 and int)