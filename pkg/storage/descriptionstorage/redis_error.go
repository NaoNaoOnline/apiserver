package descriptionstorage

import (
	"github.com/xh3b4sd/tracer"
)

var descriptionEventLimitError = &tracer.Error{
	Kind: "descriptionEventLimitError",
	Desc: "The request expects an upper limit of 50 description objects per event. The upper limit of 50 description objects per event was found. Therefore the request failed.",
}

var descriptionLikeAlreadyExistsError = &tracer.Error{
	Kind: "descriptionLikeAlreadyExistsError",
	Desc: "The request expects the calling user to only like a description that was not liked before. The description was found to have already been liked by the calling user. Therefore the request failed.",
}

var descriptionObjectNotFoundError = &tracer.Error{
	Kind: "descriptionObjectNotFoundError",
	Desc: "The request expects a description object to exist. The description object was not found to exist. Therefore the request failed.",
}

var descriptionUnlikeAlreadyExistsError = &tracer.Error{
	Kind: "descriptionUnlikeAlreadyExistsError",
	Desc: "The request expects the calling user to only unlike a description that was liked before. The description was found to have already been unliked by the calling user. Therefore the request failed.",
}

var eventAlreadyHappenedError = &tracer.Error{
	Kind: "eventAlreadyHappenedError",
	Desc: "The request expects description objects to be created until the associated event has already happened. The associated event was found to have already happened. Therefore the request failed.",
}

var eventObjectNotFoundError = &tracer.Error{
	Kind: "eventObjectNotFoundError",
	Desc: "The request expects an event object associated to the description object. The associated event object was not found. Therefore the request failed.",
}
