package liststorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var listObjectNotFoundError = &tracer.Error{
	Kind: "listObjectNotFoundError",
	Desc: "The request expects a list object to exist. The list object was not found to exist. Therefore the request failed.",
}

var listUserLimitError = &tracer.Error{
	Kind: "listUserLimitError",
	Desc: "The request expects an upper limit of 50 list objects per user. The upper limit of 50 list objects per user was found. Therefore the request failed.",
}

func IsListUserLimit(err error) bool {
	return errors.Is(err, listUserLimitError)
}
