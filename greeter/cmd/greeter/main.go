// Command greeter runs the greeter HTTP service.
package main

import (
	"log"
	"net/http"

	"greeter/internal/server"
)

func main() {
	addr := ":9090"
	log.Printf("greeter listening on %s", addr)
	if err := http.ListenAndServe(addr, server.New()); err != nil {
		log.Fatalf("greeter: %v", err)
	}
}
