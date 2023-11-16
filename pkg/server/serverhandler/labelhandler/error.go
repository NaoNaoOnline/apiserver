package labelhandler

import (
	"github.com/xh3b4sd/tracer"
)

var createKindInvalidError = &tracer.Error{
	Kind: "createKindInvalidError",
	Desc: "The request expects the label kind to be one of [cate host]. The label kind was not found to be one of [cate host]. Therefore the request failed.",
}

var labelProfileAlreadyExistsError = &tracer.Error{
	Kind: "labelProfileAlreadyExistsError",
	Desc: "The request expects the provided label profile not to exist already. The provided label profile was found to already exist. Therefore the request failed.",
}

var labelProfileNotFoundError = &tracer.Error{
	Kind: "labelProfileNotFoundError",
	Desc: "The request expects the provided label profile to exist already. The provided label profile was not found to already exist. Therefore the request failed.",
}

var searchInternConflictError = &tracer.Error{
	Kind: "searchInternConflictError",
	Desc: "The request expects the query object to contain one of [intern.labl intern.user]. The query object was not found to contain one of [intern.labl intern.user]. Therefore the request failed.",
}
