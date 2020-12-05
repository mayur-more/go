package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Text: ")
    str, _ := reader.ReadString('\n')
    fmt.Println(str)
}
/*
% go run consoleInput2.go
Enter Text: one
one

% go run consoleInput2.go
Enter Text: one two three
one two three

*/