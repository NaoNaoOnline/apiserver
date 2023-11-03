package generic

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

// Uni returns the unique items of the given list.
func Uni[T string | objectid.ID | int64](lis []T) []T {
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
