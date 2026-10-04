// Package server wires the greeter service's routes.
package server

import (
	"net/http"

	"greeter/internal/handlers"
)

// New builds the greeter service's HTTP handler.
func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handlers.GetGreeting)
	mux.HandleFunc("GET /health", handlers.GetHealth)
	return mux
}
