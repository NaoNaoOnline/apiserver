package eventhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var searchInternConflictError = &tracer.Error{
	Kind: "searchInternConflictError",
	Desc: "The request expects the query object to contain one of [intern.evnt intern.user]. The query object was not found to contain one of [intern.evnt intern.user]. Therefore the request failed.",
}

func IsSearchInternConflict(err error) bool {
	return errors.Is(err, searchInternConflictError)
}

var searchInternEmptyError = &tracer.Error{
	Kind: "searchInternEmptyError",
	Desc: "The request expects the query object to contain one of [intern.evnt intern.user]. The query object was not found to contain one of [intern.evnt intern.user]. Therefore the request failed.",
}

func IsSearchInternEmpty(err error) bool {
	return errors.Is(err, searchInternEmptyError)
}

var queryObjectConflictError = &tracer.Error{
	Kind: "queryObjectConflictError",
	Desc: "The request expects the query object to contain one of [intern public symbol]. The query object was not found to contain one of [intern public symbol]. Therefore the request failed.",
}

func IsQueryObjectConflictError(err error) bool {
	return errors.Is(err, queryObjectConflictError)
}

var searchPublicEmptyError = &tracer.Error{
	Kind: "searchPublicEmptyError",
	Desc: "The request expects the query object to contain one of [public.cate public.host]. The query object was not found to contain one of [public.cate public.host]. Therefore the request failed.",
}

func IsSearchPublicEmpty(err error) bool {
	return errors.Is(err, searchPublicEmptyError)
}

var searchSymbolConflictError = &tracer.Error{
	Kind: "searchSymbolConflictError",
	Desc: "The request expects the query object to contain one of [symbol.ltst symbol.rctn]. The query object was not found to contain one of [symbol.ltst symbol.rctn]. Therefore the request failed.",
}

func IsSearchSymbolConflict(err error) bool {
	return errors.Is(err, searchSymbolConflictError)
}

var searchSymbolEmptyError = &tracer.Error{
	Kind: "searchSymbolEmptyError",
	Desc: "The request expects the query object to contain one of [symbol.ltst symbol.rctn]. The query object was not found to contain one of [symbol.ltst symbol.rctn]. Therefore the request failed.",
}

func IsSearchSymbolEmpty(err error) bool {
	return errors.Is(err, searchSymbolEmptyError)
}
