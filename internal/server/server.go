// Package server implements an HTTP server with health check endpoint.
package server

import (
	"fmt"
	"net/http"

	"github.com/tab-sama/go-tyk-template/internal/config"
)

// Server represents an HTTP server with ping functionality.
type Server struct {
	cfg    *config.Config
	server *http.Server
}

// NewServer creates a new Server instance.
func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
		server: &http.Server{
			Addr:        cfg.ListenAddress,
			ReadTimeout: cfg.TimeOut,
		},
	}
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	s.setupRoutes()

	err := s.server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("server: %w", err)
	}

	return nil
}

// pingHandler handles the /ping endpoint.
func (s *Server) pingHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte(`{"message":"pong"}`))
	if err != nil {
		panic(err)
	}
}

// setupRoutes configures the HTTP routes.
func (s *Server) setupRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", s.pingHandler)

	s.server.Handler = mux
}
