package config

import (
	"testing"
	"time"
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
