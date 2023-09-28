package userid

import (
	"context"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

type t string

var k t = "userid"

func NewContext(ctx context.Context, v objectid.String) context.Context {
	return context.WithValue(ctx, k, v)
}

func FromContext(ctx context.Context) objectid.String {
	v, _ := ctx.Value(k).(objectid.String)
	return v
}
