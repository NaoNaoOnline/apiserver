package user

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var invalidInputError = &tracer.Error{
	Kind: "invalidInputError",
}

func IsInvalidInput(err error) bool {
	return errors.Is(err, invalidInputError)
}

var notFoundError = &tracer.Error{
	Kind: "notFoundError",
}

func IsNotFound(err error) bool {
	return errors.Is(err, notFoundError)
}
