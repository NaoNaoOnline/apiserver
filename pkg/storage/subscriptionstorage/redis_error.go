package subscriptionstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var subscriptionObjectNotFoundError = &tracer.Error{
	Kind: "subscriptionObjectNotFoundError",
	Desc: "The request expects a subscription object to exist. The subscription object was not found to exist. Therefore the request failed.",
}

func IsSubscriptionObjectNotFound(err error) bool {
	return errors.Is(err, subscriptionObjectNotFoundError)
}

var walletObjectNotFoundError = &tracer.Error{
	Kind: "walletObjectNotFoundError",
	Desc: "The request expects a wallet object to exist. The wallet object was not found to exist. Therefore the request failed.",
}
