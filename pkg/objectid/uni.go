package objectid

func Uni[T string | String](lis []T) []T {
	see := map[T]struct{}{}

	var uni []T
	for _, x := range lis {
		{
			_, exi := see[x]
			if exi {
				continue
			}
			see[x] = struct{}{}
		}

		{
			uni = append(uni, x)
		}
	}

	return uni
}
