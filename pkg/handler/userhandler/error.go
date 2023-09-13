package userhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var searchNameConflictError = &tracer.Error{
	Kind: "searchNameConflictError",
	Desc: "The request expects public.name to be the only field provided within the given search query object. Fields other than public.name were found to be set within the given search query object. Therefore the request failed.",
}

func IsSearchNameConflict(err error) bool {
	return errors.Is(err, searchNameConflictError)
}

var searchUserConflictError = &tracer.Error{
	Kind: "searchUserConflictError",
	Desc: "The request expects intern.user to be the only field provided within the given search query object. Fields other than intern.user were found to be set within the given search query object. Therefore the request failed.",
}

func IsSearchUserConflict(err error) bool {
	return errors.Is(err, searchUserConflictError)
}
