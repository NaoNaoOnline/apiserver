package scoreid

import "fmt"

func Fmt[T string | String](ids []T, str string) []string {
	var key []string

	for _, x := range ids {
		key = append(key, fmt.Sprint(str, x))
	}

	return key
}
