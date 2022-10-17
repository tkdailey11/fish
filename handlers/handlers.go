package handlers

import (
	"net/http"

	"github.com/tkdailey11/fish/pkg/fish"
)

// Handlers base struct
type Handlers struct{}

type OffspringReq struct {
	parent1 Fish
	parent2 string
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) PossibleOffspring(w http.ResponseWriter, r *http.Request) {
}

func (h *Handlers) PossibleParents(w http.ResponseWriter, r *http.Request) {
}
