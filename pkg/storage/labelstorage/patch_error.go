package labelstorage

import (
	"github.com/xh3b4sd/tracer"
)

var jsonPatchOperationEmptyError = &tracer.Error{
	Kind: "jsonPatchOperationEmptyError",
	Desc: "The request expects the JSON-Patch operation not to be empty. The JSON-Patch operation was found to be empty. Therefore the request failed.",
}

var jsonPatchOperationInvalidError = &tracer.Error{
	Kind: "jsonPatchOperationInvalidError",
	Desc: "The request expects the JSON-Patch operation to be one of [add remove]. The JSON-Patch operation was not found to be one of [add remove]. Therefore the request failed.",
}

var jsonPatchPathEmptyError = &tracer.Error{
	Kind: "jsonPatchPathEmptyError",
	Desc: "The request expects the JSON-Patch path not to be empty. The JSON-Patch path was found to be empty. Therefore the request failed.",
}

var jsonPatchPathInvalidError = &tracer.Error{
	Kind: "jsonPatchPathInvalidError",
	Desc: "The request expects the JSON-Patch path prefix to be one of [/prfl/data/]. The JSON-Patch path prefix was not found to be one of [/prfl/data/]. Therefore the request failed.",
}
