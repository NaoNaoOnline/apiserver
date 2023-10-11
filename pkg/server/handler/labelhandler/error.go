package labelhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var createKindInvalidError = &tracer.Error{
	Kind: "createKindInvalidError",
	Desc: "The request expects the label kind to be one of [cate host]. The label kind was not found to be one of [cate host]. Therefore the request failed.",
}

func IsCreateKindInvalid(err error) bool {
	return errors.Is(err, createKindInvalidError)
}

var userIDEmptyError = &tracer.Error{
	Kind: "userIDEmptyError",
	Desc: "The request expects a valid OAuth access token mapping to an internal user ID. No user ID was found. Therefore the request failed.",
}

func IsUserIDEmpty(err error) bool {
	return errors.Is(err, userIDEmptyError)
}
