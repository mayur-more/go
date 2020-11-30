package main

import (
	"fmt"
)

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func main() {
	values := []int{1, 2, 3, 4}
	fmt.Println(values)
	doubleAt(values, 2)
	fmt.Println(values)
}

/*

% go run params.go 
[1 2 3 4]
[1 2 6 4]

*/