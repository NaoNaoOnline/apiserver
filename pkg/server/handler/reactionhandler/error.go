package reactionhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var queryObjectEmptyError = &tracer.Error{
	Kind: "queryObjectEmptyError",
	Desc: "The request expects the query object not to be empty. The query object was found to be empty. Therefore the request failed.",
}

func IsQueryObjectEmpty(err error) bool {
	return errors.Is(err, queryObjectEmptyError)
}

var searchKindEmptyError = &tracer.Error{
	Kind: "searchKindEmptyError",
	Desc: "The request expects public.kind not to be empty. public.kind was found to be empty. Therefore the request failed.",
}

func IsSearchKindEmpty(err error) bool {
	return errors.Is(err, searchKindEmptyError)
}
