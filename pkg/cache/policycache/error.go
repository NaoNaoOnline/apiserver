package policycache

import "github.com/xh3b4sd/tracer"

var policyBufferEmptyError = &tracer.Error{
	Kind: "policyBufferEmptyError",
	Desc: "The request expects the policy buffer not to be empty. The policy buffer was found to be empty. Therefore the request failed.",
}
