package liststorage

import (
	"github.com/xh3b4sd/tracer"
)

var listObjectNotFoundError = &tracer.Error{
	Kind: "listObjectNotFoundError",
	Desc: "The request expects a list object to exist. The list object was not found to exist. Therefore the request failed.",
}
