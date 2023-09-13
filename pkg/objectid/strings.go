package objectid

func Strings(str []string) []String {
	var ids []String

	for _, x := range str {
		ids = append(ids, String(x))
	}

	return ids
}
