package policyhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/cache/policycache"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type UpdateHandlerConfig struct {
	Cid []int64
	Log logger.Interface
	Pol policycache.Interface
}

type UpdateHandler struct {
	cid []int64
	log logger.Interface
	pol policycache.Interface
}

func NewUpdateHandler(c UpdateHandlerConfig) *UpdateHandler {
	if len(c.Cid) == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cid must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Pol == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Pol must not be empty", c)))
	}

	var han *UpdateHandler
	{
		han = &UpdateHandler{
			cid: c.Cid,
			log: c.Log,
			pol: c.Pol,
		}
	}

	return han
}
