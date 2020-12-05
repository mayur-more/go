package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Text: ")
    str, _ := reader.ReadString('\n')
    fmt.Println(str)
    
    fmt.Print("Enter a number: ")
    str, _ = reader.ReadString('\n')
    f, err := strconv.ParseFloat(str, 64)
    if err != nil {
    	fmt.Println(err)
    }else {
    	fmt.Println("Value of number:", f)
    }    
}
/*
 % go run consoleInput3.go
Enter Text: one two
one two

Enter a number: 4.2
strconv.ParseFloat: parsing "4.2\n": invalid syntax

*/