package eventstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var eventObjectNotFoundError = &tracer.Error{
	Kind: "eventObjectNotFoundError",
	Desc: "The request expects an event object to exist. The event object was not found to exist. Therefore the request failed.",
}

func IsEventObjectNotFound(err error) bool {
	return errors.Is(err, eventObjectNotFoundError)
}

var hostParticipationConflictError = &tracer.Error{
	Kind: "hostParticipationConflictError",
	Desc: "The request expects the host not to participate on multiple events at the same time. The host was found to particpate on multiple events at the same time. Therefore the request failed.",
}

func IsEventParticipationConflict(err error) bool {
	return errors.Is(err, hostParticipationConflictError)
}

var labelObjectNotFoundError = &tracer.Error{
	Kind: "labelObjectNotFoundError",
	Desc: "The request expects a label object associated to the event object. The associated label object was not found. Therefore the request failed.",
}

func IsLabelObjectNotFound(err error) bool {
	return errors.Is(err, labelObjectNotFoundError)
}
