package reactionhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var searchKindEmptyError = &tracer.Error{
	Kind: "searchKindEmptyError",
	Desc: "The request expects public.kind not to be empty. public.kind was found to be empty. Therefore the request failed.",
}

func IsSearchKindEmpty(err error) bool {
	return errors.Is(err, searchKindEmptyError)
}
