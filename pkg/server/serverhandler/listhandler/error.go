package listhandler

import (
	"github.com/xh3b4sd/tracer"
)

var createDescEmptyError = &tracer.Error{
	Kind: "createDescEmptyError",
	Desc: "The request expects the list description not to be empty. The list description was found to be empty. Therefore the request failed.",
}

var updateEmptyError = &tracer.Error{
	Kind: "updateEmptyError",
	Desc: "The request expects the query object to contain all of [intern update]. The query object was not found to contain all of [intern update]. Therefore the request failed.",
}
