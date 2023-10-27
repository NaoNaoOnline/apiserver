package labelhandler

import (
	"github.com/xh3b4sd/tracer"
)

var createKindInvalidError = &tracer.Error{
	Kind: "createKindInvalidError",
	Desc: "The request expects the label kind to be one of [cate host]. The label kind was not found to be one of [cate host]. Therefore the request failed.",
}
