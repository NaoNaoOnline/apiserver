package policyhandler

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var policyMemberError = &tracer.Error{
	Kind: "policyMemberError",
	Desc: "The request expects the caller to be a policy member. The caller was not found to be a policy member. Therefore the request failed.",
}

func IsPolicyMember(err error) bool {
	return errors.Is(err, policyMemberError)
}

var searchKindConflictError = &tracer.Error{
	Kind: "searchKindConflictError",
	Desc: "The request expects public.kind not to be provided if symbol.ltst is one of [default aggregated]. public.kind was found to be provided when symbol.ltst was one of [default aggregated]. Therefore the request failed.",
}

func IsSearchKindConflict(err error) bool {
	return errors.Is(err, searchKindConflictError)
}

var updateSyncInvalidError = &tracer.Error{
	Kind: "updateSyncInvalidError",
	Desc: `The request expects a single query context with symbol.sync set to "default". The request was not found to have a single query context with symbol.sync set to "default". Therefore the request failed.`,
}

func IsSearchSyncInvalid(err error) bool {
	return errors.Is(err, updateSyncInvalidError)
}

var userIDEmptyError = &tracer.Error{
	Kind: "userIDEmptyError",
	Desc: "The request expects a valid OAuth access token mapping to an internal user ID. No user ID was found. Therefore the request failed.",
}

func IsUserIDEmpty(err error) bool {
	return errors.Is(err, userIDEmptyError)
}