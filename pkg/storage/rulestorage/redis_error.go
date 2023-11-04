package rulestorage

import (
	"github.com/xh3b4sd/tracer"
)

var listObjectNotFoundError = &tracer.Error{
	Kind: "listObjectNotFoundError",
	Desc: "The request expects list object referenced in the rule object to exist. The list object referenced in the rule object was not found to exist. Therefore the request failed.",
}

var ruleObjectNotFoundError = &tracer.Error{
	Kind: "ruleObjectNotFoundError",
	Desc: "The request expects a rule object to exist. The rule object was not found to exist. Therefore the request failed.",
}
