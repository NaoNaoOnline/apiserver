package objectid

import "time"

func Time(tim ...time.Time) int64 {
	if len(tim) == 1 && !tim[0].IsZero() {
		return tim[0].Unix()
	}

	return time.Now().Unix()
}
