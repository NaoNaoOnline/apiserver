package reactionstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var reactionObjectAlreadyExistsError = &tracer.Error{
	Kind: "reactionObjectAlreadyExistsError",
	Desc: "The request expects a reaction object not to exist already. The reaction object was found to exist already. Therefore the request failed.",
}

func IsReactionObjectAlreadyExists(err error) bool {
	return errors.Is(err, reactionObjectAlreadyExistsError)
}

var reactionObjectNotFoundError = &tracer.Error{
	Kind: "reactionObjectNotFoundError",
	Desc: "The request expects a reaction object to exist. The reaction object was not found to exist. Therefore the request failed.",
}

func IsReactionObjectNotFound(err error) bool {
	return errors.Is(err, reactionObjectNotFoundError)
}
