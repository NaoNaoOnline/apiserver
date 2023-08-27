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

var eventLinkInvalidError = &tracer.Error{
	Kind: "eventLinkInvalidError",
	Desc: "The request expects a valid event link for the event object. The event link must not be empty and it must be a valid URL. No valid event link was found for the request. Therefore it failed.",
}

func IsEventLinkInvalid(err error) bool {
	return errors.Is(err, eventLinkInvalidError)
}

var eventNotFoundError = &tracer.Error{
	Kind: "eventNotFoundError",
	Desc: "The request expected an event object to be found for the given event ID. No event object was found for the request. Therefore it failed.",
}

func IsEventNotFound(err error) bool {
	return errors.Is(err, eventNotFoundError)
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

func IsLabelNotFound(err error) bool {
	return errors.Is(err, labelNotFoundError)
}

var tooManyLabelsError = &tracer.Error{
	Kind: "tooManyLabelsError",
	Desc: "The request expects a maximum of 5 category and host labels for the event object. Too many labels were found for the request. Therefore it failed.",
}

func IsTooManyLabels(err error) bool {
	return errors.Is(err, tooManyLabelsError)
}

var userIDEmptyError = &tracer.Error{
	Kind: "userIDEmptyError",
	Desc: "The request expects a valid OAuth access token mapping to an internal user ID. No user ID was found for the request. Therefore it failed.",
}

func IsUserIDEmpty(err error) bool {
	return errors.Is(err, userIDEmptyError)
}
