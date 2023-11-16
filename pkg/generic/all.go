package generic

// All expresses whether the given list contains all of the given subset. So if
// any item of sub cannot be found inside all, then All returns false.
func All[T string | int64](all []T, sub []T) bool {
	if len(all) == 0 || len(sub) == 0 {
		return false
	}

	for _, x := range sub {
		var exi bool

		for _, y := range all {
			if x == y {
				exi = true
				break
			}
		}

		if !exi {
			return false
		}
	}

	return true
}
