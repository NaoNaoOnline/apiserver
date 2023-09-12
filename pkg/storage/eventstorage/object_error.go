package eventstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var eventDurationEmptyError = &tracer.Error{
	Kind: "eventDurationEmptyError",
	Desc: "The request expects the event duration not to be empty. The event duration was found to be empty. Therefore the request failed.",
}

func IsEventDurationEmpty(err error) bool {
	return errors.Is(err, eventDurationEmptyError)
}

var eventDurationLimitError = &tracer.Error{
	Kind: "eventDurationLimitError",
	Desc: "The request expects the event duration not to be over 4 hours. The event duration was found to be over 4 hours. Therefore the request failed.",
}

func IsEventDurationLimit(err error) bool {
	return errors.Is(err, eventDurationLimitError)
}

var eventDurationNegativeError = &tracer.Error{
	Kind: "eventDurationNegativeError",
	Desc: "The request expects the event duration not to be negative. The event duration was found to be negative. Therefore the request failed.",
}

func IsEventDurationNegative(err error) bool {
	return errors.Is(err, eventDurationNegativeError)
}

var eventLabelDuplicateError = &tracer.Error{
	Kind: "eventLabelDuplicateError",
	Desc: "The request expects the event labels not to be duplicated. The event labels were found to be duplicated. Therefore the request failed.",
}

func IsEventLabelDuplicate(err error) bool {
	return errors.Is(err, eventLabelDuplicateError)
}

var eventLabelEmptyError = &tracer.Error{
	Kind: "eventLabelEmptyError",
	Desc: "The request expects the event labels not to be empty. The event labels were found to be empty. Therefore the request failed.",
}

func IsEventLabelEmpty(err error) bool {
	return errors.Is(err, eventLabelEmptyError)
}

var eventLabelLimitError = &tracer.Error{
	Kind: "eventLabelLimitError",
	Desc: "The request expects an upper limit of 5 label IDs per event. The upper limit of 5 label IDs per event was found. Therefore the request failed.",
}

func IsEventLabelLimit(err error) bool {
	return errors.Is(err, eventLabelLimitError)
}

var eventLinkEmptyError = &tracer.Error{
	Kind: "eventLinkEmptyError",
	Desc: "The request expects the event link not to be empty. The event link was found to be empty. Therefore the request failed.",
}

func IsEventLinkEmpty(err error) bool {
	return errors.Is(err, eventLinkEmptyError)
}

var eventLinkFormatError = &tracer.Error{
	Kind: "eventLinkFormatError",
	Desc: "The request expects the event link to be a valid https URL. The event link was not found to be a valid https URL. Therefore the request failed.",
}

func IsEventLinkFormat(err error) bool {
	return errors.Is(err, eventLinkFormatError)
}

var eventTimeEmptyError = &tracer.Error{
	Kind: "eventTimeEmptyError",
	Desc: "The request expects the event time not to be empty. The event time was found to be empty. Therefore the request failed.",
}

func IsEventTimeEmpty(err error) bool {
	return errors.Is(err, eventTimeEmptyError)
}

var eventTimeFutureError = &tracer.Error{
	Kind: "eventTimeFutureError",
	Desc: "The request expects the event time to be within the next 30 days. The event time was not found to be within the next 30 days. Therefore the request failed.",
}

func IsEventTimeFuture(err error) bool {
	return errors.Is(err, eventTimeFutureError)
}

var eventTimePastError = &tracer.Error{
	Kind: "eventTimePastError",
	Desc: "The request expects the event time not to be in the past. The event time was found to be in the past. Therefore the request failed.",
}

func IsEventTimePast(err error) bool {
	return errors.Is(err, eventTimePastError)
}
