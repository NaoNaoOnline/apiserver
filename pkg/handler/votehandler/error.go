package votehandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var userIDEmptyError = &tracer.Error{
	Kind: "userIDEmptyError",
	Desc: "The request expects a valid OAuth access token mapping to an internal user ID. No user ID was found for the request. Therefore it failed.",
}

func IsUserIDEmpty(err error) bool {
	return errors.Is(err, userIDEmptyError)
}

var userNotOwnerError = &tracer.Error{
	Kind: "userNotOwnerError",
	Desc: "The request expects the calling user to be the owner of the requested resource. The calling user was not found to be the owner of the requested resource. Therefore it failed.",
}

func IsUserNotOwner(err error) bool {
	return errors.Is(err, userNotOwnerError)
}
