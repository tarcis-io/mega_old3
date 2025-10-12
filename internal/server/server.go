package server

import (
	"log/slog"
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

		logger *slog.Logger
	}
)

func New(config *config.Config, logger *slog.Logger) (*Server, error) {
	server := &Server{}
	return server, nil
}

func (server *Server) Run() error {
	return nil
}
