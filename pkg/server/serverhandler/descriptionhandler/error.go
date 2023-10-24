package descriptionhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var descriptionDeletePeriodError = &tracer.Error{
	Kind: "descriptionDeletePeriodError",
	Desc: "The request expects the description object to be deleted within 5 minutes of resource creation. The description object was tried to be deleted after 5 minutes of resource creation. Therefore the request failed.",
}

func IsDeletePeriodPast(err error) bool {
	return errors.Is(err, descriptionDeletePeriodError)
}

var descriptionUpdatePeriodError = &tracer.Error{
	Kind: "descriptionUpdatePeriodError",
	Desc: "The request expects the description object to be updated within 5 minutes of resource creation. The description object was tried to be updated after 5 minutes of resource creation. Therefore the request failed.",
}

func IsUpdatePeriodPast(err error) bool {
	return errors.Is(err, descriptionUpdatePeriodError)
}

var descriptionRequirementError = &tracer.Error{
	Kind: "descriptionRequirementError",
	Desc: "The request expects the only description object of an event not to be deleted. The only description object of an event was tried to be deleted. Therefore the request failed.",
}

func IsDescriptionRequirement(err error) bool {
	return errors.Is(err, descriptionRequirementError)
}

var eventAlreadyHappenedError = &tracer.Error{
	Kind: "eventAlreadyHappenedError",
	Desc: "The request expects vote objects to be created or deleted until the associated event has already happened. The associated event was found to have already happened. Therefore the request failed.",
}

func IsEventAlreadyHappened(err error) bool {
	return errors.Is(err, eventAlreadyHappenedError)
}

var eventDeletedError = &tracer.Error{
	Kind: "eventDeletedError",
	Desc: "The request expects vote objects to be created or deleted until the associated event was already deleted. The associated event was found to have already been deleted. Therefore the request failed.",
}

func IsEventDeleted(err error) bool {
	return errors.Is(err, eventDeletedError)
}

var updateEmptyError = &tracer.Error{
	Kind: "updateEmptyError",
	Desc: "The request expects the query object to contain all of [intern update]. The query object was not found to contain one of [intern.evnt intern.user]. Therefore the request failed.",
}

func IsUpdateEmpty(err error) bool {
	return errors.Is(err, updateEmptyError)
}
