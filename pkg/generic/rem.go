package generic

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

// Rem iterates over the given list removes all of the given subset. So if
// any item of sub can be found inside all, then Rem will filter it out.
func Rem[T string | objectid.ID | int64](all []T, sub []T) []T {
	if len(all) == 0 || len(sub) == 0 {
		return all
	}

	var rem []T
	for _, x := range all {
		var exi bool

		for _, y := range sub {
			if x == y {
				exi = true
				break
			}
		}

		if !exi {
			rem = append(rem, x)
		}
	}

	return rem
}
