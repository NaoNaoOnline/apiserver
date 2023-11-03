package objectid

import (
	"fmt"
	"strings"
)

func Frst(lis []string) []ID {
	var ids []ID

	for _, x := range lis {
		ids = append(ids, splt(x)[0])
	}

	return ids
}

func Pair(fir ID, sec ID) string {
	return fmt.Sprintf("%s,%s", fir, sec)
}

func splt(str string) [2]ID {
	spl := strings.Split(str, ",")
	return [2]ID{ID(spl[0]), ID(spl[1])}
}
