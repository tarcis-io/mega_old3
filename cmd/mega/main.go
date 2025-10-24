// Package main is the entry point for the mega application.
package main

import (
	"log/slog"
	"os"

	"mega/internal/config"
)

func main() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}
	if err := run(config, nil); err != nil {
		os.Exit(1)
	}
}

func run(config *config.Config, logger *slog.Logger) error {
	return nil
}
