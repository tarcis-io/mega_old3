package main

import (
	"log"
	"log/slog"

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
	return nil, nil
}
