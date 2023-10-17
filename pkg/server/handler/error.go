package handler

import (
	"github.com/xh3b4sd/tracer"
)

var ExecutionFailedError = &tracer.Error{
	Kind: "ExecutionFailedError",
	Desc: "This internal error implies a severe malfunction of the system.",
}

var PolicyMemberError = &tracer.Error{
	Kind: "PolicyMemberError",
	Desc: "The request expects the caller to be a policy member. The caller was not found to be a policy member. Therefore the request failed.",
}

var QueryObjectEmptyError = &tracer.Error{
	Kind: "QueryObjectEmptyError",
	Desc: "The request expects the query object not to be empty. The query object was found to be empty. Therefore the request failed.",
}

var UserIDEmptyError = &tracer.Error{
	Kind: "UserIDEmptyError",
	Desc: "The request expects a valid OAuth access token mapping to an internal user ID. No user ID was found. Therefore the request failed.",
}

var UserNotOwnerError = &tracer.Error{
	Kind: "UserNotOwnerError",
	Desc: "The request expects the calling user to be the owner of the requested resource. The calling user was not found to be the owner of the requested resource. Therefore the request failed.",
}
