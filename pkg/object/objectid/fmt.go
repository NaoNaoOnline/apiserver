package objectid

import "fmt"

func Fmt[T string | ID](ids []T, str string) []string {
	var key []string

	for _, x := range ids {
		key = append(key, fmt.Sprintf(str, x))
	}

	return key
}
