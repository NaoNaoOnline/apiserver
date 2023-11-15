package descriptionstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/format/descriptionformat"
	"github.com/xh3b4sd/tracer"
)

var descriptionLikeNegativeError = &tracer.Error{
	Kind: "descriptionLikeNegativeError",
	Desc: "The request expects the description like not to be negative. The description like was found to be negative. Therefore the request failed.",
}

var descriptionTextEmptyError = &tracer.Error{
	Kind: "descriptionTextEmptyError",
	Desc: "The request expects the description text not to be empty. The description text was found to be empty. Therefore the request failed.",
}

var descriptionTextFormatError = descriptionformat.Errorf("description", "text")

var descriptionTextLengthError = &tracer.Error{
	Kind: "descriptionTextLengthError",
	Desc: "The request expects the description text to have between 20 and 120 characters. The description text was not found to have between 20 and 120 characters. Therefore the request failed.",
}

var descriptionTextURLError = &tracer.Error{
	Kind: "descriptionTextURLError",
	Desc: "The request expects the description text not to contain any URL. The description text was found to contain URLs. Therefore the request failed.",
}

var eventIDEmptyError = &tracer.Error{
	Kind: "eventIDEmptyError",
	Desc: "The request expects the event ID not to be empty. The event ID was found to be empty. Therefore the request failed.",
}
