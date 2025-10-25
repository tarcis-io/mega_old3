package main

import (
	"fmt"
	"os"

	"mega/internal/config"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "application stopped unexpectedly: %v", err)
		os.Exit(1)
	}
}

func run() error {
	_, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}
	return nil
}
