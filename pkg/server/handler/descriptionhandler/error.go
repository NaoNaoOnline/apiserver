package descriptionhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var descriptionDeletePeriodError = &tracer.Error{
	Kind: "descriptionDeletePeriodError",
	Desc: "The request expects the description object to be deleted within 5 minutes of resource creation. The description object was tried to be deleted after 5 minutes of resource creation. Therefore the request failed.",
}

func IsDeletePeriodPast(err error) bool {
	return errors.Is(err, descriptionDeletePeriodError)
}

var descriptionUpdatePeriodError = &tracer.Error{
	Kind: "descriptionUpdatePeriodError",
	Desc: "The request expects the description object to be updated within 5 minutes of resource creation. The description object was tried to be updated after 5 minutes of resource creation. Therefore the request failed.",
}

func IsUpdatePeriodPast(err error) bool {
	return errors.Is(err, descriptionUpdatePeriodError)
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
