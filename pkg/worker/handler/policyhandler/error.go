package policyhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var policyLatestInvalidError = &tracer.Error{
	Kind: "policyLatestInvalidError",
	Desc: "The task expects policy.naonao.io/latest to contain 4 comma separated block numbers. policy.naonao.io/latest was not found to contain 4 comma separated block numbers. Therefore the task failed.",
}

func IsPolicyLatestInvalid(err error) bool {
	return errors.Is(err, policyLatestInvalidError)
}
