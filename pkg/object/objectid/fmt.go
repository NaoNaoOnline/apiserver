package objectid

import "fmt"

func Fmt[T string | ID](lis []T, str string) []string {
	var key []string

	for _, x := range lis {
		key = append(key, fmt.Sprintf(str, x))
	}

	return key
}
