package eventhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var searchEvntConflictError = &tracer.Error{
	Kind: "searchEvntConflictError",
	Desc: "The request expects intern.evnt to be the only field provided within the given query object. Fields other than intern.evnt were found to be set within the given query object. Therefore the request failed.",
}

func IsSearchEvntConflict(err error) bool {
	return errors.Is(err, searchEvntConflictError)
}

var searchUserConflictError = &tracer.Error{
	Kind: "searchUserConflictError",
	Desc: "The request expects intern.user to be the only field provided within the given query object. Fields other than intern.user were found to be set within the given query object. Therefore the request failed.",
}

func IsSearchUserConflict(err error) bool {
	return errors.Is(err, searchUserConflictError)
}
