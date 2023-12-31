package subjectclaim

import (
	"context"
)

type t string

var k t = "subjectclaim"

func NewContext(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, k, v)
}

func FromContext(ctx context.Context) string {
	v, _ := ctx.Value(k).(string)
	return v
}
