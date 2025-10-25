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
		fmt.Fprintf(os.Stderr, "application stopped unexpectedly: %v", err)
		os.Exit(1)
	}
}

func run() error {
	config, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}
	logger, err := newLogger(config)
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}
	_, err = server.New(config, logger)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	return nil
}

func newLogger(cfg *config.Config) (*slog.Logger, error) {
	loggerHandlerOptions := &slog.HandlerOptions{
		Level: cfg.LogLevel,
	}
	var loggerHandler slog.Handler
	switch cfg.LogFormat {
	case config.LogFormatJSON:
		loggerHandler = slog.NewJSONHandler(os.Stdout, loggerHandlerOptions)
	case config.LogFormatText:
		loggerHandler = slog.NewTextHandler(os.Stdout, loggerHandlerOptions)
	default:
		return nil, fmt.Errorf("unknown log format: %s", cfg.LogFormat)
	}
	logger := slog.New(loggerHandler)
	return logger, nil
}
