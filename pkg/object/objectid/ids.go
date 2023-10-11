package objectid

func IDs(str []string) []ID {
	var lis []ID

	for _, x := range str {
		lis = append(lis, ID(x))
	}

	return lis
}
