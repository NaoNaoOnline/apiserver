package policystorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var policyAcceNegativeError = &tracer.Error{
	Kind: "policyAcceNegativeError",
	Desc: "The request expects the policy access to be a positive integer. The policy access was not found to be a positive integer. Therefore the request failed.",
}

func IsPolicyAcceNegative(err error) bool {
	return errors.Is(err, policyAcceNegativeError)
}

var policyTimeEmptyError = &tracer.Error{
	Kind: "policyTimeEmptyError",
	Desc: "The request expects the external policy creation timestamp not to be empty. The external policy creation timestamp was found to be empty. Therefore the request failed.",
}

func IsPolicyTimeEmpty(err error) bool {
	return errors.Is(err, policyTimeEmptyError)
}

var policyKindInvalidError = &tracer.Error{
	Kind: "policyKindInvalidError",
	Desc: "The request expects the policy kind to be one of [CreateMember CreateSystem DeleteMember DeleteSystem]. The policy kind was not found to be one of [CreateMember CreateSystem DeleteMember DeleteSystem]. Therefore the request failed.",
}

func IsPolicyKindInvalid(err error) bool {
	return errors.Is(err, policyKindInvalidError)
}

var policyMembEmptyError = &tracer.Error{
	Kind: "policyMembEmptyError",
	Desc: "The request expects the policy member not to be empty. The policy member was found to be empty. Therefore the request failed.",
}

func IsPolicyMembEmpty(err error) bool {
	return errors.Is(err, policyMembEmptyError)
}

var policyMembFormatError = &tracer.Error{
	Kind: "policyMembFormatError",
	Desc: "The request expects the policy member to be in hex format including 0x prefix. The policy member was not found to be in hex format including 0x prefix. Therefore the request failed.",
}

func IsPolicyMembFormat(err error) bool {
	return errors.Is(err, policyMembFormatError)
}

var policyMembLengthError = &tracer.Error{
	Kind: "policyMembLengthError",
	Desc: "The request expects the policy member to have 42 characters. The policy member was not found to have 42 characters. Therefore the request failed.",
}

func IsPolicyMembLength(err error) bool {
	return errors.Is(err, policyMembLengthError)
}

var policySystNegativeError = &tracer.Error{
	Kind: "policySystNegativeError",
	Desc: "The request expects the policy system to be a positive integer. The policy system was not found to be a positive integer. Therefore the request failed.",
}

func IsPolicySystNegative(err error) bool {
	return errors.Is(err, policySystNegativeError)
}
