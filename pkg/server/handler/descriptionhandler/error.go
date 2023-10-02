package descriptionhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var userIDEmptyError = &tracer.Error{
	Kind: "userIDEmptyError",
	Desc: "The request expects a valid OAuth access token mapping to an internal user ID. No user ID was found. Therefore the request failed.",
}

func IsUserIDEmpty(err error) bool {
	return errors.Is(err, userIDEmptyError)
}

var updatePeriodPastError = &tracer.Error{
	Kind: "updatePeriodPastError",
	Desc: "The request expects changes on the description object to happen within 5 minutes of resource creation. The changes on the description object were found to be after 5 minutes of resource creation. Therefore the request failed.",
}

func IsUpdatePeriodPast(err error) bool {
	return errors.Is(err, updatePeriodPastError)
}

var userNotOwnerError = &tracer.Error{
	Kind: "userNotOwnerError",
	Desc: "The request expects the calling user to be the owner of the requested resource. The calling user was not found to be the owner of the requested resource. Therefore the request failed.",
}

func IsUserNotOwner(err error) bool {
	return errors.Is(err, userNotOwnerError)
}
