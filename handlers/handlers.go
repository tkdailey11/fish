package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/tkdailey11/fish/pkg/fish"
)

// Handlers base struct
type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) PossibleOffspring(w http.ResponseWriter, r *http.Request) {
	var fishListStr []string
	err := json.NewDecoder(r.Body).Decode(&fishListStr)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	var fishList []fish.Fish
	for _, str := range fishListStr {
		fishList = append(fishList, fish.FishFromString(str))
	}

	allPossible := fish.GetPossibleFish(fishList)

	result := "["

	for _, possible := range allPossible {
		result = fmt.Sprintf("%s\"%s\",", result, possible.Name())
	}

	result = strings.TrimSuffix(result, ",")
	result = result + "]"

	fmt.Fprintf(w, "%s", result)
}

func (h *Handlers) PossibleParents(w http.ResponseWriter, r *http.Request) {
	var offspring fish.Fish
	err := json.NewDecoder(r.Body).Decode(&offspring)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	fishList := offspring.GetPossibleParents()

	result := "["

	for _, f := range fishList {
		result = fmt.Sprintf("%s\"%s\",", result, f.Name())
	}

	result = strings.TrimSuffix(result, ",")
	result = result + "]"

	fmt.Fprintf(w, "%s", result)
}
