package userhandler

import (
	"github.com/xh3b4sd/tracer"
)

var searchInternConflictError = &tracer.Error{
	Kind: "searchInternConflictError",
	Desc: "The request expects intern.user to be the only field provided within the given query object. Fields other than intern.user were found to be set within the given query object. Therefore the request failed.",
}

var searchInternEmptyError = &tracer.Error{
	Kind: "searchInternEmptyError",
	Desc: "The request expects intern.user not to be empty. intern.user was found to be empty. Therefore the request failed.",
}

var searchPublicConflictError = &tracer.Error{
	Kind: "searchPublicConflictError",
	Desc: "The request expects public.name to be the only field provided within the given query object. Fields other than public.name were found to be set within the given query object. Therefore the request failed.",
}

var searchPublicEmptyError = &tracer.Error{
	Kind: "searchPublicEmptyError",
	Desc: "The request expects public.name not to be empty. public.name was found to be empty. Therefore the request failed.",
}

var searchSymbolConflictError = &tracer.Error{
	Kind: "searchSymbolConflictError",
	Desc: "The request expects symbol.user to be the only field provided within the given query object. Fields other than symbol.user were found to be set within the given query object. Therefore the request failed.",
}

var searchSymbolEmptyError = &tracer.Error{
	Kind: "searchSymbolEmptyError",
	Desc: "The request expects symbol.user not to be empty. symbol.user was found to be empty. Therefore the request failed.",
}

var searchSymbolInvalidError = &tracer.Error{
	Kind: "searchSymbolInvalidError",
	Desc: `The request expects symbol.user to be set to "self". symbol.user was not found to be set to "self". Therefore the request failed.`,
}

var updateEmptyError = &tracer.Error{
	Kind: "updateEmptyError",
	Desc: "The request expects the query object to contain all of [intern update]. The query object was not found to contain all of [intern update]. Therefore the request failed.",
}

var nameUpdatePeriodError = &tracer.Error{
	Kind: "descriptionUpdatePeriodError",
	Desc: "The request expects the user name to be updated once within the past 7 days. The user name was tried to be updated more than once within the past 7 days. Therefore the request failed.",
}
