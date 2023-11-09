package rulestorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var listObjectNotFoundError = &tracer.Error{
	Kind: "listObjectNotFoundError",
	Desc: "The request expects list object referenced in the rule object to exist. The list object referenced in the rule object was not found to exist. Therefore the request failed.",
}

var ruleListLimitError = &tracer.Error{
	Kind: "ruleListLimitError",
	Desc: "The request expects an upper limit of 100 rule objects per list. The upper limit of 100 rule objects per list was found. Therefore the request failed.",
}

func IsRuleListLimit(err error) bool {
	return errors.Is(err, ruleListLimitError)
}

var ruleObjectNotFoundError = &tracer.Error{
	Kind: "ruleObjectNotFoundError",
	Desc: "The request expects a rule object to exist. The rule object was not found to exist. Therefore the request failed.",
}
