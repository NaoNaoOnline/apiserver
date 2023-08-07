package userstorage

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

var subjectClaimEmptyError = &tracer.Error{
	Kind: "subjectClaimEmptyError",
	Desc: "The request expects a valid OAuth access token containing an external subject claim. No subject claim was found for the request. Therefore it failed.",
}

func IsSubjectClaimEmpty(err error) bool {
	return errors.Is(err, subjectClaimEmptyError)
}
