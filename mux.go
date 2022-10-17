package main

import (
	"net/http"

	"github.com/tkdailey11/fish/handlers"
)

// Routes configuration
func initRoutes() {

	// Constructing handlers struct
	h := handlers.NewHandlers()

	// Binding Functions with Handlers
	http.HandleFunc("/api/offspring", h.PossibleOffspring)
	http.HandleFunc("/api/parents", h.PossibleParents)
}
