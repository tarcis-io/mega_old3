// Package config loads, parses, and provides the application configuration.
package config

import (
	"fmt"
	"net"
	"os"
	"time"
)

type (
	// Config holds the application configuration.
	Config struct {
		// ServerAddress specifies the TCP address for the server to listen on,
		// in the form "host:port" (e.g. "127.0.0.1:3000").
		// Default: "localhost:8080".
		ServerAddress string

		// ServerReadTimeout is the maximum duration for reading the entire
		// request, including the body.
		// It helps prevent the server from hanging requests indefinitely.
		// Default: 5 * time.Second.
		ServerReadTimeout time.Duration

		// ServerReadHeaderTimeout is the amount of time allowed to read
		// request headers.
		// It helps prevent the server from hanging requests indefinitely and
		// is enforced before the request body is read.
		// Default: 2 * time.Second.
		ServerReadHeaderTimeout time.Duration

		// ServerWriteTimeout is the maximum duration before timing out writes
		// of the response.
		// It helps prevent the server from hanging responses indefinitely.
		// Default: 10 * time.Second.
		ServerWriteTimeout time.Duration

		// ServerIdleTimeout is the maximum amount of time to wait for the next
		// request when keep-alives are enabled.
		// It helps free up server resources from idle connections.
		// Default: 60 * time.Second.
		ServerIdleTimeout time.Duration

		// ServerShutdownTimeout is the maximum duration to wait for active
		// connections to drain before forcing them to close.
		// It ensures a graceful shutdown where ongoing requests can complete.
		// Default: 15 * time.Second.
		ServerShutdownTimeout time.Duration
	}

	parser struct {
		errs []error
	}
)

func New() (*Config, error) {
	config := &Config{}
	return config, nil
}

func newParser() *parser {
	parser := &parser{
		errs: []error{},
	}
	return parser
}

func (parser *parser) env(envKey, envDefault string) string {
	if env, isSet := os.LookupEnv(envKey); isSet {
		return env
	}
	return envDefault
}

func (parser *parser) hostPort(envKey, envDefault string) string {
	env := parser.env(envKey, envDefault)
	host, port, err := net.SplitHostPort(env)
	if err != nil {
		parser.errs = append(parser.errs, fmt.Errorf("failed to parse \"host:port\" (%s) got=%q: %w", envKey, env, err))
		return ""
	}
	return net.JoinHostPort(host, port)
}

func (parser *parser) duration(envKey, envDefault string) time.Duration {
	env := parser.env(envKey, envDefault)
	duration, err := time.ParseDuration(env)
	if err != nil {
		parser.errs = append(parser.errs, fmt.Errorf("failed to parse duration (%s) got=%q: %w", envKey, env, err))
		return 0
	}
	if duration <= 0 {
		parser.errs = append(parser.errs, fmt.Errorf("duration (%s) must be greater than zero got=%q", envKey, env))
		return 0
	}
	return duration
}
