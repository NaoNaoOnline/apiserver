package generic

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

// Dup expresses whether the given list contains duplicated items.
func Dup[T string | objectid.ID | int64](lis []T) bool {
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
