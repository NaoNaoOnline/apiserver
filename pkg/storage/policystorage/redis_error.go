package policystorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var policyObjectNotFoundError = &tracer.Error{
	Kind: "policyObjectNotFoundError",
	Desc: "The request expects a policy object to exist. The policy object was not found to exist. Therefore the request failed.",
}

func IsPolicyObjectNotFound(err error) bool {
	return errors.Is(err, policyObjectNotFoundError)
}
