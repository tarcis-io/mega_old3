package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"mega/internal/config"
	"mega/internal/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "application stopped unexpectedly: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	config, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}
	logger, err := newLogger(config, os.Stdout)
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}
	server, err := server.New(config, logger)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	logger.Info("Running application")
	if err := server.Run(); err != nil {
		logger.Error("Application stopped unexpectedly", "error", err)
		return err
	}
	logger.Info("Application stopped successfully")
	return nil
}

func newLogger(cfg *config.Config, writer io.Writer) (*slog.Logger, error) {
	loggerHandlerOptions := &slog.HandlerOptions{
		Level: cfg.LogLevel,
	}
	var loggerHandler slog.Handler
	switch cfg.LogFormat {
	case config.LogFormatJSON:
		loggerHandler = slog.NewJSONHandler(writer, loggerHandlerOptions)
	case config.LogFormatText:
		loggerHandler = slog.NewTextHandler(writer, loggerHandlerOptions)
	default:
		return nil, fmt.Errorf("unknown log format: %s", cfg.LogFormat)
	}
	logger := slog.New(loggerHandler)
	return logger, nil
}
