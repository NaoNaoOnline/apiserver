package walletstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var walletAddrChangeError = &tracer.Error{
	Kind: "walletAddrChangeError",
	Desc: "The request expects the wallet address not to change during consecutive signature verification challenges. The wallet address was found to have changed during the current signature verification challenge. Therefore the request failed.",
}

func IsWalletAddrChange(err error) bool {
	return errors.Is(err, walletAddrChangeError)
}

var walletKindInvalidError = &tracer.Error{
	Kind: "walletKindInvalidError",
	Desc: "The request expects the wallet kind to be one of [eth]. The wallet kind was not found to be one of [eth]. Therefore the request failed.",
}

func IsWalletKindInvalid(err error) bool {
	return errors.Is(err, walletKindInvalidError)
}

var walletMessEmptyError = &tracer.Error{
	Kind: "walletMessEmptyError",
	Desc: "The request expects the wallet message not to be empty. The wallet message was found to be empty. Therefore the request failed.",
}

func IsWalletMessEmpty(err error) bool {
	return errors.Is(err, walletMessEmptyError)
}

var walletPubkEmptyError = &tracer.Error{
	Kind: "walletPubkEmptyError",
	Desc: "The request expects the wallet public key not to be empty. The wallet public key was found to be empty. Therefore the request failed.",
}

func IsWalletPubkEmpty(err error) bool {
	return errors.Is(err, walletPubkEmptyError)
}

var walletPubkFormatError = &tracer.Error{
	Kind: "walletPubkFormatError",
	Desc: "The request expects the wallet public key to be in hex format including 0x prefix. The wallet public key was not found to be in hex format including 0x prefix. Therefore the request failed.",
}

func IsWalletPubkFormat(err error) bool {
	return errors.Is(err, walletPubkFormatError)
}

var walletPubkLengthError = &tracer.Error{
	Kind: "walletPubkLengthError",
	Desc: "The request expects the wallet public key to have 132 characters. The wallet public key was not found to have 132 characters. Therefore the request failed.",
}

func IsWalletPubkLength(err error) bool {
	return errors.Is(err, walletPubkLengthError)
}

var walletSignEmptyError = &tracer.Error{
	Kind: "walletSignEmptyError",
	Desc: "The request expects the wallet signature not to be empty. The wallet signature was found to be empty. Therefore the request failed.",
}

func IsWalletSignEmpty(err error) bool {
	return errors.Is(err, walletSignEmptyError)
}

var walletSignFormatError = &tracer.Error{
	Kind: "walletSignFormatError",
	Desc: "The request expects the wallet signature to be in hex format including 0x prefix. The wallet signature was not found to be in hex format including 0x prefix. Therefore the request failed.",
}

func IsWalletSignFormat(err error) bool {
	return errors.Is(err, walletSignFormatError)
}

var walletSignLengthError = &tracer.Error{
	Kind: "walletSignLengthError",
	Desc: "The request expects the wallet signature to have 132 characters. The wallet signature was not found to have 132 characters. Therefore the request failed.",
}

func IsWalletSignLength(err error) bool {
	return errors.Is(err, walletSignLengthError)
}

var walletSignatureInvalidError = &tracer.Error{
	Kind: "walletSignatureInvalidError",
	Desc: "The request expects the wallet signature to be valid. The wallet signature was not found to be valid. Therefore the request failed.",
}

func IsWalletSignatureInvalid(err error) bool {
	return errors.Is(err, walletSignatureInvalidError)
}
