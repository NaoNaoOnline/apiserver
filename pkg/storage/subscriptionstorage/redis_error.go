package subscriptionstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var subscriptionCrtrInvalidError = &tracer.Error{
	Kind: "subscriptionCrtrInvalidError",
	Desc: "The request expects the creator addresses to be legitimate. The creator addresses were not found to be legitimate. Therefore the request failed.",
}

var subscriptionObjectNotFoundError = &tracer.Error{
	Kind: "subscriptionObjectNotFoundError",
	Desc: "The request expects a subscription object to exist. The subscription object was not found to exist. Therefore the request failed.",
}

func IsSubscriptionObjectNotFound(err error) bool {
	return errors.Is(err, subscriptionObjectNotFoundError)
}
