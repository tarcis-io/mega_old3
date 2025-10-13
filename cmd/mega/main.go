// Package main is the entry point for the mega application.
package main

import (
	"fmt"
	"log/slog"
	"os"

	"mega/internal/config"
	"mega/internal/server"
)

// main is the entry point for the mega application.
// It delegates to the run function and handles the final exit code.
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Running application")
	if err := run(logger); err != nil {
		logger.Error("Application stopped unexpectedly", "error", err)
		os.Exit(1)
	}
	logger.Info("Application stopped successfully")
}

// run loads the configuration, creates a new server and runs it.
// It is responsible for the entire application lifecycle and returns an error
// if any step fails.
func run(logger *slog.Logger) error {
	config, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	server, err := server.New(config, logger)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	if err := server.Run(); err != nil {
		return fmt.Errorf("server stopped unexpectedly: %w", err)
	}
	return nil
}
