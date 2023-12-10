package rulehandler

import (
	"github.com/xh3b4sd/tracer"
)

var createKindInvalidError = &tracer.Error{
	Kind: "createKindInvalidError",
	Desc: "The request expects the rule kind to be one of [cate evnt host user]. The rule kind was not found to be one of [cate evnt host user]. Therefore the request failed.",
}

var listDeletedError = &tracer.Error{
	Kind: "listDeletedError",
	Desc: "The request expects rules to be added or removed until the associated list has been deleted. The associated list was found to have already been deleted. Therefore the request failed.",
}

var resourceIDEmptyError = &tracer.Error{
	Kind: "resourceIDEmptyError",
	Desc: "The request expects the query object to contain one of [public.excl public.incl]. The query object was not found to contain one of [public.excl public.incl]. Therefore the request failed.",
}

var listRuleDuplicateError = &tracer.Error{
	Kind: "listRuleDuplicateError",
	Desc: "The request expects the list rule not to be duplicated. The list rule was found to be duplicated. Therefore the request failed.",
}

var ruleDynamicListError = &tracer.Error{
	Kind: "ruleDynamicListError",
	Desc: "The request expects dynamic rules to be added to dynamic lists. The rule to be added was found to be static. Therefore the request failed.",
}

var ruleListLimitError = &tracer.Error{
	Kind: "ruleListLimitError",
	Desc: "The request expects an upper limit of 100 rule objects per list. The upper limit of 100 rule objects per list was found. Therefore the request failed.",
}

var ruleStaticExclError = &tracer.Error{
	Kind: "ruleStaticExclError",
	Desc: "The request expects static rules not to define any excludes. The rule to be added was found to define excludes. Therefore the request failed.",
}

var ruleStaticListError = &tracer.Error{
	Kind: "ruleStaticListError",
	Desc: "The request expects static rules to be added to static lists. The rule to be added was found to be dynamic. Therefore the request failed.",
}
