package labelstorage

import (
	"errors"

	"github.com/twitchtv/twirp"
	"github.com/xh3b4sd/tracer"
)

var invalidLabelKindError = &tracer.Error{
	Kind: "invalidLabelKindError",
	Desc: "The request expects a valid label kind for the label object, e.g. cate or host. No valid label kind was found for the request. Therefore it failed.",
}

func IsInvalidLabelKind(err error) bool {
	return errors.Is(err, invalidLabelKindError)
}

var labelAlreadyExistsError = &tracer.Error{
	Kind: "labelAlreadyExistsError",
	Code: string(twirp.InvalidArgument),
	Desc: "Labels must be unique. A label with the provided name was found to exist already. Therefore the request failed.",
}

func IsLabelAlreadyExists(err error) bool {
	return errors.Is(err, labelAlreadyExistsError)
}

var labelNameEmptyError = &tracer.Error{
	Kind: "labelNameEmptyError",
	Desc: "The request expects a valid label name for the label object. No label name was found for the request. Therefore it failed.",
}

func IsLabelNameEmpty(err error) bool {
	return errors.Is(err, labelNameEmptyError)
}

var labelNotFoundError = &tracer.Error{
	Kind: "labelNotFoundError",
	Desc: "The request expects a valid label ID for the label object. No label object was found for the request. Therefore it failed.",
}

func IsLabelNotFound(err error) bool {
	return errors.Is(err, labelNotFoundError)
}

var userIDEmptyError = &tracer.Error{
	Kind: "userIDEmptyError",
	Desc: "The request expects a valid OAuth access token mapping to an internal user ID. No user ID was found for the request. Therefore it failed.",
}

func IsUserIDEmpty(err error) bool {
	return errors.Is(err, userIDEmptyError)
}
