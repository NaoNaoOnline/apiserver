package descriptionstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var descriptionTextEmptyError = &tracer.Error{
	Kind: "descriptionTextEmptyError",
	Desc: "The request expects the description text not to be empty. The description text was found to be empty for the request. Therefore the request failed.",
}

func IsDescriptionTextEmpty(err error) bool {
	return errors.Is(err, descriptionTextEmptyError)
}

var descriptionTextFormatError = &tracer.Error{
	Kind: "descriptionTextFormatError",
	Desc: `The request expects the description text to contain words, numbers or: , . : - ' " ! $ % & #. The description text was found to contain invalid characters for the request. Therefore the request failed.`,
}

func IsDescriptionTextFormat(err error) bool {
	return errors.Is(err, descriptionTextFormatError)
}

var descriptionTextLengthError = &tracer.Error{
	Kind: "descriptionTextLengthError",
	Desc: "The request expects the description text to have between 20 and 120 characters. The description text was not found to have between 20 and 120 characters for the request. Therefore the request failed.",
}

func IsDescriptionTextLength(err error) bool {
	return errors.Is(err, descriptionTextLengthError)
}

var eventIDEmptyError = &tracer.Error{
	Kind: "eventIDEmptyError",
	Desc: "The request expects the event ID not to be empty. The event ID was found to be empty for the request. Therefore the request failed.",
}

func IsEventIDEmpty(err error) bool {
	return errors.Is(err, eventIDEmptyError)
}
