package rulehandler

import (
	"github.com/xh3b4sd/tracer"
)

var createKindInvalidError = &tracer.Error{
	Kind: "createKindInvalidError",
	Desc: "The request expects the rule kind to be one of [cate host user]. The rule kind was not found to be one of [cate host user]. Therefore the request failed.",
}

var resourceIDEmptyError = &tracer.Error{
	Kind: "resourceIDEmptyError",
	Desc: "The request expects the query object to contain one of [public.excl public.incl]. The query object was not found to contain one of [public.excl public.incl]. Therefore the request failed.",
}