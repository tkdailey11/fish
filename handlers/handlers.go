package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	set "github.com/deckarep/golang-set/v2"

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

	var fishList set.Set[fish.Fish]
	for _, str := range fishListStr {
		fishList.Add(fish.FishFromString(str))
	}

	allPossible, err := fish.GetPossibleOffspring(fishList).MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s", string(allPossible))
}

func (h *Handlers) PossibleParents(w http.ResponseWriter, r *http.Request) {
	var offspring string
	err := json.NewDecoder(r.Body).Decode(&offspring)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}


	babyFish := fish.FishFromString(offspring)
	parents := babyFish.GetPossibleParents()

	fmt.Fprintf(w, "%d %s", parents.Cardinality(), parents.String())

	// fishList, err := babyFish.GetPossibleParents().MarshalJSON()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// fmt.Fprintf(w, "%s", string(fishList))
}

func (h *Handlers) PossibleOffspringFins(w http.ResponseWriter, r *http.Request) {
	var finsListStr []string
	err := json.NewDecoder(r.Body).Decode(&finsListStr)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	var finsList set.Set[fish.Fin]
	for _, str := range finsListStr {
		finsList.Add(fish.GetFin(str))
	}

	allPossible, err := fish.GetPossibleOffspringFins(finsList).MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s", string(allPossible))
}

func (h *Handlers) PossibleParentsFins(w http.ResponseWriter, r *http.Request) {
	var offspring fish.Fin
	err := json.NewDecoder(r.Body).Decode(&offspring)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	finsList, err := offspring.GetPossibleParentsFins().MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s", string(finsList))
}

func (h *Handlers) PossibleOffspringSpecies(w http.ResponseWriter, r *http.Request) {
	var speciesListStr []string
	err := json.NewDecoder(r.Body).Decode(&speciesListStr)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	var speciesList set.Set[fish.Species]
	for _, str := range speciesListStr {
		speciesList.Add(fish.GetSpecies(str))
	}

	allPossible, err := fish.GetPossibleOffspringSpecies(speciesList).MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s", string(allPossible))
}

func (h *Handlers) PossibleParentsSpecies(w http.ResponseWriter, r *http.Request) {
	var offspring fish.Species
	err := json.NewDecoder(r.Body).Decode(&offspring)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	speciesList, err := offspring.GetPossibleParentsSpecies().MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s", string(speciesList))
}

func (h *Handlers) PossibleFins(w http.ResponseWriter, r *http.Request) {
	var finListStr []string
	err := json.NewDecoder(r.Body).Decode(&finListStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fins := make([]fish.Fin, len(finListStr))
	for i, f := range finListStr {
		fins[i] = fish.GetFin(f)
	}

	possibleFins := fish.GetPossibleFinStrings(fins)

	fmt.Fprintf(w, "Fins: %v", possibleFins)
}

func (h *Handlers) PossibleSpecies(w http.ResponseWriter, r *http.Request) {}
