package objectid

// Uni returns a list of unique items.
func Uni[T string | ID](lis []T) []T {
	see := map[T]struct{}{}

	var uni []T
	for _, x := range lis {
		{
			_, exi := see[x]
			if exi {
				continue
			}
		}

		{
			see[x] = struct{}{}
		}

		{
			uni = append(uni, x)
		}
	}

	return uni
}
