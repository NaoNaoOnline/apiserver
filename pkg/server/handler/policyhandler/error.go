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

var queryObjectEmptyError = &tracer.Error{
	Kind: "queryObjectEmptyError",
	Desc: "The request expects the query object not to be empty. The query object was found to be empty. Therefore the request failed.",
}

func IsQueryObjectEmpty(err error) bool {
	return errors.Is(err, queryObjectEmptyError)
}

var searchKindConflictError = &tracer.Error{
	Kind: "searchKindConflictError",
	Desc: "The request expects public.kind not to be provided if symbol.ltst is one of [default aggregated]. public.kind was found to be provided when symbol.ltst was one of [default aggregated]. Therefore the request failed.",
}

func IsSearchKindConflict(err error) bool {
	return errors.Is(err, searchKindConflictError)
}

var searchLtstEmptyError = &tracer.Error{
	Kind: "searchLtstEmptyError",
	Desc: "The request expects symbol.ltst not to be empty. symbol.ltst was found to be empty. Therefore the request failed.",
}

func IsSearchLtstEmpty(err error) bool {
	return errors.Is(err, searchLtstEmptyError)
}

var updateSyncInvalidError = &tracer.Error{
	Kind: "updateSyncInvalidError",
	Desc: `The request expects symbol.sync to be set to "default". symbol.sync was not found to be set to "default". Therefore the request failed.`,
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
