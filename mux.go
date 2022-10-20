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

	http.HandleFunc("/api/fin/offspring", h.PossibleOffspringFins)
	http.HandleFunc("/api/fin/parents", h.PossibleParentsFins)
	
	http.HandleFunc("/api/species/offspring", h.PossibleOffspringSpecies)
	http.HandleFunc("/api/species/parents", h.PossibleParentsSpecies)
}
