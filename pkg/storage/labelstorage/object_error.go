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

var invalidLabelKindError = &tracer.Error{
	Kind: "invalidLabelKindError",
	Desc: "The request expects a valid label kind for the label object, e.g. cate or host. No valid label kind was found for the request. Therefore it failed.",
}

func IsInvalidLabelKind(err error) bool {
	return errors.Is(err, invalidLabelKindError)
}

var labelNameEmptyError = &tracer.Error{
	Kind: "labelNameEmptyError",
	Desc: "The request expects a valid label name for the label object. No label name was found for the request. Therefore it failed.",
}

func IsLabelNameEmpty(err error) bool {
	return errors.Is(err, labelNameEmptyError)
}

var labelNameFormatError = &tracer.Error{
	Kind: "labelNameFormatError",
	Desc: "The request expects the label name to contain words or numbers. The label name was found to contain invalid characters for the request. Therefore it failed.",
}

func IsLabelNameFormat(err error) bool {
	return errors.Is(err, labelNameFormatError)
}
