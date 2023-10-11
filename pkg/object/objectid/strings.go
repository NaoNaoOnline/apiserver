package objectid

func Strings(ids []ID) []string {
	var str []string

	for _, x := range ids {
		str = append(str, x.String())
	}

	return str
}
