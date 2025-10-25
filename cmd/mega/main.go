package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"mega/internal/config"
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
