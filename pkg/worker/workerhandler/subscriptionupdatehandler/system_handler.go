package subscriptionupdatehandler

import (
	"fmt"

	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type UpdateHandlerConfig struct {
	Cid []int64
	Loc locker.Interface
	Log logger.Interface
}

type UpdateHandler struct {
	cid []int64
	loc locker.Interface
	log logger.Interface
}

func NewUpdateHandler(c UpdateHandlerConfig) *UpdateHandler {
	if len(c.Cid) == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cid must not be empty", c)))
	}
	if c.Loc == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Loc must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	var han *UpdateHandler
	{
		han = &UpdateHandler{
			cid: c.Cid,
			loc: c.Loc,
			log: c.Log,
		}
	}

	return han
}
