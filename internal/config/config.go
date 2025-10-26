// Package config loads, parses, and provides the application configuration.
package config

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"strings"
	"time"
)

// Supported log formats.
const (
	LogFormatJSON = "JSON"
	LogFormatText = "TEXT"
)

// Supported log outputs.
const (
	LogOutputStdout = "STDOUT"
	LogOutputStderr = "STDERR"
)

// Log environment variable key and default values.
const (
	logLevelEnvKey      = "LOG_LEVEL"
	logLevelEnvDefault  = "INFO"
	logFormatEnvKey     = "LOG_FORMAT"
	logFormatEnvDefault = LogFormatJSON
	logOutputEnvKey     = "LOG_OUTPUT"
	logOutputEnvDefault = LogOutputStdout
)

// Server environment variable key and default values.
const (
	serverAddressEnvKey               = "SERVER_ADDRESS"
	serverAddressEnvDefault           = "localhost:8080"
	serverReadTimeoutEnvKey           = "SERVER_READ_TIMEOUT"
	serverReadTimeoutEnvDefault       = "5s"
	serverReadHeaderTimeoutEnvKey     = "SERVER_READ_HEADER_TIMEOUT"
	serverReadHeaderTimeoutEnvDefault = "2s"
	serverWriteTimeoutEnvKey          = "SERVER_WRITE_TIMEOUT"
	serverWriteTimeoutEnvDefault      = "10s"
	serverIdleTimeoutEnvKey           = "SERVER_IDLE_TIMEOUT"
	serverIdleTimeoutEnvDefault       = "60s"
	serverShutdownTimeoutEnvKey       = "SERVER_SHUTDOWN_TIMEOUT"
	serverShutdownTimeoutEnvDefault   = "15s"
)

type (
	// Config holds the application configuration.
	Config struct {
		// LogLevel specifies the minimum level of log messages to output.
		// This will be slog.LevelDebug, slog.LevelInfo, slog.LevelWarn,
		// slog.LevelError, or a numerical level.
		// Default: slog.LevelInfo.
		LogLevel slog.Level

		// LogFormat specifies the output format for log messages.
		// Valid values are "JSON" or "TEXT".
		// Default: "JSON".
		LogFormat string

		// LogOutput specifies the destination for log messages.
		// This will be os.Stdout, os.Stderr, or an *os.File.
		// If the writer is an *os.File, the caller is responsible for closing
		// it.
		// Default: os.Stdout.
		LogOutput io.Writer

		// ServerAddress specifies the TCP address for the server to listen on,
		// in the form "host:port" (e.g. "127.0.0.1:3000").
		// Default: "localhost:8080".
		ServerAddress string

		// ServerReadTimeout is the maximum duration for reading the entire
		// request, including the body.
		// It helps prevent the server from hanging on requests indefinitely.
		// Default: 5 * time.Second.
		ServerReadTimeout time.Duration

		// ServerReadHeaderTimeout is the amount of time allowed to read
		// request headers.
		// It helps prevent the server from hanging on requests indefinitely
		// and is enforced before the request body is read.
		// Default: 2 * time.Second.
		ServerReadHeaderTimeout time.Duration

		// ServerWriteTimeout is the maximum duration before timing out writes
		// of the response.
		// It helps prevent the server from hanging on responses indefinitely.
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
)

// New creates and returns a new Config instance by loading, parsing, and
// validating the application configuration.
// It returns a fully initialized Config on success.
// If any configuration value fails to parse or is invalid, it returns a nil
// Config and a single error that aggregates all the errors found.
func New() (*Config, error) {
	parser := newParser()
	config := &Config{
		LogLevel:                parser.logLevel(logLevelEnvKey, logLevelEnvDefault),
		LogFormat:               parser.logFormat(logFormatEnvKey, logFormatEnvDefault),
		LogOutput:               parser.logOutput(logOutputEnvKey, logOutputEnvDefault),
		ServerAddress:           parser.hostPort(serverAddressEnvKey, serverAddressEnvDefault),
		ServerReadTimeout:       parser.duration(serverReadTimeoutEnvKey, serverReadTimeoutEnvDefault),
		ServerReadHeaderTimeout: parser.duration(serverReadHeaderTimeoutEnvKey, serverReadHeaderTimeoutEnvDefault),
		ServerWriteTimeout:      parser.duration(serverWriteTimeoutEnvKey, serverWriteTimeoutEnvDefault),
		ServerIdleTimeout:       parser.duration(serverIdleTimeoutEnvKey, serverIdleTimeoutEnvDefault),
		ServerShutdownTimeout:   parser.duration(serverShutdownTimeoutEnvKey, serverShutdownTimeoutEnvDefault),
	}
	if err := parser.err(); err != nil {
		return nil, fmt.Errorf("failed to create config: %w", err)
	}
	return config, nil
}

type (
	// parser is a helper struct for parsing configuration values.
	parser struct {
		// errs holds any errors that occurred during parsing.
		errs []error
	}
)

// newParser creates and returns a new initialized parser instance.
func newParser() *parser {
	return &parser{}
}

// err returns a single error containing all errors found during the parsing
// process.
// If no errors were recorded, it returns nil.
func (parser *parser) err() error {
	if len(parser.errs) == 0 {
		return nil
	}
	return errors.Join(parser.errs...)
}

// appendError adds a new error to the parser's internal slice of errors.
// It allows the parser to accumulate all configuration errors and report them
// at once.
func (parser *parser) appendError(err error) {
	parser.errs = append(parser.errs, err)
}

// env retrieves the value of an environment variable named by the key, and
// returns it.
// If the variable is not set, it returns the provided default value.
func (parser *parser) env(envKey, envDefault string) string {
	if env, isSet := os.LookupEnv(envKey); isSet {
		return env
	}
	return envDefault
}

// logLevel retrieves a log level string from an environment variable, parses
// it, and returns it as a slog.Level.
// It accepts "DEBUG", "INFO", "WARN" or "ERROR" (case-insensitive) or a
// numerical level (e.g., -4, 0, 4, 8).
// If parsing or validation fails, it records the error and returns a zero
// log level.
func (parser *parser) logLevel(envKey, envDefault string) slog.Level {
	env := parser.env(envKey, envDefault)
	var logLevel slog.Level
	if err := logLevel.UnmarshalText([]byte(env)); err != nil {
		parser.appendError(fmt.Errorf("failed to parse log level (%s) got=%q: %w", envKey, env, err))
		return 0
	}
	return logLevel
}

// logFormat retrieves a log format string from an environment variable,
// validates it, and returns it.
// It accepts "JSON" or "TEXT" (case-insensitive).
// If validation fails, it records the error and returns an empty string.
func (parser *parser) logFormat(envKey, envDefault string) string {
	switch env := strings.ToUpper(parser.env(envKey, envDefault)); env {
	case LogFormatJSON:
		return LogFormatJSON
	case LogFormatText:
		return LogFormatText
	default:
		parser.appendError(fmt.Errorf("failed to parse log format (%s) got=%q: it must be either %q or %q", envKey, env, LogFormatJSON, LogFormatText))
		return ""
	}
}

// logOutput retrieves a log output string from an environment variable,
// validates it, and returns it as an io.Writer.
// It accepts "STDOUT", "STDERR" (case-insensitive), or a file path.
// If a file path is provided, the returned io.Writer will be an *os.File that
// the caller is responsible for closing.
// If parsing or validation fails, it records the error and returns nil.
func (parser *parser) logOutput(envKey, envDefault string) io.Writer {
	switch env := parser.env(envKey, envDefault); strings.ToUpper(env) {
	case LogOutputStdout:
		return os.Stdout
	case LogOutputStderr:
		return os.Stderr
	case "":
		parser.appendError(fmt.Errorf("failed to parse log output (%s) got=%q: it must be either %q, %q, or a file path", envKey, env, LogOutputStdout, LogOutputStderr))
		return nil
	default:
		file, err := os.OpenFile(env, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			parser.appendError(fmt.Errorf("failed to open log output file (%s) got=%q: %w", envKey, env, err))
			return nil
		}
		return file
	}
}

// hostPort retrieves a "host:port" string from an environment variable,
// validates it, and returns it.
// It also checks if the port is within the IANA range.
// If validation fails, it records the error and returns an empty string.
func (parser *parser) hostPort(envKey, envDefault string) string {
	env := parser.env(envKey, envDefault)
	if _, err := net.ResolveTCPAddr("tcp", env); err != nil {
		parser.appendError(fmt.Errorf("failed to parse \"host:port\" (%s) got=%q: %w", envKey, env, err))
		return ""
	}
	return env
}

// duration retrieves a duration string from an environment variable, parses
// it, and returns it as a time.Duration.
// It also checks if the duration is greater than zero.
// If parsing or validation fails, it records the error and returns a zero
// duration.
func (parser *parser) duration(envKey, envDefault string) time.Duration {
	env := parser.env(envKey, envDefault)
	duration, err := time.ParseDuration(env)
	if err != nil {
		parser.appendError(fmt.Errorf("failed to parse duration (%s) got=%q: %w", envKey, env, err))
		return 0
	}
	if duration <= 0 {
		parser.appendError(fmt.Errorf("duration (%s) must be greater than zero got=%q", envKey, env))
		return 0
	}
	return duration
}
