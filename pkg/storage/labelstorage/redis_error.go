package labelstorage

import (
	"errors"

	"github.com/twitchtv/twirp"
	"github.com/xh3b4sd/tracer"
)

var labelAlreadyExistsError = &tracer.Error{
	Kind: "labelAlreadyExistsError",
	Code: string(twirp.InvalidArgument),
	Desc: "Labels must be unique. A label with the provided name was found to exist already. Therefore the request failed.",
}

func IsLabelAlreadyExists(err error) bool {
	return errors.Is(err, labelAlreadyExistsError)
}

var labelNotFoundError = &tracer.Error{
	Kind: "labelNotFoundError",
	Desc: "The request expects a valid label ID for the label object. No label object was found for the request. Therefore it failed.",
}

func IsLabelNotFound(err error) bool {
	return errors.Is(err, labelNotFoundError)
}
