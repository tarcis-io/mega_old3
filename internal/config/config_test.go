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
