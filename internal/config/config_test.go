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
}

func mustParseDuration(duration string) time.Duration {
	timeDuration, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
	}
	return timeDuration
}
