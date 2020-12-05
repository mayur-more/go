package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	
	root, _ := filepath.Abs(".")
	fmt.Println("Processing path", root)
	
	err := filepath.Walk(root, processPath)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func processPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	
	if path != "." {
		if info.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			fmt.Println("File:", path)
		}
	}
	
	return nil
}

/*

 % go run walkDirs.go 
Processing path /Users/mayurmore/eclipse-workspace/golang_practice/Learning_Go/src/ch7
Directory: /Users/mayurmore/eclipse-workspace/golang_practice/Learning_Go/src/ch7
File: /Users/mayurmore/eclipse-workspace/golang_practice/Learning_Go/src/ch7/fromBytes.txt
File: /Users/mayurmore/eclipse-workspace/golang_practice/Learning_Go/src/ch7/fromString.txt
File: /Users/mayurmore/eclipse-workspace/golang_practice/Learning_Go/src/ch7/hello.txt
File: /Users/mayurmore/eclipse-workspace/golang_practice/Learning_Go/src/ch7/readFile.go
File: /Users/mayurmore/eclipse-workspace/golang_practice/Learning_Go/src/ch7/walkDirs.go
File: /Users/mayurmore/eclipse-workspace/golang_practice/Learning_Go/src/ch7/writeFile.go

*/