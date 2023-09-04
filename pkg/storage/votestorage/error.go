package votestorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var descriptionIDEmptyError = &tracer.Error{
	Kind: "descriptionIDEmptyError",
	Desc: "The request expects a valid description ID for the associated description object. No description ID was found for the request. Therefore it failed.",
}

func IsDescriptionIDEmpty(err error) bool {
	return errors.Is(err, descriptionIDEmptyError)
}

var descriptionNotFoundError = &tracer.Error{
	Kind: "descriptionNotFoundError",
	Desc: "The request expects a valid description ID for the associated description object. No description object was found for the request. Therefore it failed.",
}

func IsDescriptionNotFound(err error) bool {
	return errors.Is(err, descriptionNotFoundError)
}

var reactionIDEmptyError = &tracer.Error{
	Kind: "reactionIDEmptyError",
	Desc: "The request expects a valid reaction ID for the associated reaction object. No reaction ID was found for the request. Therefore it failed.",
}

func IsReactionIDEmpty(err error) bool {
	return errors.Is(err, reactionIDEmptyError)
}

var reactionNotFoundError = &tracer.Error{
	Kind: "reactionNotFoundError",
	Desc: "The request expects a valid reaction ID for the associated reaction object. No reaction object was found for the request. Therefore it failed.",
}

func IsReactionNotFound(err error) bool {
	return errors.Is(err, reactionNotFoundError)
}

var voteLimitError = &tracer.Error{
	Kind: "voteLimitError",
	Desc: "The request expects an upper limit of 5 vote objects per event per user. The upper limit of 5 vote objects per event per user was found for the request. Therefore it failed.",
}

func IsVoteLimit(err error) bool {
	return errors.Is(err, voteLimitError)
}

var voteNotFoundError = &tracer.Error{
	Kind: "voteNotFoundError",
	Desc: "The request expects a valid vote ID for the associated vote object. No vote object was found for the request. Therefore it failed.",
}

func IsVoteNotFound(err error) bool {
	return errors.Is(err, voteNotFoundError)
}
