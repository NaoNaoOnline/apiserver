package userid

import (
	"context"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
)

type t string

var k t = "userid"

func NewContext(ctx context.Context, v scoreid.String) context.Context {
	return context.WithValue(ctx, k, v)
}

func FromContext(ctx context.Context) scoreid.String {
	v, _ := ctx.Value(k).(scoreid.String)
	return v
}
