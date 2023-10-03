package votehandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var descriptionDeletedError = &tracer.Error{
	Kind: "descriptionDeletedError",
	Desc: "The request expects vote objects to be created or deleted until the associated description was already deleted. The associated description was found to have already been deleted. Therefore the request failed.",
}

func IsDescriptionDeleted(err error) bool {
	return errors.Is(err, descriptionDeletedError)
}

var eventAlreadyHappenedError = &tracer.Error{
	Kind: "eventAlreadyHappenedError",
	Desc: "The request expects vote objects to be created or deleted until the associated event has already happened. The associated event was found to have already happened. Therefore the request failed.",
}

func IsEventAlreadyHappened(err error) bool {
	return errors.Is(err, eventAlreadyHappenedError)
}

var eventDeletedError = &tracer.Error{
	Kind: "eventDeletedError",
	Desc: "The request expects vote objects to be created or deleted until the associated event was already deleted. The associated event was found to have already been deleted. Therefore the request failed.",
}

func IsEventDeleted(err error) bool {
	return errors.Is(err, eventDeletedError)
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