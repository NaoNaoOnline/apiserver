package user

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var subjectClaimEmptyError = &tracer.Error{
	Kind: "subjectClaimEmptyError",
	Desc: "The request expects a valid OAuth access token containing an external subject claim. No subject claim was found for the request. Therefore it failed.",
}

func IsSubjectClaimEmpty(err error) bool {
	return errors.Is(err, subjectClaimEmptyError)
}
