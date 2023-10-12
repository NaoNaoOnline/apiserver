package wallethandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var searchKindConflictError = &tracer.Error{
	Kind: "searchKindConflictError",
	Desc: "The request expects public.kind to be the only field provided within the given search query object. Fields other than public.kind were found to be set within the given search query object. Therefore the request failed.",
}

func IsSearchKindConflict(err error) bool {
	return errors.Is(err, searchKindConflictError)
}

var searchWlltConflictError = &tracer.Error{
	Kind: "searchWlltConflictError",
	Desc: "The request expects intern.wllt to be the only field provided within the given search query object. Fields other than intern.wllt were found to be set within the given search query object. Therefore the request failed.",
}

func IsSearchWlltConflict(err error) bool {
	return errors.Is(err, searchWlltConflictError)
}
