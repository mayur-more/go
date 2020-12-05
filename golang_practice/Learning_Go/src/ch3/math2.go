package main

import (
	"fmt"
	"math/big"
	"math"
)

func main() {
	
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum: ", intSum)
	
	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum: ", floatSum)
	
	var b1,b2, b3, bigsum big.Float
	b1.SetFloat64(23.5)
	b2.SetFloat64(65.1)
	b3.SetFloat64(76.3)
	
	bigsum.Add(&b1, &b2).Add(&bigsum, &b3)
	fmt.Printf("Bigsum sum: %.10g\n", &bigsum)
	
	circulRadius := 15.5
	circumference := circulRadius * math.Pi
	fmt.Printf("circumference: %.2f\n", circumference)
}

/*

 % go run math2.go
Integer sum:  125
Float sum:  164.89999999999998
Bigsum sum: 164.9
circumference: 48.69

*/