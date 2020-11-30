// pkg/errors example
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

// Config holds configuration
type Config struct {
	// configuration fields go here (redacted)
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't open configuration file")
	}

	defer file.Close()

	cfg := &Config{}
	// Parse file here (redacted)
	return cfg, nil

}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main() {
	setupLogging()
	cfg, err := readConfig("/path/to/config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		log.Printf("error: %+v", err)
		os.Exit(1)
	}

	// Normal operation (redacted)
	fmt.Println(cfg)
}

/*

% go run errors_complete.go 
errors_complete.go:9:2: cannot find package "github.com/pkg/errors" in any of:
	/usr/local/Cellar/go/1.15.5/libexec/src/github.com/pkg/errors (from $GOROOT)
	/Users/mayurmore/go/src/github.com/pkg/errors (from $GOPATH)

% go get github.com/pkg/errors

% go run errors_complete.go   
error: can't open configuration file: open /path/to/config.toml: no such file or directory
exit status 1

*/