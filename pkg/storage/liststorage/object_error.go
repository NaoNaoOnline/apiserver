package liststorage

import (
	"github.com/xh3b4sd/tracer"
)

var listDescEmptyError = &tracer.Error{
	Kind: "listDescEmptyError",
	Desc: "The request expects the list description not to be empty. The list description was found to be empty. Therefore the request failed.",
}

var listDescFormatError = &tracer.Error{
	Kind: "listDescFormatError",
	Desc: `The request expects the list description to contain words, numbers or: , . : - ' " ! $ % & #. The list description was found to contain invalid characters. Therefore the request failed.`,
}

var listDescLengthError = &tracer.Error{
	Kind: "listDescLengthError",
	Desc: "The request expects the list description to have between 3 and 120 characters. The list description was not found to have between 3 and 20 characters. Therefore the request failed.",
}
