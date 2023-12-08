package feedstorage

import (
	"github.com/xh3b4sd/tracer"
)

var feedEvntEmptyError = &tracer.Error{
	Kind: "feedEvntEmptyError",
	Desc: "The request expects the feed event not to be empty. The feed event was found to be empty. Therefore the request failed.",
}

var feedKindInvalidError = &tracer.Error{
	Kind: "feedKindInvalidError",
	Desc: "The request expects the feed kind to be one of [cate evnt host user]. The feed kind was not found to be one of [cate evnt host user]. Therefore the request failed.",
}

var feedListEmptyError = &tracer.Error{
	Kind: "feedListEmptyError",
	Desc: "The request expects the feed list not to be empty. The feed list was found to be empty. Therefore the request failed.",
}

var feedObctEmptyError = &tracer.Error{
	Kind: "feedObctEmptyError",
	Desc: "The request expects the feed object not to be empty. The feed object was found to be empty. Therefore the request failed.",
}
