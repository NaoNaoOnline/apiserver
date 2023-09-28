package objectid

func Strings(str []string) []ID {
	var ids []ID

	for _, x := range str {
		ids = append(ids, ID(x))
	}

	return ids
}
