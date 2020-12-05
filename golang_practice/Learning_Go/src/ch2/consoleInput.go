package main

import (
    "fmt"
)

func main() {
    var s string
    fmt.Scanln(&s)
    fmt.Println(s)
}
/*
% go run consoleInput.go  
one
one

% go run consoleInput.go
one two three
one
*/