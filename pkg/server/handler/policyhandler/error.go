package policyhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var searchKindConflictError = &tracer.Error{
	Kind: "searchKindConflictError",
	Desc: "The request expects public.kind not to be provided if symbol.ltst is one of [default aggregated]. public.kind was found to be provided when symbol.ltst was one of [default aggregated]. Therefore the request failed.",
}

func IsSearchKindConflict(err error) bool {
	return errors.Is(err, searchKindConflictError)
}
