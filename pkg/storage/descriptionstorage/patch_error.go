package descriptionstorage

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var jsonPatchOperationEmptyError = &tracer.Error{
	Kind: "jsonPatchOperationEmptyError",
	Desc: "The request expects the JSON-Patch operation not to be empty. The JSON-Patch operation was found to be empty. Therefore the request failed.",
}

func IsJsonPatchOperationEmpty(err error) bool {
	return errors.Is(err, jsonPatchOperationEmptyError)
}

var jsonPatchOperationInvalidError = &tracer.Error{
	Kind: "jsonPatchOperationInvalidError",
	Desc: "The request expects the JSON-Patch operation to be one of [replace]. The JSON-Patch operation was not found to be one of [replace]. Therefore the request failed.",
}

func IsJsonPatchOperationInvalid(err error) bool {
	return errors.Is(err, jsonPatchOperationInvalidError)
}

var jsonPatchPathEmptyError = &tracer.Error{
	Kind: "jsonPatchPathEmptyError",
	Desc: "The request expects the JSON-Patch path not to be empty. The JSON-Patch path was found to be empty. Therefore the request failed.",
}

func IsJsonPatchPathEmpty(err error) bool {
	return errors.Is(err, jsonPatchPathEmptyError)
}

var jsonPatchPathInvalidError = &tracer.Error{
	Kind: "jsonPatchPathInvalidError",
	Desc: "The request expects the JSON-Patch path to be one of [/text]. The JSON-Patch path was not found to be one of [/text]. Therefore the request failed.",
}

func IsJsonPatchPathInvalid(err error) bool {
	return errors.Is(err, jsonPatchPathInvalidError)
}
