package fish

import (
	"encoding/json"
	"strings"

	set "github.com/deckarep/golang-set/v2"
)

type Species int

func (s *Species) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	sTemp := GetSpecies(str)
	s = &sTemp

	return nil
}

func (s Species) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Name())
}

const (
	angelfish  Species = 0b000000000000000000001
	arrowfish  Species = 0b000000000000000000010
	betta      Species = 0b000000000000000000100
	carp       Species = 0b000000000000000001000
	catfish    Species = 0b000000000000000010000
	comet      Species = 0b000000000000000100000
	fatfish    Species = 0b000000000000001000000
	flashfish  Species = 0b000000000000010000000
	foxface    Species = 0b000000000000100000000
	goldfish   Species = 0b000000000001000000000
	goldshark  Species = 0b000000000010000000000
	leaffish   Species = 0b000000000100000000000
	pufferfish Species = 0b000000001000000000000
	pygmy      Species = 0b000000010000000000000
	rainbow    Species = 0b000000100000000000000
	shark      Species = 0b000001000000000000000
	snooper    Species = 0b000010000000000000000
	snout      Species = 0b000100000000000000000
	spotanus   Species = 0b001000000000000000000
	stickfish  Species = 0b010000000000000000000
	tigerfish  Species = 0b100000000000000000000
	unknownS   Species = 0
)

func (s Species) Name() string {
	switch s.GetBase() {
	case angelfish:
		return "Angelfish"
	case arrowfish:
		return "Arrowfish"
	case betta:
		return "Betta"
	case carp:
		return "Carp"
	case catfish:
		return "Catfish"
	case comet:
		return "Comet"
	case fatfish:
		return "Fatfish"
	case flashfish:
		return "Flashfish"
	case foxface:
		return "Foxface"
	case goldfish:
		return "Goldfish"
	case goldshark:
		return "Goldshark"
	case leaffish:
		return "Leaffish"
	case pufferfish:
		return "Pufferfish"
	case pygmy:
		return "Pygmy"
	case rainbow:
		return "Rainbow"
	case shark:
		return "Shark"
	case snooper:
		return "Snooper"
	case snout:
		return "Snout"
	case spotanus:
		return "Spotanus"
	case stickfish:
		return "Stickfish"
	case tigerfish:
		return "Tigerfish"
	default:
		return "Unknown"
	}
}

func GetSpecies(name string) Species {
	switch strings.ToLower(name) {
	case "angelfish":
		return angelfish
	case "arrowfish":
		return arrowfish
	case "betta":
		return betta
	case "carp":
		return carp
	case "catfish":
		return catfish
	case "comet":
		return comet
	case "fatfish":
		return fatfish
	case "flashfish":
		return flashfish
	case "foxface":
		return foxface
	case "goldfish":
		return goldfish
	case "goldshark":
		return goldshark
	case "leaffish":
		return leaffish
	case "pufferfish":
		return pufferfish
	case "pygmy":
		return pygmy
	case "rainbow":
		return rainbow
	case "shark":
		return shark
	case "snooper":
		return snooper
	case "snout":
		return snout
	case "spotanus":
		return spotanus
	case "stickfish":
		return stickfish
	case "tigerfish":
		return tigerfish
	default:
		return unknownS
	}
}

func (s Species) Breed(other Species) Species {
	baby := s | other
	return baby.GetBase()
}

func (s Species) GetBase() Species {
	switch s {
	case angelfish, angelfish | tigerfish, arrowfish | pufferfish, arrowfish | stickfish, carp | shark, catfish | leaffish, catfish | pufferfish, comet | fatfish, comet | snout, fatfish | rainbow, flashfish | goldshark, flashfish | snooper, goldshark | tigerfish, leaffish | shark, rainbow | snooper, snout | stickfish:
		return angelfish
	case arrowfish, arrowfish | snout, catfish | fatfish, catfish | snout, fatfish | shark, shark | snooper:
		return arrowfish
	case betta, betta | pygmy, carp | goldfish, carp | pygmy, foxface | leaffish, foxface | pufferfish, goldfish | leaffish, pufferfish | spotanus, spotanus | stickfish:
		return betta
	case carp, betta | carp, betta | leaffish, comet | foxface, comet | spotanus, foxface | stickfish, goldfish | pufferfish, goldfish | stickfish, leaffish | pygmy, pufferfish | pygmy, rainbow | spotanus:
		return carp
	case catfish, arrowfish | catfish, arrowfish | shark, shark | snout:
		return catfish
	case comet, angelfish | betta, angelfish | carp, arrowfish | spotanus, betta | goldshark, carp | tigerfish, comet | stickfish, fatfish | foxface, fatfish | goldfish, flashfish | leaffish, flashfish | pufferfish, foxface | snout, goldfish | snooper, goldshark | pygmy, leaffish | tigerfish, pufferfish | rainbow, pygmy | snooper, rainbow | stickfish, snout | spotanus:
		return comet
	case fatfish, angelfish | arrowfish, angelfish | catfish, arrowfish | goldshark, catfish | tigerfish, fatfish | snooper, flashfish | shark, goldshark | snout, shark | tigerfish, snooper | snout:
		return fatfish
	case flashfish, angelfish | comet, angelfish | stickfish, arrowfish | betta, arrowfish | pygmy, betta | snout, carp | fatfish, carp | snout, catfish | goldfish, catfish | pygmy, comet | tigerfish, fatfish | leaffish, flashfish | rainbow, foxface | shark, goldfish | shark, goldshark | pufferfish, goldshark | stickfish, leaffish | snooper, pufferfish | snooper, rainbow | tigerfish:
		return flashfish
	case foxface, foxface | spotanus, goldfish | spotanus:
		return foxface
	case goldfish, betta | spotanus, foxface | goldfish, foxface | pygmy, pygmy | spotanus:
		return goldfish
	case goldshark, angelfish | goldshark, angelfish | snooper, arrowfish | comet, arrowfish | rainbow, catfish | comet, catfish | stickfish, fatfish | flashfish, fatfish | tigerfish, flashfish | snout, pufferfish | shark, rainbow | snout, shark | stickfish, snooper | tigerfish:
		return goldshark
	case leaffish, betta | pufferfish, betta | stickfish, carp | leaffish, carp | pufferfish, comet | goldfish, comet | pygmy, flashfish | foxface, flashfish | spotanus, foxface | rainbow, goldfish | rainbow, pygmy | stickfish, spotanus | tigerfish:
		return leaffish
	case pufferfish, angelfish | foxface, angelfish | spotanus, betta | comet, betta | rainbow, carp | comet, carp | stickfish, flashfish | goldfish, flashfish | pygmy, foxface | tigerfish, goldfish | tigerfish, goldshark | spotanus, leaffish | pufferfish, leaffish | stickfish, pygmy | rainbow:
		return pufferfish
	case pygmy, betta | foxface, betta | goldfish, carp | foxface, carp | spotanus, goldfish | pygmy, leaffish | spotanus:
		return pygmy
	case rainbow, angelfish | leaffish, angelfish | pufferfish, arrowfish | foxface, arrowfish | goldfish, betta | fatfish, betta | snooper, carp | goldshark, carp | snooper, catfish | foxface, catfish | spotanus, comet | flashfish, comet | rainbow, fatfish | pygmy, flashfish | stickfish, goldfish | snout, goldshark | leaffish, pufferfish | tigerfish, pygmy | snout, shark | spotanus, stickfish | tigerfish:
		return rainbow
	case shark, catfish | shark:
		return shark
	case snooper, angelfish | fatfish, angelfish | snout, arrowfish | flashfish, arrowfish | tigerfish, catfish | flashfish, catfish | rainbow, comet | shark, fatfish | goldshark, goldshark | snooper, rainbow | shark, snout | tigerfish:
		return snooper
	case snout, angelfish | shark, arrowfish | fatfish, arrowfish | snooper, catfish | goldshark, catfish | snooper, fatfish | snout, goldshark | shark:
		return snout
	case spotanus:
		return spotanus
	case stickfish, angelfish | goldfish, angelfish | pygmy, betta | flashfish, betta | tigerfish, carp | flashfish, carp | rainbow, comet | leaffish, comet | pufferfish, fatfish | spotanus, foxface | goldshark, foxface | snooper, goldfish | goldshark, leaffish | rainbow, pufferfish | stickfish, pygmy | tigerfish, snooper | spotanus:
		return stickfish
	case tigerfish, angelfish | flashfish, angelfish | rainbow, arrowfish | carp, arrowfish | leaffish, betta | catfish, betta | shark, carp | catfish, comet | goldshark, comet | snooper, fatfish | pufferfish, fatfish | stickfish, flashfish | tigerfish, goldshark | rainbow, leaffish | snout, pufferfish | snout, pygmy | shark, snooper | stickfish:
		return tigerfish
	default:
		return 0
	}
}

func (s Species) GetPossibleParentsSpecies() set.Set[Species] {
	return set.NewSet[Species]()
}

func GetPossibleOffspringSpecies(species set.Set[Species]) set.Set[Species] {
	var possibleSpecies set.Set[Species]
	slc := species.ToSlice()
	for i := 0; i < len(slc)-1; i++ {
		s1 := slc[i]
		for j := i + 1; j < len(slc); j++ {
			s2 := slc[j]
			possibleSpecies.Add(s1.Breed(s2))
		}
	}
	return possibleSpecies
}
