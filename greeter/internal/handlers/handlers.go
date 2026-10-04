// Package handlers implements the greeter service's HTTP handlers.
package handlers

import (
	"encoding/json"
	"net/http"
)

// Greeting is the response body for GET /hello.
type Greeting struct {
	Name    string `json:"name,omitempty"`
	Message string `json:"message"`
}

// Health is the response body for GET /health.
type Health struct {
	Status string `json:"status"`
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

// GetGreeting handles GET /hello?name=X, returning a greeting addressed to
// name when provided, or a generic greeting otherwise. It never fails.
func GetGreeting(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		writeJSON(w, http.StatusOK, Greeting{Message: "Hello, World!"})
		return
	}
	writeJSON(w, http.StatusOK, Greeting{Name: name, Message: "Hello, " + name + "!"})
}

// GetHealth handles GET /health, a liveness check.
func GetHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, Health{Status: "ok"})
}
