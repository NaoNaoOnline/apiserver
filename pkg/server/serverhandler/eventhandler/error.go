package eventhandler

import (
	"github.com/xh3b4sd/tracer"
)

var eventDeletedError = &tracer.Error{
	Kind: "eventDeletedError",
	Desc: "The request expects event objects to be clicked until they are deleted. The event object was found to have already been deleted. Therefore the request failed.",
}

var queryObjectConflictError = &tracer.Error{
	Kind: "queryObjectConflictError",
	Desc: "The request expects the query object to contain one of [intern public symbol]. The query object was not found to contain one of [intern public symbol]. Therefore the request failed.",
}

var searchFeedPremiumError = &tracer.Error{
	Kind: "searchFeedPremiumError",
	Desc: "The request expects the user to have a premium subscription in order to receive list notifications. The user was not found to have a premium subscription. Therefore the request failed.",
}

var searchInternConflictError = &tracer.Error{
	Kind: "searchInternConflictError",
	Desc: "The request expects the query object to contain one of [intern.evnt intern.user]. The query object was not found to contain one of [intern.evnt intern.user]. Therefore the request failed.",
}

var searchInternEmptyError = &tracer.Error{
	Kind: "searchInternEmptyError",
	Desc: "The request expects the query object to contain one of [intern.evnt intern.user]. The query object was not found to contain one of [intern.evnt intern.user]. Therefore the request failed.",
}

var searchPublicEmptyError = &tracer.Error{
	Kind: "searchPublicEmptyError",
	Desc: "The request expects the query object to contain one of [public.cate public.host]. The query object was not found to contain one of [public.cate public.host]. Therefore the request failed.",
}

var searchSymbolConflictError = &tracer.Error{
	Kind: "searchSymbolConflictError",
	Desc: "The request expects the query object to contain one of [symbol.like symbol.list symbol.time]. The query object was not found to contain one of [symbol.like symbol.list symbol.time]. Therefore the request failed.",
}

var searchSymbolEmptyError = &tracer.Error{
	Kind: "searchSymbolEmptyError",
	Desc: "The request expects the query object to contain one of [symbol.like symbol.list symbol.time]. The query object was not found to contain one of [symbol.like symbol.list symbol.time]. Therefore the request failed.",
}

var searchSymbolPageError = &tracer.Error{
	Kind: "searchSymbolPageError",
	Desc: "The request expects the query filter to contain all of [filter.strt filter.stop] if one of [symbol.like symbol.time] is set to page. The query filter was not found to contain all of [filter.strt filter.stop]. Therefore the request failed.",
}

var searchSymbolTimeError = &tracer.Error{
	Kind: "searchSymbolTimeError",
	Desc: "The request expects the query object to contain one of [dflt hpnd upcm] if symbol.time is configured. The query object was not found to contain one of [dflt hpnd upcm]. Therefore the request failed.",
}

var updateSymbolInvalidError = &tracer.Error{
	Kind: "updateSymbolInvalidError",
	Desc: "The request expects symbol.link to be one of [add]. symbol.link was not found to be one of [add]. Therefore the request failed.",
}
