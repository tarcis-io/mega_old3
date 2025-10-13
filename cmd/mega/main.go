package main

import (
	"fmt"
	"log/slog"
	"os"

	"mega/internal/config"
	"mega/internal/server"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Running application")
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
	logger.Info("Application stopped successfully")
	return nil
}
