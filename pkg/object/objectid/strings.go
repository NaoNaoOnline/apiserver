package objectid

func Strings(ids []ID) []string {
	var lis []string

	for _, x := range ids {
		lis = append(lis, x.String())
	}

	return lis
}
