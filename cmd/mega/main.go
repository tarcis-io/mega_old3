package main

import (
	"fmt"
	"log/slog"
	"os"

	"mega/internal/config"
)

func main() {
	slog.Info("Running application")
	if err := run(); err != nil {
		slog.Error("Application stopped unexpectedly", "error", err)
		os.Exit(1)
	}
	slog.Info("Application stopped successfully")
}

func run() error {
	config, err := setupConfig()
	if err != nil {
		return err
	}
	_, err = setupLogger(config)
	if err != nil {
		return err
	}
	return nil
}

func setupConfig() (*config.Config, error) {
	config, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create config: %w", err)
	}
	return config, nil
}

func setupLogger(cfg *config.Config) (*slog.Logger, error) {
	return nil, nil
}
