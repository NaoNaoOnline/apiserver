package objectid

// Dup returns whether the given list contains duplicates.
func Dup[T string | ID](lis []T) bool {
	see := map[T]struct{}{}

	for _, x := range lis {
		{
			_, exi := see[x]
			if exi {
				return true
			}
		}

		{
			see[x] = struct{}{}
		}
	}

	return false
}
