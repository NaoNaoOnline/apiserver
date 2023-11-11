package objectid

import (
	"fmt"
	"strings"
)

func Frst(key []string) []ID {
	var ids []ID

	for _, x := range key {
		// Regardless whether the given keys are in fact paired or not, we can
		// safely take the first element from the split.
		spl := strings.Split(x, ",")
		ids = append(ids, ID(spl[0]))
	}

	return ids
}

func Scnd(key []string) []ID {
	var ids []ID

	for _, x := range key {
		// Below we can only use the second element if the given keys are in fact
		// paired.
		spl := strings.Split(x, ",")
		if len(spl) == 2 {
			ids = append(ids, ID(spl[1]))
		}
	}

	return ids
}

func Pair(fir ID, sec ID) string {
	return fmt.Sprintf("%s,%s", fir, sec)
}
