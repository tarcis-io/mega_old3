package main

import (
	"log/slog"
	"os"
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
	return nil
}
