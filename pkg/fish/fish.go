package fish

import (
	"encoding/json"
	"fmt"
	"strings"

	set "github.com/deckarep/golang-set/v2"
)

type Fish struct {
	FinType     Fin
	SpeciesType Species
}

func (f *Fish) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	parts := strings.Split(strings.ToLower(s), " ")
	if len(parts) != 2 {
		return nil
	}

	f.FinType = GetFin(parts[0])
	f.SpeciesType = GetSpecies(parts[1])

	return nil
}

func (f Fish) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Name())
}

func FishFromString(s string) Fish {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return Fish{}
	}

	return Fish {
		FinType: GetFin(strings.ToLower(parts[0])),
		SpeciesType: GetSpecies(strings.ToLower(parts[1])),
	}
}

func (f Fish) Name() string {
	return fmt.Sprintf("%s %s", f.FinType.Name(), f.SpeciesType.Name())
}

func (f Fish) Breed(other Fish) Fish {
	return Fish{
		FinType:     f.FinType.Breed(other.FinType),
		SpeciesType: f.SpeciesType.Breed(other.SpeciesType),
	}
}

func (f Fish) GetPossibleParents() set.Set[Fish] {
	return set.NewThreadUnsafeSet[Fish]()
}

func GetPossibleOffspring(fish set.Set[Fish]) set.Set[Fish] {
	var possibleFish set.Set[Fish]
	slc := fish.ToSlice()
	for i := 0; i < len(slc)-1; i++ {
		f1 := slc[i]
		for j := i + 1; j < len(slc); j++ {
			f2 := slc[j]
			possibleFish.Add(f1.Breed(f2))
		}
	}
	return possibleFish
}
