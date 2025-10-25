package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"mega/internal/config"
	"mega/internal/server"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}
	logger, err := newLogger(config)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	slog.SetDefault(logger)
	slog.Info("Running application")
	if err := run(); err != nil {
		slog.Error("Application stopped unexpectedly", "error", err)
		os.Exit(1)
	}
	slog.Info("Application stopped successfully")
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
	slog.SetDefault(logger)
	server, err := server.New(config, logger)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	return server.Run()
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
