package userhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var queryObjectEmptyError = &tracer.Error{
	Kind: "queryObjectEmptyError",
	Desc: "The request expects the query object not to be empty. The query object was found to be empty. Therefore the request failed.",
}

func IsQueryObjectEmpty(err error) bool {
	return errors.Is(err, queryObjectEmptyError)
}

var searchInternEmptyError = &tracer.Error{
	Kind: "searchInternEmptyError",
	Desc: "The request expects intern.user not to be empty. intern.user was found to be empty. Therefore the request failed.",
}

func IsSearchInternEmpty(err error) bool {
	return errors.Is(err, searchInternEmptyError)
}

var searchPublicEmptyError = &tracer.Error{
	Kind: "searchPublicEmptyError",
	Desc: "The request expects public.name not to be empty. public.name was found to be empty. Therefore the request failed.",
}

func IsSearchPublicEmpty(err error) bool {
	return errors.Is(err, searchPublicEmptyError)
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
