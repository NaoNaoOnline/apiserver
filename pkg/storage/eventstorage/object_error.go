package eventstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var eventDurationEmptyError = &tracer.Error{
	Kind: "eventDurationEmptyError",
	Desc: "The request expects the event duration not to be empty. The event duration was found to be empty for the request. Therefore it failed.",
}

func IsEventDurationEmpty(err error) bool {
	return errors.Is(err, eventDurationEmptyError)
}

var eventDurationLimitError = &tracer.Error{
	Kind: "eventDurationLimitError",
	Desc: "The request expects the event duration not to be over 4 hours. The event duration was found to be over 4 hours for the request. Therefore it failed.",
}

func IsEventDurationLimit(err error) bool {
	return errors.Is(err, eventDurationLimitError)
}

var eventDurationNegativeError = &tracer.Error{
	Kind: "eventDurationNegativeError",
	Desc: "The request expects the event duration not to be negative. The event duration was found to be negative for the request. Therefore it failed.",
}

func IsEventDurationNegative(err error) bool {
	return errors.Is(err, eventDurationNegativeError)
}

var eventLinkEmptyError = &tracer.Error{
	Kind: "eventLinkEmptyError",
	Desc: "The request expects the event link not to be empty. The event link was found to be empty for the request. Therefore it failed.",
}

func IsEventLinkEmpty(err error) bool {
	return errors.Is(err, eventLinkEmptyError)
}

var eventLinkFormatError = &tracer.Error{
	Kind: "eventLinkFormatError",
	Desc: "The request expects the event link to be a valid https URL. The event link was not found to be a valid https URL for the request. Therefore it failed.",
}

func IsLabelLinkFormat(err error) bool {
	return errors.Is(err, eventLinkFormatError)
}

var eventTimeEmptyError = &tracer.Error{
	Kind: "eventTimeEmptyError",
	Desc: "The request expects the event time not to be empty. The event time was found to be empty for the request. Therefore it failed.",
}

func IsEventTimeEmpty(err error) bool {
	return errors.Is(err, eventTimeEmptyError)
}

var eventTimePastError = &tracer.Error{
	Kind: "eventTimePastError",
	Desc: "The request expects the event time not to be in the past. The event time was found to be in the past for the request. Therefore it failed.",
}

func IsEventTimePast(err error) bool {
	return errors.Is(err, eventTimePastError)
}

var eventLabelEmptyError = &tracer.Error{
	Kind: "eventLabelEmptyError",
	Desc: "The request expects the event labels not to be empty. The event labels were found to be empty for the request. Therefore it failed.",
}

func IsEventLabelEmpty(err error) bool {
	return errors.Is(err, eventLabelEmptyError)
}

var eventLabelLimitError = &tracer.Error{
	Kind: "eventLabelLimitError",
	Desc: "The request expects an upper limit of 5 label IDs per event. The upper limit of 5 label IDs per event was found for the request. Therefore it failed.",
}

func IsLabelLimitError(err error) bool {
	return errors.Is(err, eventLabelLimitError)
}
