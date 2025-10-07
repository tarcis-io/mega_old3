package config

import (
	"time"
)

type (
	Config struct {
		ServerAddress string

		ServerReadTimeout time.Duration

		ServerReadHeaderTimeout time.Duration

		ServerWriteTimeout time.Duration

		ServerIdleTimeout time.Duration

		ServerShutdownTimeout time.Duration
	}
)
