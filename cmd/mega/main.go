package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "application stopped unexpectedly: %v", err)
		os.Exit(1)
	}
}

func run() error {
	return nil
}
