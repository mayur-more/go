package main

import "fmt"

func main() {
	
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	//aNumber := 42
	//isTrue := isTrue

	stringLength, err := fmt.Println(str1, str2, str3)
	
	if err == nil {
		fmt.Println("String Length : ", stringLength)
	}		
}

/*
 ch2 % go run consoleOutput2.go 
The quick red fox jumped over the lazy brown dog.
String Length :  50
*/