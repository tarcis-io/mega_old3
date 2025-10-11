package config

import (
	"testing"
	"time"
)

type (
	testCase struct {
		name       string
		envValues  map[string]string
		wantConfig *Config
		wantError  bool
	}
)

func TestNew(t *testing.T) {
	testCases := []testCase{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for envKey, envValue := range tc.envValues {
				t.Setenv(envKey, envValue)
			}
		})
	}
}

func mustParseDuration(duration string) time.Duration {
	timeDuration, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
	}
	return timeDuration
}
