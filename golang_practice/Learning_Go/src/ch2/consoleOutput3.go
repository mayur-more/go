package main

import "fmt"

func main() {
	
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	fmt.Println(str1, str2, str3)	
	
	stringLength, _ := fmt.Println(str1, str2, str3)
	
//	if err == nil {
		fmt.Println("String Length: ", stringLength)
//	}		
	
	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	fmt.Printf("Value of aNumber as a float: %.2f\n", float64(aNumber))
	
	fmt.Printf("Data Types: %T, %T, %T, %T and %T\n", str1, str2, str3, aNumber, isTrue)
	
	mtString := fmt.Sprintf("Data Types: %T, %T, %T, %T and %T\n", str1, str2, str3, aNumber, isTrue)
	fmt.Print(mtString)
}

/*
  % go run consoleOutput3.go
The quick red fox jumped over the lazy brown dog.
The quick red fox jumped over the lazy brown dog.
String Length:  50
Value of aNumber: 42
Value of isTrue: true
Value of aNumber as a float: 42.00
Data Types: string, string, string, int and bool
Data Types: string, string, string, int and bool
*/