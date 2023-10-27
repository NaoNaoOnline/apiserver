package rulestorage

import "github.com/xh3b4sd/tracer"

var listIDEmptyError = &tracer.Error{
	Kind: "eventIDEmptyError",
	Desc: "The request expects the list ID not to be empty. The list ID was found to be empty. Therefore the request failed.",
}

var resourceIDEmptyError = &tracer.Error{
	Kind: "resourceIDEmptyError",
	Desc: "The request expects the rule object to contain one of [Object.Excl Object.Incl]. The rule object was not found to contain one of [Object.Excl Object.Incl]. Therefore the request failed.",
}

var ruleKindInvalidError = &tracer.Error{
	Kind: "ruleKindInvalidError",
	Desc: "The request expects the rule kind to be one of [cate host user]. The rule kind was not found to be one of [cate host user]. Therefore the request failed.",
}

var ruleListEmptyError = &tracer.Error{
	Kind: "ruleListEmptyError",
	Desc: "The request expects the rule list not to be empty. The rule list was found to be empty. Therefore the request failed.",
}
