package labelstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var fieldUnsupportedError = &tracer.Error{
	Kind: "fieldUnsupportedError",
	Desc: "Neither desc, disc nor twit are supported fields right now, you shadowy super coder. Let's talk!",
}

func IsFieldUnsupported(err error) bool {
	return errors.Is(err, fieldUnsupportedError)
}

var labelKindInvalidError = &tracer.Error{
	Kind: "labelKindInvalidError",
	Desc: "The request expects the label kind to be one of [bltn cate host]. The label kind was not found to be one of [bltn cate host]. Therefore the request failed.",
}

func IsLabelKindInvalid(err error) bool {
	return errors.Is(err, labelKindInvalidError)
}

var labelNameEmptyError = &tracer.Error{
	Kind: "labelNameEmptyError",
	Desc: "The request expects the label name not to be empty. The label name was found to be empty. Therefore the request failed.",
}

func IsLabelNameEmpty(err error) bool {
	return errors.Is(err, labelNameEmptyError)
}

var labelNameFormatError = &tracer.Error{
	Kind: "labelNameFormatError",
	Desc: "The request expects the label name to contain words or numbers. The label name was found to contain invalid characters. Therefore the request failed.",
}

func IsLabelNameFormat(err error) bool {
	return errors.Is(err, labelNameFormatError)
}

var labelNameLengthError = &tracer.Error{
	Kind: "labelNameLengthError",
	Desc: "The request expects the label name to have between 3 and 20 characters. The label name was not found to have between 3 and 20 characters. Therefore the request failed.",
}

func IsDescriptionNameLength(err error) bool {
	return errors.Is(err, labelNameLengthError)
}
