package rulestorage

import (
	"github.com/xh3b4sd/tracer"
)

var resourceObjectNotFoundError = &tracer.Error{
	Kind: "resourceObjectNotFoundError",
	Desc: "The request expects all resource objects referenced in the rule object to exist. Not all of the resource objects referenced in the rule object were found to exist. Therefore the request failed.",
}
