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

var userNotFoundError = &tracer.Error{
	Kind: "userNotFoundError",
	Desc: "The request expected an user object to be found for the user ID. No user object was found. Therefore the request failed.",
}

func IsNotFound(err error) bool {
	return errors.Is(err, userNotFoundError)
}

var subjectClaimEmptyError = &tracer.Error{
	Kind: "subjectClaimEmptyError",
	Desc: "The request expects a valid OAuth access token containing an external subject claim. No subject claim was found. Therefore the request failed.",
}

func IsSubjectClaimEmpty(err error) bool {
	return errors.Is(err, subjectClaimEmptyError)
}

var subjectClaimMappingError = &tracer.Error{
	Kind: "subjectClaimMappingError",
	Desc: "The request expects a mapping between external subject claim and internal user ID. No mapping was found. Therefore the request failed.",
}

func IsSubjectClaimMapping(err error) bool {
	return errors.Is(err, subjectClaimMappingError)
}
