package config

import (
	"reflect"
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

func mustParseDuration(duration string) time.Duration {
	timeDuration, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
	}
	return timeDuration
}
