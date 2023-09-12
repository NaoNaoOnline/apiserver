package labelstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var labelObjectAlreadyExistsError = &tracer.Error{
	Kind: "labelObjectAlreadyExistsError",
	Desc: "The request expects a label object not to exist already. The label object was found to exist already. Therefore the request failed.",
}

func IsLabelObjectAlreadyExists(err error) bool {
	return errors.Is(err, labelObjectAlreadyExistsError)
}

var labelObjectNotFoundError = &tracer.Error{
	Kind: "labelObjectNotFoundError",
	Desc: "The request expects an label object to exist. The label object was not found to exist. Therefore the request failed.",
}

func IsEventObjectNotFound(err error) bool {
	return errors.Is(err, labelObjectNotFoundError)
}
