package generic

// Any expresses whether the given list contains any of the given subset. So if
// any item of sub can be found inside all, then Any returns true.
func Any[T string | int64](all []T, sub []T) bool {
	for _, x := range all {
		for _, y := range sub {
			if x == y {
				return true
			}
		}
	}

	return false
}
