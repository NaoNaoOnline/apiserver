package wallethandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var searchInternConflictError = &tracer.Error{
	Kind: "searchInternConflictError",
	Desc: "The request expects intern.wllt to be the only field provided within the given query object. Fields other than intern.wllt were found to be set within the given query object. Therefore the request failed.",
}

func IsSearchInternConflict(err error) bool {
	return errors.Is(err, searchInternConflictError)
}

var searchInternEmptyError = &tracer.Error{
	Kind: "searchInternEmptyError",
	Desc: "The request expects intern.wllt not to be empty. intern.wllt was found to be empty. Therefore the request failed.",
}

func IsSearchInternEmpty(err error) bool {
	return errors.Is(err, searchInternEmptyError)
}

var searchPublicConflictError = &tracer.Error{
	Kind: "searchPublicConflictError",
	Desc: "The request expects public.kind to be the only field provided within the given query object. Fields other than public.kind were found to be set within the given query object. Therefore the request failed.",
}

func IsSearchPublicConflict(err error) bool {
	return errors.Is(err, searchPublicConflictError)
}

var searchPublicEmptyError = &tracer.Error{
	Kind: "searchPublicEmptyError",
	Desc: "The request expects public.kind not to be empty. public.kind was found to be empty. Therefore the request failed.",
}

func IsSearchPublicEmpty(err error) bool {
	return errors.Is(err, searchPublicEmptyError)
}

var updateInternEmptyError = &tracer.Error{
	Kind: "updateInternEmptyError",
	Desc: "The request expects intern.wllt not to be empty. intern.wllt was found to be empty. Therefore the request failed.",
}

func IsUpdateInternEmpty(err error) bool {
	return errors.Is(err, updateInternEmptyError)
}
