package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Text: ")
    str, _ := reader.ReadString('\n')
    fmt.Println(str)
    
    fmt.Print("Enter a number: ")
    str, _ = reader.ReadString('\n')
    f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
    if err != nil {
    	fmt.Println(err)
    }else {
    	fmt.Println("Value of number:", f)
    }    
}
/*
 % go run consoleInput4.go
Enter Text: Some one
Some one

Enter a number: 43
Value of number: 43

% go run consoleInput4.go
Enter Text: Mayur
Mayur

Enter a number: 5.6
Value of number: 5.6
*/