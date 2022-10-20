package fish

import (
	"encoding/json"
	"strings"

	set "github.com/deckarep/golang-set/v2"
)

type Fin int
type FinParent struct {
	Parent1 Fin
	Parent2 Fin
}

func (f *Fin) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	fTemp := GetFin(s)
	f = &fTemp

	return nil
}

func (f Fin) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Name())
}

const (
	bluetip   Fin = 0b000000000000000000001
	canary    Fin = 0b000000000000000000010
	crimson   Fin = 0b000000000000000000100
	daffodil  Fin = 0b000000000000000001000
	greenfin  Fin = 0b000000000000000010000
	hooked    Fin = 0b000000000000000100000
	imperial  Fin = 0b000000000000001000000
	orange    Fin = 0b000000000000010000000
	oriental  Fin = 0b000000000000100000000
	peachy    Fin = 0b000000000001000000000
	pink      Fin = 0b000000000010000000000
	razorback Fin = 0b000000000100000000000
	royal     Fin = 0b000000001000000000000
	rusty     Fin = 0b000000010000000000000
	sapphire  Fin = 0b000000100000000000000
	serrated  Fin = 0b000001000000000000000
	silky     Fin = 0b000010000000000000000
	snubbed   Fin = 0b000100000000000000000
	spined    Fin = 0b001000000000000000000
	striped   Fin = 0b010000000000000000000
	tangerine Fin = 0b100000000000000000000
	unknownF  Fin = 0
)

func (f Fin) Name() string {
	switch f.GetBase() {
	case bluetip:
		return "Bluetip"
	case canary:
		return "Canary"
	case crimson:
		return "Crimson"
	case daffodil:
		return "Daffodil"
	case greenfin:
		return "Greenfin"
	case hooked:
		return "Hooked"
	case imperial:
		return "Imperial"
	case orange:
		return "Orange"
	case oriental:
		return "Oriental"
	case peachy:
		return "Peachy"
	case pink:
		return "Pink"
	case razorback:
		return "Razorback"
	case royal:
		return "Royal"
	case rusty:
		return "Rusty"
	case sapphire:
		return "Sapphire"
	case serrated:
		return "Serrated"
	case silky:
		return "Silky"
	case snubbed:
		return "Snubbed"
	case spined:
		return "Spined"
	case striped:
		return "Striped"
	case tangerine:
		return "Tangerine"
	default:
		return "Unknown"
	}
}

func GetFin(name string) Fin {
	switch strings.ToLower(name) {
	case "bluetip":
		return bluetip
	case "canary":
		return canary
	case "crimson":
		return crimson
	case "daffodil":
		return daffodil
	case "greenfin":
		return greenfin
	case "hooked":
		return hooked
	case "imperial":
		return imperial
	case "orange":
		return orange
	case "oriental":
		return oriental
	case "peachy":
		return peachy
	case "pink":
		return pink
	case "razorback":
		return razorback
	case "royal":
		return royal
	case "rusty":
		return rusty
	case "sapphire":
		return sapphire
	case "serrated":
		return serrated
	case "silky":
		return silky
	case "snubbed":
		return snubbed
	case "spined":
		return spined
	case "striped":
		return striped
	case "tangerine":
		return tangerine
	default:
		return unknownF
	}
}

func (f Fin) Breed(other Fin) Fin {
	baby := f | other
	return baby.GetBase()
}

func (f Fin) GetBase() Fin {
	switch f {
	case bluetip, bluetip | spined, greenfin | spined:
		return bluetip
	case canary, canary | oriental, crimson | imperial, crimson | peachy, daffodil | sapphire, daffodil | striped, hooked | rusty, hooked | tangerine, imperial | sapphire, orange | silky, oriental | razorback, peachy | tangerine, pink | serrated, pink | snubbed, razorback | royal, royal | snubbed, rusty | silky, serrated | striped:
		return canary
	case crimson, bluetip | oriental, bluetip | royal, canary | spined, crimson | tangerine, greenfin | pink, greenfin | royal, orange | pink, orange | sapphire, orange | striped, oriental | spined, rusty | sapphire, rusty | striped, sapphire | tangerine:
		return crimson
	case daffodil, canary | hooked, canary | silky, daffodil | serrated, hooked | razorback, imperial | serrated, imperial | snubbed, oriental | silky, peachy | razorback, peachy | snubbed:
		return daffodil
	case greenfin, bluetip | greenfin, bluetip | orange, orange | spined, rusty | spined:
		return greenfin
	case hooked, hooked | peachy, imperial | silky, peachy | silky:
		return hooked
	case imperial, daffodil | imperial, daffodil | peachy, hooked | serrated, hooked | snubbed, peachy | serrated, razorback | silky, silky | snubbed:
		return imperial
	case orange, bluetip | rusty, bluetip | tangerine, crimson | spined, greenfin | orange, greenfin | rusty, spined | tangerine:
		return orange
	case oriental, bluetip | silky, canary | pink, canary | royal, crimson | daffodil, crimson | serrated, daffodil | tangerine, greenfin | hooked, greenfin | silky, hooked | orange, imperial | rusty, imperial | tangerine, orange | peachy, oriental | royal, peachy | rusty, pink | razorback, razorback | striped, sapphire | serrated, sapphire | snubbed, snubbed | striped:
		return oriental
	case peachy, daffodil | hooked, daffodil | silky, hooked | imperial, imperial | peachy, serrated | silky:
		return peachy
	case pink, bluetip | daffodil, canary | crimson, canary | tangerine, crimson | oriental, daffodil | greenfin, imperial | spined, orange | serrated, orange | snubbed, oriental | sapphire, peachy | spined, pink | striped, razorback | rusty, razorback | tangerine, royal | sapphire, royal | striped, rusty | snubbed:
		return pink
	case razorback, canary | razorback, canary | snubbed, crimson | hooked, crimson | silky, daffodil | pink, daffodil | royal, hooked | sapphire, imperial | pink, imperial | striped, oriental | serrated, oriental | snubbed, peachy | sapphire, peachy | striped, royal | serrated, silky | tangerine:
		return razorback
	case royal, bluetip | hooked, bluetip | peachy, canary | sapphire, canary | striped, crimson | razorback, crimson | snubbed, daffodil | orange, daffodil | rusty, greenfin | imperial, greenfin | peachy, hooked | spined, imperial | orange, oriental | pink, oriental | striped, pink | royal, razorback | sapphire, rusty | serrated, serrated | tangerine, silky | spined, snubbed | tangerine:
		return royal
	case rusty, bluetip | crimson, bluetip | sapphire, crimson | greenfin, greenfin | tangerine, orange | rusty, orange | tangerine, sapphire | spined, spined | striped:
		return rusty
	case sapphire, bluetip | canary, bluetip | razorback, canary | greenfin, crimson | sapphire, crimson | striped, greenfin | oriental, orange | oriental, orange | royal, pink | rusty, pink | tangerine, razorback | spined, royal | rusty, snubbed | spined, striped | tangerine:
		return sapphire
	case serrated, canary | imperial, canary | peachy, daffodil | razorback, daffodil | snubbed, hooked | oriental, hooked | royal, imperial | razorback, oriental | peachy, pink | silky, royal | silky, serrated | snubbed:
		return serrated
	case silky, hooked | silky:
		return silky
	case snubbed, canary | daffodil, canary | serrated, daffodil | oriental, hooked | pink, hooked | striped, imperial | oriental, imperial | royal, peachy | pink, peachy | royal, razorback | serrated, razorback | snubbed, sapphire | silky, silky | striped:
		return snubbed
	case spined:
		return spined
	case striped, bluetip | imperial, bluetip | serrated, bluetip | snubbed, canary | orange, canary | rusty, crimson | pink, crimson | royal, daffodil | spined, greenfin | razorback, greenfin | serrated, greenfin | snubbed, orange | razorback, oriental | rusty, oriental | tangerine, pink | sapphire, royal | tangerine, sapphire | striped, serrated | spined:
		return striped
	case tangerine, bluetip | pink, bluetip | striped, crimson | orange, crimson | rusty, greenfin | sapphire, greenfin | striped, pink | spined, royal | spined, rusty | tangerine:
		return tangerine
	default:
		return 0
	}
}

func (f Fin) GetPossibleParentsFins() set.Set[FinParent] {
	parents := set.NewThreadUnsafeSet[FinParent]()
	switch f {
	case bluetip:
		parents.Add(FinParent{Parent1: bluetip, Parent2: bluetip})
		parents.Add(FinParent{Parent1: bluetip, Parent2: spined})
		parents.Add(FinParent{Parent1: greenfin, Parent2: spined})
		break
	case canary:
		parents.Add(FinParent{Parent1: canary, Parent2: canary})
		parents.Add(FinParent{Parent1: canary, Parent2: oriental})
		parents.Add(FinParent{Parent1: crimson, Parent2: imperial})
		parents.Add(FinParent{Parent1: crimson, Parent2: peachy})
		parents.Add(FinParent{Parent1: daffodil, Parent2: sapphire})
		parents.Add(FinParent{Parent1: daffodil, Parent2: striped})
		parents.Add(FinParent{Parent1: hooked, Parent2: rusty})
		parents.Add(FinParent{Parent1: hooked, Parent2: tangerine})
		parents.Add(FinParent{Parent1: imperial, Parent2: sapphire})
		parents.Add(FinParent{Parent1: orange, Parent2: silky})
		parents.Add(FinParent{Parent1: oriental, Parent2: razorback})
		parents.Add(FinParent{Parent1: peachy, Parent2: tangerine})
		parents.Add(FinParent{Parent1: pink, Parent2: serrated})
		parents.Add(FinParent{Parent1: pink, Parent2: snubbed})
		parents.Add(FinParent{Parent1: razorback, Parent2: royal})
		parents.Add(FinParent{Parent1: royal, Parent2: snubbed})
		parents.Add(FinParent{Parent1: rusty, Parent2: silky})
		parents.Add(FinParent{Parent1: serrated, Parent2: striped})
		break
	case crimson:
		parents.Add(FinParent{Parent1: crimson, Parent2: crimson})
		parents.Add(FinParent{Parent1: bluetip, Parent2: oriental})
		parents.Add(FinParent{Parent1: bluetip, Parent2: royal})
		parents.Add(FinParent{Parent1: canary, Parent2: spined})
		parents.Add(FinParent{Parent1: crimson, Parent2: tangerine})
		parents.Add(FinParent{Parent1: greenfin, Parent2: pink})
		parents.Add(FinParent{Parent1: greenfin, Parent2: royal})
		parents.Add(FinParent{Parent1: orange, Parent2: pink})
		parents.Add(FinParent{Parent1: orange, Parent2: sapphire})
		parents.Add(FinParent{Parent1: orange, Parent2: striped})
		parents.Add(FinParent{Parent1: oriental, Parent2: spined})
		parents.Add(FinParent{Parent1: rusty, Parent2: sapphire})
		parents.Add(FinParent{Parent1: rusty, Parent2: striped})
		parents.Add(FinParent{Parent1: sapphire, Parent2: tangerine})
		break
	case daffodil:
		parents.Add(FinParent{Parent1: daffodil, Parent2: daffodil})
		parents.Add(FinParent{Parent1: canary, Parent2: hooked})
		parents.Add(FinParent{Parent1: canary, Parent2: silky})
		parents.Add(FinParent{Parent1: daffodil, Parent2: serrated})
		parents.Add(FinParent{Parent1: hooked, Parent2: razorback})
		parents.Add(FinParent{Parent1: imperial, Parent2: serrated})
		parents.Add(FinParent{Parent1: imperial, Parent2: snubbed})
		parents.Add(FinParent{Parent1: oriental, Parent2: silky})
		parents.Add(FinParent{Parent1: peachy, Parent2: razorback})
		parents.Add(FinParent{Parent1: peachy, Parent2: snubbed})
		break
	case greenfin:
		parents.Add(FinParent{Parent1: greenfin, Parent2: greenfin})
		parents.Add(FinParent{Parent1: bluetip, Parent2: greenfin})
		parents.Add(FinParent{Parent1: bluetip, Parent2: orange})
		parents.Add(FinParent{Parent1: orange, Parent2: spined})
		parents.Add(FinParent{Parent1: rusty, Parent2: spined})
		break
	case hooked:
		parents.Add(FinParent{Parent1: hooked, Parent2: hooked})
		parents.Add(FinParent{Parent1: hooked, Parent2: peachy})
		parents.Add(FinParent{Parent1: imperial, Parent2: silky})
		parents.Add(FinParent{Parent1: peachy, Parent2: silky})
		break
	case imperial:
		parents.Add(FinParent{Parent1: imperial, Parent2: imperial})
		parents.Add(FinParent{Parent1: daffodil, Parent2: imperial})
		parents.Add(FinParent{Parent1: daffodil, Parent2: peachy})
		parents.Add(FinParent{Parent1: hooked, Parent2: serrated})
		parents.Add(FinParent{Parent1: hooked, Parent2: snubbed})
		parents.Add(FinParent{Parent1: peachy, Parent2: serrated})
		parents.Add(FinParent{Parent1: razorback, Parent2: silky})
		parents.Add(FinParent{Parent1: silky, Parent2: snubbed})
		break
	case orange:
		parents.Add(FinParent{Parent1: orange, Parent2: orange})
		parents.Add(FinParent{Parent1: bluetip, Parent2: rusty})
		parents.Add(FinParent{Parent1: bluetip, Parent2: tangerine})
		parents.Add(FinParent{Parent1: crimson, Parent2: spined})
		parents.Add(FinParent{Parent1: greenfin, Parent2: orange})
		parents.Add(FinParent{Parent1: greenfin, Parent2: rusty})
		parents.Add(FinParent{Parent1: spined, Parent2: tangerine})
		break
	case oriental:
		parents.Add(FinParent{Parent1: oriental, Parent2: oriental})
		parents.Add(FinParent{Parent1: bluetip, Parent2: silky})
		parents.Add(FinParent{Parent1: canary, Parent2: pink})
		parents.Add(FinParent{Parent1: canary, Parent2: royal})
		parents.Add(FinParent{Parent1: crimson, Parent2: daffodil})
		parents.Add(FinParent{Parent1: crimson, Parent2: serrated})
		parents.Add(FinParent{Parent1: daffodil, Parent2: tangerine})
		parents.Add(FinParent{Parent1: greenfin, Parent2: hooked})
		parents.Add(FinParent{Parent1: greenfin, Parent2: silky})
		parents.Add(FinParent{Parent1: hooked, Parent2: orange})
		parents.Add(FinParent{Parent1: imperial, Parent2: rusty})
		parents.Add(FinParent{Parent1: imperial, Parent2: tangerine})
		parents.Add(FinParent{Parent1: orange, Parent2: peachy})
		parents.Add(FinParent{Parent1: oriental, Parent2: royal})
		parents.Add(FinParent{Parent1: peachy, Parent2: rusty})
		parents.Add(FinParent{Parent1: pink, Parent2: razorback})
		parents.Add(FinParent{Parent1: razorback, Parent2: striped})
		parents.Add(FinParent{Parent1: sapphire, Parent2: serrated})
		parents.Add(FinParent{Parent1: sapphire, Parent2: snubbed})
		parents.Add(FinParent{Parent1: snubbed, Parent2: striped})
		break
	case peachy:
		parents.Add(FinParent{Parent1: peachy, Parent2: peachy})
		parents.Add(FinParent{Parent1: daffodil, Parent2: hooked})
		parents.Add(FinParent{Parent1: daffodil, Parent2: silky})
		parents.Add(FinParent{Parent1: hooked, Parent2: imperial})
		parents.Add(FinParent{Parent1: imperial, Parent2: peachy})
		parents.Add(FinParent{Parent1: serrated, Parent2: silky})
		break
	case pink, bluetip | daffodil, canary | crimson, canary | tangerine, crimson | oriental, daffodil | greenfin, imperial | spined, orange | serrated, orange | snubbed, oriental | sapphire, peachy | spined, pink | striped, razorback | rusty, razorback | tangerine, royal | sapphire, royal | striped, rusty | snubbed:
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		return pink
	case razorback, canary | razorback, canary | snubbed, crimson | hooked, crimson | silky, daffodil | pink, daffodil | royal, hooked | sapphire, imperial | pink, imperial | striped, oriental | serrated, oriental | snubbed, peachy | sapphire, peachy | striped, royal | serrated, silky | tangerine:
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		return razorback
	case royal, bluetip | hooked, bluetip | peachy, canary | sapphire, canary | striped, crimson | razorback, crimson | snubbed, daffodil | orange, daffodil | rusty, greenfin | imperial, greenfin | peachy, hooked | spined, imperial | orange, oriental | pink, oriental | striped, pink | royal, razorback | sapphire, rusty | serrated, serrated | tangerine, silky | spined, snubbed | tangerine:
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		return royal
	case rusty:
		parents.Add(FinParent{Parent1: rusty, Parent2: rusty})
		parents.Add(FinParent{Parent1: bluetip, Parent2: crimson})
		parents.Add(FinParent{Parent1: bluetip, Parent2: sapphire})
		parents.Add(FinParent{Parent1: crimson, Parent2: greenfin})
		parents.Add(FinParent{Parent1: greenfin, Parent2: tangerine})
		parents.Add(FinParent{Parent1: orange, Parent2: rusty})
		parents.Add(FinParent{Parent1: orange, Parent2: tangerine})
		parents.Add(FinParent{Parent1: sapphire, Parent2: spined})
		parents.Add(FinParent{Parent1: spined, Parent2: striped})
		break
	case sapphire, bluetip | canary, bluetip | razorback, canary | greenfin, crimson | sapphire, crimson | striped, greenfin | oriental, orange | oriental, orange | royal, pink | rusty, pink | tangerine, razorback | spined, royal | rusty, snubbed | spined, striped | tangerine:
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		return sapphire
	case serrated, canary | imperial, canary | peachy, daffodil | razorback, daffodil | snubbed, hooked | oriental, hooked | royal, imperial | razorback, oriental | peachy, pink | silky, royal | silky, serrated | snubbed:
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		return serrated
	case silky:
		parents.Add(FinParent{Parent1: silky, Parent2: silky})
		parents.Add(FinParent{Parent1: hooked, Parent2: silky})
		break
	case snubbed, canary | daffodil, canary | serrated, daffodil | oriental, hooked | pink, hooked | striped, imperial | oriental, imperial | royal, peachy | pink, peachy | royal, razorback | serrated, razorback | snubbed, sapphire | silky, silky | striped:
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		return snubbed
	case spined:
		parents.Add(FinParent{Parent1: spined, Parent2: spined})
		break
	case striped, bluetip | imperial, bluetip | serrated, bluetip | snubbed, canary | orange, canary | rusty, crimson | pink, crimson | royal, daffodil | spined, greenfin | razorback, greenfin | serrated, greenfin | snubbed, orange | razorback, oriental | rusty, oriental | tangerine, pink | sapphire, royal | tangerine, sapphire | striped, serrated | spined:
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		parents.Add(FinParent{Parent1: 1, Parent2: 1})
		return striped
	case tangerine:
		parents.Add(FinParent{Parent1: tangerine, Parent2: tangerine})
		parents.Add(FinParent{Parent1: bluetip, Parent2: pink})
		parents.Add(FinParent{Parent1: bluetip, Parent2: striped})
		parents.Add(FinParent{Parent1: crimson, Parent2: orange})
		parents.Add(FinParent{Parent1: crimson, Parent2: rusty})
		parents.Add(FinParent{Parent1: greenfin, Parent2: sapphire})
		parents.Add(FinParent{Parent1: greenfin, Parent2: striped})
		parents.Add(FinParent{Parent1: pink, Parent2: spined})
		parents.Add(FinParent{Parent1: royal, Parent2: spined})
		parents.Add(FinParent{Parent1: rusty, Parent2: tangerine})
		break
	default:
		break
	}

	return parents
}

func GetPossibleOffspringFins(fins set.Set[Fin]) set.Set[Fin] {
	var possibleFins set.Set[Fin]
	slc := fins.ToSlice()
	for i := 0; i < len(slc)-1; i++ {
		f1 := slc[i]
		for j := i + 1; j < len(slc); j++ {
			f2 := slc[j]
			possibleFins.Add(f1.Breed(f2))
		}
	}
	return possibleFins
}
