package descriptionstorage

import (
	"github.com/xh3b4sd/tracer"
)

var descriptionObjectNotFoundError = &tracer.Error{
	Kind: "descriptionObjectNotFoundError",
	Desc: "The request expects a description object to exist. The description object was not found to exist. Therefore the request failed.",
}

var eventAlreadyHappenedError = &tracer.Error{
	Kind: "eventAlreadyHappenedError",
	Desc: "The request expects description objects to be created until the associated event has already happened. The associated event was found to have already happened. Therefore the request failed.",
}

var eventObjectNotFoundError = &tracer.Error{
	Kind: "eventObjectNotFoundError",
	Desc: "The request expects an event object associated to the description object. The associated event object was not found. Therefore the request failed.",
}
