package userhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var searchInternConflictError = &tracer.Error{
	Kind: "searchInternConflictError",
	Desc: "The request expects intern.user to be the only field provided within the given query object. Fields other than intern.user were found to be set within the given query object. Therefore the request failed.",
}

func IsSearchInternConflict(err error) bool {
	return errors.Is(err, searchInternConflictError)
}

var searchInternEmptyError = &tracer.Error{
	Kind: "searchInternEmptyError",
	Desc: "The request expects intern.user not to be empty. intern.user was found to be empty. Therefore the request failed.",
}

func IsSearchInternEmpty(err error) bool {
	return errors.Is(err, searchInternEmptyError)
}

var searchPublicConflictError = &tracer.Error{
	Kind: "searchPublicConflictError",
	Desc: "The request expects public.name to be the only field provided within the given query object. Fields other than public.name were found to be set within the given query object. Therefore the request failed.",
}

func IsSearchPublicConflict(err error) bool {
	return errors.Is(err, searchPublicConflictError)
}

var searchPublicEmptyError = &tracer.Error{
	Kind: "searchPublicEmptyError",
	Desc: "The request expects public.name not to be empty. public.name was found to be empty. Therefore the request failed.",
}

func IsSearchPublicEmpty(err error) bool {
	return errors.Is(err, searchPublicEmptyError)
}

var searchSymbolConflictError = &tracer.Error{
	Kind: "searchSymbolConflictError",
	Desc: "The request expects symbol.user to be the only field provided within the given query object. Fields other than symbol.user were found to be set within the given query object. Therefore the request failed.",
}

func IsSearchSymbolConflict(err error) bool {
	return errors.Is(err, searchSymbolConflictError)
}

var searchSymbolEmptyError = &tracer.Error{
	Kind: "searchSymbolEmptyError",
	Desc: "The request expects symbol.user not to be empty. symbol.user was found to be empty. Therefore the request failed.",
}

func IsSearchSymbolEmpty(err error) bool {
	return errors.Is(err, searchSymbolEmptyError)
}

var searchSymbolInvalidError = &tracer.Error{
	Kind: "searchSymbolInvalidError",
	Desc: `The request expects symbol.user to be set to "self". symbol.user was not found to be set to "self". Therefore the request failed.`,
}

func IsSearchSymbolInvalid(err error) bool {
	return errors.Is(err, searchSymbolInvalidError)
}
