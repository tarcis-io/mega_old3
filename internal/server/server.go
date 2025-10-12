package server

import (
	"net/http"
	"time"

	"mega/internal/config"
)

type (
	Server struct {
		address string

		router *http.ServeMux

		readTimeout time.Duration

		readHeaderTimeout time.Duration

		writeTimeout time.Duration

		idleTimeout time.Duration

		shutdownTimeout time.Duration
	}
)

func New(config *config.Config) (*Server, error) {
	server := &Server{}
	return server, nil
}
