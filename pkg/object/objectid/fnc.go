package objectid

func Fnc[T string | ID](lis []T, fnc func(T) string) []string {
	var key []string

	for _, x := range lis {
		key = append(key, fnc(x))
	}

	return key
}
