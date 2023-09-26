package votestorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var descriptionObjectNotFoundError = &tracer.Error{
	Kind: "descriptionObjectNotFoundError",
	Desc: "The request expects a description object to exist. The description object was not found to exist. Therefore the request failed.",
}

func IsDescriptionObjectNotFound(err error) bool {
	return errors.Is(err, descriptionObjectNotFoundError)
}

var eventAlreadyHappenedError = &tracer.Error{
	Kind: "eventAlreadyHappenedError",
	Desc: "The request expects vote objects to be created or deleted until the associated event has already happened. The associated event was found to have already happened. Therefore the request failed.",
}

func IsEventAlreadyHappened(err error) bool {
	return errors.Is(err, eventAlreadyHappenedError)
}

var eventObjectNotFoundError = &tracer.Error{
	Kind: "eventObjectNotFoundError",
	Desc: "The request expects an event object to exist. The event object was not found to exist. Therefore the request failed.",
}

func IsEventObjectNotFound(err error) bool {
	return errors.Is(err, eventObjectNotFoundError)
}

var reactionObjectNotFoundError = &tracer.Error{
	Kind: "reactionObjectNotFoundError",
	Desc: "The request expects a reaction object to exist. The reaction object was not found to exist. Therefore the request failed.",
}

func IsReactionObjectNotFound(err error) bool {
	return errors.Is(err, reactionObjectNotFoundError)
}

var voteEventLimitError = &tracer.Error{
	Kind: "voteEventLimitError",
	Desc: "The request expects an upper limit of 5 vote objects per event per user. The upper limit of 5 vote objects per event per user was found. Therefore the request failed.",
}

func IsVoteEventLimit(err error) bool {
	return errors.Is(err, voteEventLimitError)
}

var voteObjectNotFoundError = &tracer.Error{
	Kind: "voteObjectNotFoundError",
	Desc: "The request expects a vote object to exist. The vote object was not found to exist. Therefore the request failed.",
}

func IsVoteObjectNotFound(err error) bool {
	return errors.Is(err, voteObjectNotFoundError)
}

var voteUserLimitError = &tracer.Error{
	Kind: "voteUserLimitError",
	Desc: "The request expects an upper limit of 100 vote objects per user globally. The upper limit of 100 vote objects per user globally was found. Therefore the request failed.",
}

func IsVoteUserLimit(err error) bool {
	return errors.Is(err, voteUserLimitError)
}
