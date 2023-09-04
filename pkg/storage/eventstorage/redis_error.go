package eventstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var eventObjectNotFoundError = &tracer.Error{
	Kind: "eventObjectNotFoundError",
	Desc: "The request expects an event object to exist. The event object was not found to exist for the request. Therefore it failed.",
}

func IsEventObjectNotFound(err error) bool {
	return errors.Is(err, eventObjectNotFoundError)
}

var labelObjectNotFoundError = &tracer.Error{
	Kind: "labelObjectNotFoundError",
	Desc: "The request expects a label object associated to the event object. The associated label object was not found for the request. Therefore it failed.",
}

func IsLabelObjectNotFound(err error) bool {
	return errors.Is(err, labelObjectNotFoundError)
}
