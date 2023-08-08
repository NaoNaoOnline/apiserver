package eventstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var eventDurationEmptyError = &tracer.Error{
	Kind: "eventDurationEmptyError",
	Desc: "The request expects a valid event duration for the event object. No event duration was found for the request. Therefore it failed.",
}

func IsEventDurationEmpty(err error) bool {
	return errors.Is(err, eventDurationEmptyError)
}

var eventLinkEmptyError = &tracer.Error{
	Kind: "eventLinkEmptyError",
	Desc: "The request expects a valid event link for the event object. No event link was found for the request. Therefore it failed.",
}

func IsEventLinkEmpty(err error) bool {
	return errors.Is(err, eventLinkEmptyError)
}

var eventTimeInvalidError = &tracer.Error{
	Kind: "eventTimeInvalidError",
	Desc: "The request expects a valid event time for the event object. The event time must not be empty and it must be in the future. No valid event time was found for the request. Therefore it failed.",
}

func IsEventTimeInvalid(err error) bool {
	return errors.Is(err, eventTimeInvalidError)
}

var labelNotFoundError = &tracer.Error{
	Kind: "labelNotFoundError",
	Desc: "The request expects a valid label ID for the event object. No label object was found for the request. Therefore it failed.",
}

func IsEventNotFound(err error) bool {
	return errors.Is(err, labelNotFoundError)
}

var userIDEmptyError = &tracer.Error{
	Kind: "userIDEmptyError",
	Desc: "The request expects a valid OAuth access token mapping to an internal user ID. No user ID was found for the request. Therefore it failed.",
}

func IsUserIDEmpty(err error) bool {
	return errors.Is(err, userIDEmptyError)
}
