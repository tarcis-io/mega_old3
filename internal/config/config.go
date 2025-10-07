// Package config loads, parses, and provides the application configuration.
package config

import (
	"time"
)

type (
	// Config
	Config struct {
		// ServerAddress
		// Default: "localhost:8080".
		ServerAddress string

		// ServerReadTimeout
		// Default: 5 * time.Second.
		ServerReadTimeout time.Duration

		// ServerReadHeaderTimeout
		// Default: 2 * time.Second.
		ServerReadHeaderTimeout time.Duration

		// ServerWriteTimeout
		// Default: 10 * time.Second.
		ServerWriteTimeout time.Duration

		// ServerIdleTimeout
		// Default: 60 * time.Second.
		ServerIdleTimeout time.Duration

		// ServerShutdownTimeout
		// Default: 15 * time.Second.
		ServerShutdownTimeout time.Duration
	}
)
