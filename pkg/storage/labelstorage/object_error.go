package labelstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/format/descriptionformat"
	"github.com/NaoNaoOnline/apiserver/pkg/format/handleformat"
	"github.com/NaoNaoOnline/apiserver/pkg/format/nameformat"
	"github.com/xh3b4sd/tracer"
)

var labelDescFormatError = descriptionformat.Errorf("label", "desc")

var labelKindInvalidError = &tracer.Error{
	Kind: "labelKindInvalidError",
	Desc: "The request expects the label kind to be one of [bltn cate host]. The label kind was not found to be one of [bltn cate host]. Therefore the request failed.",
}

var labelNameEmptyError = &tracer.Error{
	Kind: "labelNameEmptyError",
	Desc: "The request expects the label name not to be empty. The label name was found to be empty. Therefore the request failed.",
}

var labelNameFormatError = nameformat.Errorf("label", "name")

var labelNameLengthError = &tracer.Error{
	Kind: "labelNameLengthError",
	Desc: "The request expects the label name to have between 2 and 25 characters. The label name was not found to have between 3 and 20 characters. Therefore the request failed.",
}

var labelPrflFormatError = handleformat.Errorf("label", "prfl")

var labelPrflInvalidError = &tracer.Error{
	Kind: "labelPrflInvalidError",
	Desc: "The request expects the label prfl to be one of [Twitter Warpcast]. The label prfl was not found to be one of [Twitter Warpcast]. Therefore the request failed.",
}
