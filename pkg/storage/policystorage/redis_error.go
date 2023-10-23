package policystorage

import "github.com/xh3b4sd/tracer"

var policyChIDInvalidError = &tracer.Error{
	Kind: "policyChIDInvalidError",
	Desc: "The request expects the same chain ID accross all records when buffering policies. Different chain IDs accross all records were found when buffering policies. Therefore the request failed.",
}

var policyChIDLimitError = &tracer.Error{
	Kind: "policyChIDLimitError",
	Desc: "The request expects a single chain ID per record when buffering policies. More than one chain ID per record was found when buffering policies. Therefore the request failed.",
}
