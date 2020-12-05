package main

import "fmt"

func main() {
	doSomething()
	
	addition := addValues(21, 24)
	fmt.Println("Addition: ", addition)
}

func doSomething() {
	fmt.Println("Doing something...")
}

func addValues(v1 int, v2 int) int {
	return v1 + v2
}

/*

 % go run functions.go
Doing something...
Addition:  45

*/