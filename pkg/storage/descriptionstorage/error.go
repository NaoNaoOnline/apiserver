package descriptionstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var descriptionTextEmptyError = &tracer.Error{
	Kind: "descriptionTextEmptyError",
	Desc: "The request expects a valid description text for the description object. No description text was found for the request. Therefore it failed.",
}

func IsDescriptionTextEmpty(err error) bool {
	return errors.Is(err, descriptionTextEmptyError)
}

var eventNotFoundError = &tracer.Error{
	Kind: "eventNotFoundError",
	Desc: "The request expects a valid event ID for the description object. No event object was found for the request. Therefore it failed.",
}

func IsEventNotFound(err error) bool {
	return errors.Is(err, eventNotFoundError)
}

var invalidEventIDError = &tracer.Error{
	Kind: "invalidEventIDError",
	Desc: "The request expects a valid event ID for the description object. No event ID was found for the request. Therefore it failed.",
}

func IsInvalidEventID(err error) bool {
	return errors.Is(err, invalidEventIDError)
}

var userIDEmptyError = &tracer.Error{
	Kind: "userIDEmptyError",
	Desc: "The request expects a valid OAuth access token mapping to an internal user ID. No user ID was found for the request. Therefore it failed.",
}

func IsUserIDEmpty(err error) bool {
	return errors.Is(err, userIDEmptyError)
}
