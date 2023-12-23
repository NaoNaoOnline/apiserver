package eventtemplate

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var cancelError = &tracer.Error{
	Kind: "cancelError",
	Desc: "Cancel is the error returned if resources were found to be missing, which are in fact required for the requested template to be generated. E.g. it might be that event objects or description objects are not available anymore by the time the template is tried to be generated.",
}

func IsCancel(err error) bool {
	return errors.Is(err, cancelError)
}
