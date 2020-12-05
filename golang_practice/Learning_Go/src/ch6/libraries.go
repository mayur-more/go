package main

import "fmt"

func main() {

	n1, l1 := FullName("Zaphod", "Beeblebrox")
	fmt.Printf("Fullname: %v, number of chars: %v\n\n", n1, l1)

	n2, l2 := FullNameNakedReturn("Arthur", "Dent")
	fmt.Printf("Fullname: %v, number of chars: %v\n\n", n2, l2)

}

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return 
}
/*

 % go run libraries.go 
Fullname: Zaphod Beeblebrox, number of chars: 17

Fullname: Arthur Dent, number of chars: 11

*/