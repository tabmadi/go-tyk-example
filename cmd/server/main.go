// Package main is the entry point for the application.
package main

import (
	"log"

	"github.com/tab-sama/go-tyk-template/internal/config"
	"github.com/tab-sama/go-tyk-template/internal/server"
)

func main() {
	cfg := config.Load()

	log.Printf("Starting server on %s", cfg.ListenAddress)

	srv := server.NewServer(cfg)

	err := srv.Start()
	if err != nil {
		panic(err)
	}
}
