package labelstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var alreadyExistsError = &tracer.Error{
	Kind: "alreadyExistsError",
	Desc: "Labels must be unique. A label with the provided name was found to exist already. Therefore the request failed.",
}

func IsAlreadyExists(err error) bool {
	return errors.Is(err, alreadyExistsError)
}

var invalidInputError = &tracer.Error{
	Kind: "invalidInputError",
}

func IsInvalidInput(err error) bool {
	return errors.Is(err, invalidInputError)
}
