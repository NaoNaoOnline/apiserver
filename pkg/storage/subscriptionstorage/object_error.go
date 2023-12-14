package subscriptionstorage

import (
	"errors"

	"github.com/NaoNaoOnline/apiserver/pkg/format/hexformat"
	"github.com/xh3b4sd/tracer"
)

var subscriptionCrtrDuplicateError = &tracer.Error{
	Kind: "subscriptionCrtrDuplicateError",
	Desc: "The request expects the creator addresses not to be duplicated. The creator addresses were found to be duplicated. Therefore the request failed.",
}

var subscriptionCrtrEmptyError = &tracer.Error{
	Kind: "subscriptionCrtrEmptyError",
	Desc: "The request expects the subscription creator not to be empty. The subscription creator was found to be empty. Therefore the request failed.",
}

var subscriptionCrtrFormatError = hexformat.Errorf("subscription", "crtr")

var subscriptionCrtrLengthError = &tracer.Error{
	Kind: "subscriptionCrtrLengthError",
	Desc: "The request expects the subscription creator to have 42 characters. The subscription creator was not found to have 42 characters. Therefore the request failed.",
}

var subscriptionCrtrLimitError = &tracer.Error{
	Kind: "subscriptionCrtrLimitError",
	Desc: "The request expects an upper limit of 3 creator addresses per subscription. More than 3 creator addresses per subscription were found. Therefore the request failed.",
}

var subscriptionPayrEmptyError = &tracer.Error{
	Kind: "subscriptionPayrEmptyError",
	Desc: "The request expects the subscription payer not to be empty. The subscription payer was found to be empty. Therefore the request failed.",
}

var subscriptionRcvrEmptyError = &tracer.Error{
	Kind: "subscriptionRcvrEmptyError",
	Desc: "The request expects the subscription receiver not to be empty. The subscription receiver was found to be empty. Therefore the request failed.",
}

var subscriptionUnixEmptyError = &tracer.Error{
	Kind: "subscriptionUnixEmptyError",
	Desc: "The request expects the subscription timestamp not to be empty. The subscription timestamp was found to be empty. Therefore the request failed.",
}

var subscriptionUnixInvalidError = &tracer.Error{
	Kind: "subscriptionUnixInvalidError",
	Desc: "The request expects the subscription timestamp to define the first day of the subscription period. The subscription timestamp was not found to define the first day of the subscription period. Therefore the request failed.",
}

func IsSubscriptionUnixInvalid(err error) bool {
	return errors.Is(err, subscriptionUnixInvalidError)
}

var subscriptionUnixRenewalError = &tracer.Error{
	Kind: "subscriptionUnixRenewalError",
	Desc: "The request expects the subscription to be renewed up to 7 days before the new subscription period. The subscription was tried to be renewed outside the 7 days before the new subscription period. Therefore the request failed.",
}
