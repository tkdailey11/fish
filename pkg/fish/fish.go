package fish

import (
	"fmt"
	"strings"
)

type Fish struct {
	fin     Fin
	species Species
}

func FishFromString(s string) Fish {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return Fish{}
	}

	return Fish {
		fin: GetFin(parts[0]),
		species: GetSpecies(parts[1]),
	}
}

func (f Fish) Name() string {
	return fmt.Sprintf("%s %s", f.fin.GetName(), f.species.GetName())
}

func (f Fish) Breed(other Fish) Fish {
	return Fish{
		fin:     f.fin.Breed(other.fin),
		species: f.species.Breed(other.species),
	}
}

func (f Fish) GetPossibleParents() []Fish {
	return []Fish{f}
}

func GetPossibleFish(fish []Fish) []Fish {
	var possibleFish []Fish
	for i := 0; i < len(fish)-1; i++ {
		f1 := fish[i]
		for j := i + 1; j < len(fish); j++ {
			f2 := fish[j]
			possibleFish = append(possibleFish, f1.Breed(f2))
		}
	}
	return possibleFish
}
