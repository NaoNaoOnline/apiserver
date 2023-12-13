package userstorage

import (
	"errors"

	"github.com/NaoNaoOnline/apiserver/pkg/format/nameformat"
	"github.com/xh3b4sd/tracer"
)

var userImageEmptyError = &tracer.Error{
	Kind: "userImageEmptyError",
	Desc: "The request expects the user image not to be empty. The user image was found to be empty. Therefore the request failed.",
}

func IsUserImageEmpty(err error) bool {
	return errors.Is(err, userImageEmptyError)
}

var userNameEmptyError = &tracer.Error{
	Kind: "userNameEmptyError",
	Desc: "The request expects the user name not to be empty. The user name was found to be empty. Therefore the request failed.",
}

func IsUserNameEmpty(err error) bool {
	return errors.Is(err, userNameEmptyError)
}

var userNameLengthError = &tracer.Error{
	Kind: "userNameLengthError",
	Desc: "The request expects the user name to have between 2 and 30 characters. The user name was not found to have between 2 and 30 characters. Therefore the request failed.",
}

func IsUserNameLength(err error) bool {
	return errors.Is(err, userNameLengthError)
}

var userPrflFormatError = nameformat.Errorf("user", "prfl")

var userPrflInvalidError = &tracer.Error{
	Kind: "userPrflInvalidError",
	Desc: "The request expects the user prfl to be one of [Twitter Warpcast]. The user prfl was not found to be one of [Twitter Warpcast]. Therefore the request failed.",
}

var userSubjectEmptyError = &tracer.Error{
	Kind: "userSubjectEmptyError",
	Desc: "The request expects the user's subject claim not to be empty. The user's subject claim was found to be empty. Therefore the request failed.",
}

func IsUserSubjectEmpty(err error) bool {
	return errors.Is(err, userSubjectEmptyError)
}
