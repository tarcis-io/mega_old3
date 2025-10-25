package main

import (
	"fmt"
	"log/slog"
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
	config, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}
	_, err = newLogger(config)
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}
	return nil
}

func newLogger(cfg *config.Config) (*slog.Logger, error) {
	return nil, nil
}
