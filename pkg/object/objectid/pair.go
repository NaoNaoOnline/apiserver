package objectid

import (
	"fmt"
	"strings"
)

func Frst(lis []string) []ID {
	var ids []ID

	for _, x := range lis {
		spl := strings.Split(x, ",")
		ids = append(ids, ID(spl[0]))
	}

	return ids
}

func Pair(fir ID, sec ID) string {
	return fmt.Sprintf("%s,%s", fir, sec)
}
