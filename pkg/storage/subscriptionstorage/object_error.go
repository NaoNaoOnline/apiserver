package subscriptionstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/format/hexformat"
	"github.com/xh3b4sd/tracer"
)

var subscriptionChIDEmptyError = &tracer.Error{
	Kind: "subscriptionChIDEmptyError",
	Desc: "The request expects the subscription chain ID not to be empty. The subscription chain ID was found to be empty. Therefore the request failed.",
}

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

var subscriptionSbcbEmptyError = &tracer.Error{
	Kind: "subscriptionSbcbEmptyError",
	Desc: "The request expects the subscription subscriber not to be empty. The subscription subscriber was found to be empty. Therefore the request failed.",
}

var subscriptionSbcbFormatError = hexformat.Errorf("subscription", "sbcb")

var subscriptionSbcbLengthError = &tracer.Error{
	Kind: "subscriptionSbcbLengthError",
	Desc: "The request expects the subscription subscriber to have 42 characters. The subscription subscriber was not found to have 42 characters. Therefore the request failed.",
}

var subscriptionUnixEmptyError = &tracer.Error{
	Kind: "subscriptionUnixEmptyError",
	Desc: "The request expects the subscription timestamp not to be empty. The subscription timestamp was found to be empty. Therefore the request failed.",
}

var subscriptionUnixInvalidError = &tracer.Error{
	Kind: "subscriptionUnixInvalidError",
	Desc: "The request expects the subscription timestamp to define the first day of the current month. The subscription timestamp was not found to define the first day of the current month. Therefore the request failed.",
}
