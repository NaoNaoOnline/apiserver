package policycache

import (
	"github.com/xh3b4sd/tracer"
)

var policyAcceNegativeError = &tracer.Error{
	Kind: "policyAcceNegativeError",
	Desc: "The request expects the policy access to be a positive integer. The policy access was not found to be a positive integer. Therefore the request failed.",
}

var policyBufferEmptyError = &tracer.Error{
	Kind: "policyBufferEmptyError",
	Desc: "The request expects the policy buffer not to be empty when merging policy records from multiple chains. The policy buffer was found to be empty when merging policy records from multiple chains. Therefore the request failed.",
}

var policyChIDDuplicateError = &tracer.Error{
	Kind: "policyChIDDuplicateError",
	Desc: "The request expects the policy chain ID not to be duplicated. The policy chain ID was found to be duplicated. Therefore the request failed.",
}

var policyChIDEmptyError = &tracer.Error{
	Kind: "policyChIDEmptyError",
	Desc: "The request expects the policy chain ID not to be empty. The policy chain ID was found to be empty. Therefore the request failed.",
}

var policyChIDInvalidError = &tracer.Error{
	Kind: "policyChIDInvalidError",
	Desc: "The request expects the same chain ID accross all records when buffering policies. Different chain IDs accross all records were found when buffering policies. Therefore the request failed.",
}

var policyChIDLimitError = &tracer.Error{
	Kind: "policyChIDLimitError",
	Desc: "The request expects a single chain ID per record when buffering policies. More than one chain ID per record was found when buffering policies. Therefore the request failed.",
}

var policyMembEmptyError = &tracer.Error{
	Kind: "policyMembEmptyError",
	Desc: "The request expects the policy member not to be empty. The policy member was found to be empty. Therefore the request failed.",
}

var policyMembFormatError = &tracer.Error{
	Kind: "policyMembFormatError",
	Desc: "The request expects the policy member to be in hex format including 0x prefix. The policy member was not found to be in hex format including 0x prefix. Therefore the request failed.",
}

var policyMembLengthError = &tracer.Error{
	Kind: "policyMembLengthError",
	Desc: "The request expects the policy member to have 42 characters. The policy member was not found to have 42 characters. Therefore the request failed.",
}

var policyRecordEmptyError = &tracer.Error{
	Kind: "policyRecordEmptyError",
	Desc: "The request expects the policy record not to be empty. The policy record was found to be empty. Therefore the request failed.",
}

var policySystNegativeError = &tracer.Error{
	Kind: "policySystNegativeError",
	Desc: "The request expects the policy system to be a positive integer. The policy system was not found to be a positive integer. Therefore the request failed.",
}
