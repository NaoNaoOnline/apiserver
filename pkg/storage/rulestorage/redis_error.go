package rulestorage

import (
	"github.com/xh3b4sd/tracer"
)

var ruleObjectNotFoundError = &tracer.Error{
	Kind: "ruleObjectNotFoundError",
	Desc: "The request expects a rule object to exist. The rule object was not found to exist. Therefore the request failed.",
}
