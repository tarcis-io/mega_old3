package main

import (
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
		return err
	}
	server, err := server.New(config, logger)
	if err != nil {
		return err
	}
	if err := server.Run(); err != nil {
		return err
	}
	logger.Info("Application stopped successfully")
	return nil
}
