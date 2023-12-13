package listhandler

import (
	"github.com/xh3b4sd/tracer"
)

var createDescEmptyError = &tracer.Error{
	Kind: "createDescEmptyError",
	Desc: "The request expects the list description not to be empty. The list description was found to be empty. Therefore the request failed.",
}

var createListLimitError = &tracer.Error{
	Kind: "createListLimitError",
	Desc: "The request expects an upper limit of 50 list objects per user. The upper limit of 50 list objects per user was found. Therefore the request failed.",
}

var createListPremiumError = &tracer.Error{
	Kind: "createListPremiumError",
	Desc: "The request expects the user to have a premium subscription in order to create more than 1 list. The user was not found to have a premium subscription. Therefore the request failed.",
}

var updateEmptyError = &tracer.Error{
	Kind: "updateEmptyError",
	Desc: "The request expects the query object to contain all of [intern update]. The query object was not found to contain all of [intern update]. Therefore the request failed.",
}

var updateFeedPremiumError = &tracer.Error{
	Kind: "updateFeedPremiumError",
	Desc: "The request expects the user to have a premium subscription in order to set list notifications. The user was not found to have a premium subscription. Therefore the request failed.",
}
