package descriptionhandler

import (
	"github.com/xh3b4sd/tracer"
)

var descriptionDeletePeriodError = &tracer.Error{
	Kind: "descriptionDeletePeriodError",
	Desc: "The request expects the description object to be deleted within 5 minutes of resource creation. The description object was tried to be deleted after 5 minutes of resource creation. Therefore the request failed.",
}

var descriptionUpdatePeriodError = &tracer.Error{
	Kind: "descriptionUpdatePeriodError",
	Desc: "The request expects the description object to be updated within 5 minutes of resource creation. The description object was tried to be updated after 5 minutes of resource creation. Therefore the request failed.",
}

var descriptionRequirementError = &tracer.Error{
	Kind: "descriptionRequirementError",
	Desc: "The request expects the only description object of an event not to be deleted. The only description object of an event was tried to be deleted. Therefore the request failed.",
}

var eventAlreadyHappenedError = &tracer.Error{
	Kind: "eventAlreadyHappenedError",
	Desc: "The request expects descriptions to be added, removed or modified until the associated event has already happened. The associated event was found to have already happened. Therefore the request failed.",
}

var eventDeletedError = &tracer.Error{
	Kind: "eventDeletedError",
	Desc: "The request expects descriptions to be added, removed or modified until the associated event has been deleted. The associated event was found to have already been deleted. Therefore the request failed.",
}

var updateEmptyError = &tracer.Error{
	Kind: "updateEmptyError",
	Desc: "The request expects the query object to contain all of [intern update]. The query object was not found to contain all of [intern update]. Therefore the request failed.",
}

var updateSymbolConflictError = &tracer.Error{
	Kind: "updateSymbolConflictError",
	Desc: "The request expects the query object to contain one of [symbol update]. The query object was not found to contain one of [symbol update]. Therefore the request failed.",
}

var updateSymbolInvalidError = &tracer.Error{
	Kind: "updateSymbolInvalidError",
	Desc: "The request expects symbol.like to be one of [add rem]. symbol.like was not found to be one of [add rem]. Therefore the request failed.",
}
