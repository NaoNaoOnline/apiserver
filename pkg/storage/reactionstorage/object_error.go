package reactionstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var reactionHtmlEmptyError = &tracer.Error{
	Kind: "reactionHtmlEmptyError",
	Desc: "The request expects the reaction html not to be empty. The reaction html was found to be empty. Therefore the request failed.",
}

func IsReactionHtmlEmpty(err error) bool {
	return errors.Is(err, reactionHtmlEmptyError)
}

var reactionKindInvalidError = &tracer.Error{
	Kind: "reactionKindInvalidError",
	Desc: "The request expects the reaction kind to be one of [bltn user]. The reaction kind was not found to be one of [bltn user]. Therefore the request failed.",
}

func IsLabelKindInvalid(err error) bool {
	return errors.Is(err, reactionKindInvalidError)
}

var reactionNameEmptyError = &tracer.Error{
	Kind: "reactionNameEmptyError",
	Desc: "The request expects the reaction html not to be empty. The reaction html was found to be empty. Therefore the request failed.",
}

func IsReactionNameEmpty(err error) bool {
	return errors.Is(err, reactionNameEmptyError)
}
