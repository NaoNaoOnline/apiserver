package isprem

import (
	"context"
	"time"
)

type t string

var k t = "isprem"

func NewContext(ctx context.Context, v time.Time) context.Context {
	return context.WithValue(ctx, k, v)
}

func FromContext(ctx context.Context, tim ...time.Time) bool {
	var now time.Time
	{
		if len(tim) == 1 && !tim[0].IsZero() {
			now = tim[0].UTC()
		} else {
			now = time.Now().UTC()
		}
	}

	var val time.Time
	{
		val, _ = ctx.Value(k).(time.Time)
	}

	return !val.IsZero() && now.Before(val)
}
