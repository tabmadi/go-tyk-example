// Package server implements an HTTP server with echo endpoint.
package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/tab-sama/go-tyk-template/internal/config"
)

// Server represents an HTTP server with echo functionality.
type Server struct {
	cfg *config.Config
	app *fiber.App
}

// EchoRequest represents the request body for echo endpoint.
type EchoRequest struct {
	Message string `json:"message"`
}

// EchoResponse represents the response body for the echo endpoint.
type EchoResponse struct {
	Message string `json:"message"`
}

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewServer creates a new Server instance.
func NewServer(cfg *config.Config) *Server {
	app := fiber.New(fiber.Config{
		ReadTimeout: cfg.TimeOut,
	})

	return &Server{
		cfg: cfg,
		app: app,
	}
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	s.setup()

	err := s.app.Listen(s.cfg.ListenAddress)
	if err != nil {
		return fmt.Errorf("server: %w", err)
	}

	return nil
}

// echoHandler handles the /echo endpoint.
func (s *Server) echoHandler(ctx *fiber.Ctx) error {
	var req EchoRequest

	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&ErrorResponse{
			Error: "Invalid JSON format",
		})
	}

	response := &EchoResponse{
		Message: req.Message,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// setup configures the middlewares and HTTP routes.
func (s *Server) setup() {
	s.app.Use(logger.New())
	s.app.Use(recover.New())

	s.app.Post("/echo", s.echoHandler)
}
