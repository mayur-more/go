package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`

	words := strings.Fields(text)
	counts := map[string]int{} // Empty map
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	fmt.Println(counts)
}

/*
ch2 % go run words.go 
map[a:1 and:2 catch:1 me:2 needles:2 pins:2 sail:1 sew:1 the:1 to:1 wind:1]
*/