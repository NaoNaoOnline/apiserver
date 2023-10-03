package eventhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var searchEvntConflictError = &tracer.Error{
	Kind: "searchEvntConflictError",
	Desc: "The request expects intern.evnt to be the only field provided within the given search query object. Fields other than intern.evnt were found to be set within the given search query object. Therefore the request failed.",
}

func IsSearchEvntConflict(err error) bool {
	return errors.Is(err, searchEvntConflictError)
}

var searchUserConflictError = &tracer.Error{
	Kind: "searchUserConflictError",
	Desc: "The request expects intern.user to be the only field provided within the given search query object. Fields other than intern.user were found to be set within the given search query object. Therefore the request failed.",
}

func IsSearchUserConflict(err error) bool {
	return errors.Is(err, searchUserConflictError)
}

var userIDEmptyError = &tracer.Error{
	Kind: "userIDEmptyError",
	Desc: "The request expects a valid OAuth access token mapping to an internal user ID. No user ID was found. Therefore the request failed.",
}

func IsUserIDEmpty(err error) bool {
	return errors.Is(err, userIDEmptyError)
}

var userNotOwnerError = &tracer.Error{
	Kind: "userNotOwnerError",
	Desc: "The request expects the calling user to be the owner of the requested resource. The calling user was not found to be the owner of the requested resource. Therefore the request failed.",
}

func IsUserNotOwner(err error) bool {
	return errors.Is(err, userNotOwnerError)
}
