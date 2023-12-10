package objectid

import "fmt"

func Fmt[T string | ID](lis []T, str string) []string {
	return Fnc(lis, func(x T) string { return fmt.Sprintf(str, x) })
}
