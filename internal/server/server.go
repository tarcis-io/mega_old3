package server

import (
	"net/http"
	"time"
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
