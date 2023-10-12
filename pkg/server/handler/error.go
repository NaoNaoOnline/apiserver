package handler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var ExecutionFailedError = &tracer.Error{
	Kind: "ExecutionFailedError",
	Desc: "The request expects the query object not to be empty. The query object was found to be empty. Therefore the request failed.",
}

func IsExecutionFailed(err error) bool {
	return errors.Is(err, ExecutionFailedError)
}

var QueryObjectEmptyError = &tracer.Error{
	Kind: "QueryObjectEmptyError",
	Desc: "The request expects the query object not to be empty. The query object was found to be empty. Therefore the request failed.",
}

func IsQueryObjectEmpty(err error) bool {
	return errors.Is(err, QueryObjectEmptyError)
}

var UserIDEmptyError = &tracer.Error{
	Kind: "UserIDEmptyError",
	Desc: "The request expects a valid OAuth access token mapping to an internal user ID. No user ID was found. Therefore the request failed.",
}

func IsUserIDEmpty(err error) bool {
	return errors.Is(err, UserIDEmptyError)
}

var UserNotOwnerError = &tracer.Error{
	Kind: "UserNotOwnerError",
	Desc: "The request expects the calling user to be the owner of the requested resource. The calling user was not found to be the owner of the requested resource. Therefore the request failed.",
}

func IsUserNotOwner(err error) bool {
	return errors.Is(err, UserNotOwnerError)
}
