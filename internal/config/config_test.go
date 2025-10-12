package config

import (
	"reflect"
	"testing"
	"time"
)

type (
	// testCase holds the data for a single test case.
	testCase struct {
		name       string
		envValues  map[string]string
		wantConfig *Config
		wantError  bool
	}
)

// TestNew verifies if the New function correctly creates a new Config
// instance.
// It covers default configuration values, custom valid configuration values
// and error handling for invalid configuration.
func TestNew(t *testing.T) {
	testCases := []testCase{
		{
			name: "should create a new Config instance with default values",
			wantConfig: &Config{
				ServerAddress:           serverAddressEnvDefault,
				ServerReadTimeout:       mustParseDuration(serverReadTimeoutEnvDefault),
				ServerReadHeaderTimeout: mustParseDuration(serverReadHeaderTimeoutEnvDefault),
				ServerWriteTimeout:      mustParseDuration(serverWriteTimeoutEnvDefault),
				ServerIdleTimeout:       mustParseDuration(serverIdleTimeoutEnvDefault),
				ServerShutdownTimeout:   mustParseDuration(serverShutdownTimeoutEnvDefault),
			},
			wantError: false,
		},
		{
			name: "should create a new Config instance with custom values",
			envValues: map[string]string{
				serverAddressEnvKey:           "127.0.0.1:8081",
				serverReadTimeoutEnvKey:       "20s",
				serverReadHeaderTimeoutEnvKey: "10s",
				serverWriteTimeoutEnvKey:      "20s",
				serverIdleTimeoutEnvKey:       "90s",
				serverShutdownTimeoutEnvKey:   "30s",
			},
			wantConfig: &Config{
				ServerAddress:           "127.0.0.1:8081",
				ServerReadTimeout:       mustParseDuration("20s"),
				ServerReadHeaderTimeout: mustParseDuration("10s"),
				ServerWriteTimeout:      mustParseDuration("20s"),
				ServerIdleTimeout:       mustParseDuration("90s"),
				ServerShutdownTimeout:   mustParseDuration("30s"),
			},
			wantError: false,
		},
		{
			name: "should return an error if the server address cannot be parsed: empty",
			envValues: map[string]string{
				serverAddressEnvKey: "",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server address cannot be parsed: localhost",
			envValues: map[string]string{
				serverAddressEnvKey: "localhost",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server address cannot be parsed: localhost:99999",
			envValues: map[string]string{
				serverAddressEnvKey: "localhost:99999",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server read timeout cannot be parsed: empty",
			envValues: map[string]string{
				serverReadTimeoutEnvKey: "",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server read timeout cannot be parsed: 5sec",
			envValues: map[string]string{
				serverReadTimeoutEnvKey: "5sec",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server read timeout cannot be parsed: 0s",
			envValues: map[string]string{
				serverReadTimeoutEnvKey: "0s",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server read timeout cannot be parsed: -5s",
			envValues: map[string]string{
				serverReadTimeoutEnvKey: "-5s",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server read header timeout cannot be parsed: empty",
			envValues: map[string]string{
				serverReadHeaderTimeoutEnvKey: "",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server read header timeout cannot be parsed: 5sec",
			envValues: map[string]string{
				serverReadHeaderTimeoutEnvKey: "5sec",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server read header timeout cannot be parsed: 0s",
			envValues: map[string]string{
				serverReadHeaderTimeoutEnvKey: "0s",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server read header timeout cannot be parsed: -5s",
			envValues: map[string]string{
				serverReadHeaderTimeoutEnvKey: "-5s",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server write timeout cannot be parsed: empty",
			envValues: map[string]string{
				serverWriteTimeoutEnvKey: "",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server write timeout cannot be parsed: 5sec",
			envValues: map[string]string{
				serverWriteTimeoutEnvKey: "5sec",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server write timeout cannot be parsed: 0s",
			envValues: map[string]string{
				serverWriteTimeoutEnvKey: "0s",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server write timeout cannot be parsed: -5s",
			envValues: map[string]string{
				serverWriteTimeoutEnvKey: "-5s",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server idle timeout cannot be parsed: empty",
			envValues: map[string]string{
				serverIdleTimeoutEnvKey: "",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server idle timeout cannot be parsed: 5sec",
			envValues: map[string]string{
				serverIdleTimeoutEnvKey: "5sec",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server idle timeout cannot be parsed: 0s",
			envValues: map[string]string{
				serverIdleTimeoutEnvKey: "0s",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server idle timeout cannot be parsed: -5s",
			envValues: map[string]string{
				serverIdleTimeoutEnvKey: "-5s",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server shutdown timeout cannot be parsed: empty",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server shutdown timeout cannot be parsed: 5sec",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "5sec",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server shutdown timeout cannot be parsed: 0s",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "0s",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server shutdown timeout cannot be parsed: -5s",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "-5s",
			},
			wantError: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for envKey, envValue := range tc.envValues {
				t.Setenv(envKey, envValue)
			}
			config, err := New()
			if (err != nil) != tc.wantError {
				t.Fatalf("New() error got=%v wantError=%t", err, tc.wantError)
			}
			if tc.wantError {
				return
			}
			if !reflect.DeepEqual(config, tc.wantConfig) {
				t.Fatalf("New() *Config\ngot=%#v\nwant=%#v", config, tc.wantConfig)
			}
		})
	}
}

// mustParseDuration is a helper function that parses a duration string into a
// time.Duration.
// It panics if the parsing fails.
func mustParseDuration(duration string) time.Duration {
	timeDuration, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
	}
	return timeDuration
}
