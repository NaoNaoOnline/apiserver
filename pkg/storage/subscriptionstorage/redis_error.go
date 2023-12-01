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

var userNotFoundError = &tracer.Error{
	Kind: "userNotFoundError",
	Desc: "The request expected user objects to be found for the subscription payer and subscription receiver. The user objects for subscription payer and subscription receiver could not be found. Therefore the request failed.",
}
