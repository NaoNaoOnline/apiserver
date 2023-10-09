package policystorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var policyObjectAlreadyExistsError = &tracer.Error{
	Kind: "policyObjectAlreadyExistsError",
	Desc: "The request expects a policy object not to exist already. The policy object was found to exist already. Therefore the request failed.",
}

func IsLabelObjectAlreadyExists(err error) bool {
	return errors.Is(err, policyObjectAlreadyExistsError)
}

var policyObjectNotFoundError = &tracer.Error{
	Kind: "policyObjectNotFoundError",
	Desc: "The request expects a policy object to exist. The policy object was not found to exist. Therefore the request failed.",
}

func IsLabelObjectNotFound(err error) bool {
	return errors.Is(err, policyObjectNotFoundError)
}
