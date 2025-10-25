package main

import (
	"fmt"
	"log/slog"
	"os"

	"mega/internal/config"
	"mega/internal/server"
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
	logger, err := setupLogger(config)
	if err != nil {
		return err
	}
	server, err := setupServer(config, logger)
	if err != nil {
		return err
	}
	if err := server.Run(); err != nil {
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
		return nil, fmt.Errorf("failed to create logger: unknown log format: %s", cfg.LogFormat)
	}
	logger := slog.New(loggerHandler)
	return logger, nil
}

func setupServer(cfg *config.Config, logger *slog.Logger) (*server.Server, error) {
	server, err := server.New(cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}
	return server, nil
}
