package votestorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var descriptionIDEmptyError = &tracer.Error{
	Kind: "descriptionIDEmptyError",
	Desc: "The request expects the description ID not to be empty. The description ID was found to be empty for the request. Therefore it failed.",
}

func IsDescriptionIDEmpty(err error) bool {
	return errors.Is(err, descriptionIDEmptyError)
}

var reactionIDEmptyError = &tracer.Error{
	Kind: "reactionIDEmptyError",
	Desc: "The request expects the reaction ID not to be empty. The reaction ID was found to be empty for the request. Therefore it failed.",
}

func IsReactionIDEmpty(err error) bool {
	return errors.Is(err, reactionIDEmptyError)
}
