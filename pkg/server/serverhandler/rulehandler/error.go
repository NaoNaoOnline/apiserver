package rulehandler

import (
	"github.com/xh3b4sd/tracer"
)

var createKindInvalidError = &tracer.Error{
	Kind: "createKindInvalidError",
	Desc: "The request expects the rule kind to be one of [cate evnt host like user]. The rule kind was not found to be one of [cate evnt host like user]. Therefore the request failed.",
}

var listDeletedError = &tracer.Error{
	Kind: "listDeletedError",
	Desc: "The request expects rules to be added or removed until the associated list has been deleted. The associated list was found to have already been deleted. Therefore the request failed.",
}

var resourceIDEmptyError = &tracer.Error{
	Kind: "resourceIDEmptyError",
	Desc: "The request expects the query object to contain one of [public.excl public.incl]. The query object was not found to contain one of [public.excl public.incl]. Therefore the request failed.",
}
