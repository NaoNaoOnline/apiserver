package liststorage

import (
	"errors"

	"github.com/NaoNaoOnline/apiserver/pkg/format/descriptionformat"
	"github.com/xh3b4sd/tracer"
)

var listDescEmptyError = &tracer.Error{
	Kind: "listDescEmptyError",
	Desc: "The request expects the list description not to be empty. The list description was found to be empty. Therefore the request failed.",
}

var listDescFormatError = descriptionformat.Errorf("list", "desc")

var listDescLengthError = &tracer.Error{
	Kind: "listDescLengthError",
	Desc: "The request expects the list description to have between 2 and 40 characters. The list description was not found to have between 3 and 20 characters. Therefore the request failed.",
}

func IsListDescLength(err error) bool {
	return errors.Is(err, listDescLengthError)
}
