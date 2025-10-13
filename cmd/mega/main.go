package main

import (
	"log/slog"
	"os"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Running application")
	return nil
}
