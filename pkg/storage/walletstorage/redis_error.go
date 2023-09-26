package walletstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var walletObjectNotFoundError = &tracer.Error{
	Kind: "walletObjectNotFoundError",
	Desc: "The request expects a wallet object to exist. The wallet object was not found to exist. Therefore the request failed.",
}

func IsWalletObjectNotFound(err error) bool {
	return errors.Is(err, walletObjectNotFoundError)
}

var walletUserLimitError = &tracer.Error{
	Kind: "walletUserLimitError",
	Desc: "The request expects an upper limit of 5 wallet objects per user globally. The upper limit of 5 wallet objects per user globally was found. Therefore the request failed.",
}

func IsWalletUserLimit(err error) bool {
	return errors.Is(err, walletUserLimitError)
}
