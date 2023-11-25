package subscriptionstorage

import (
	"github.com/xh3b4sd/tracer"
)

var subscriptionChIDEmptyError = &tracer.Error{
	Kind: "subscriptionChIDEmptyError",
	Desc: "The request expects the subscription chain ID not to be empty. The subscription chain ID was found to be empty. Therefore the request failed.",
}
