package policystorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/format/hexformat"
	"github.com/xh3b4sd/tracer"
)

var policyAcceNegativeError = &tracer.Error{
	Kind: "policyAcceNegativeError",
	Desc: "The request expects the policy access to be a positive integer. The policy access was not found to be a positive integer. Therefore the request failed.",
}

var policyChIDDuplicateError = &tracer.Error{
	Kind: "policyChIDDuplicateError",
	Desc: "The request expects the policy chain ID not to be duplicated. The policy chain ID was found to be duplicated. Therefore the request failed.",
}

var policyChIDEmptyError = &tracer.Error{
	Kind: "policyChIDEmptyError",
	Desc: "The request expects the policy chain ID not to be empty. The policy chain ID was found to be empty. Therefore the request failed.",
}

var policyMembEmptyError = &tracer.Error{
	Kind: "policyMembEmptyError",
	Desc: "The request expects the policy member not to be empty. The policy member was found to be empty. Therefore the request failed.",
}

var policyMembFormatError = hexformat.Errorf("policy", "memb")

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
