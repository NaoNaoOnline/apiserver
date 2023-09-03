package descriptionstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var descriptionObjectNotFoundError = &tracer.Error{
	Kind: "descriptionObjectNotFoundError",
	Desc: "The request expects a description object to exist. The description object was not found to exist for the request. Therefore it failed.",
}

func IsDescriptionObjectNotFound(err error) bool {
	return errors.Is(err, descriptionObjectNotFoundError)
}

var eventObjectNotFoundError = &tracer.Error{
	Kind: "eventObjectNotFoundError",
	Desc: "The request expects an event object associated to the description object. The associated event object was not found for the request. Therefore it failed.",
}

func IsEventObjectNotFound(err error) bool {
	return errors.Is(err, eventObjectNotFoundError)
}
