package main

import (
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
	_, err := setupConfig()
	if err != nil {
		return err
	}
	return nil
}

func setupConfig() (*config.Config, error) {
	return nil, nil
}
