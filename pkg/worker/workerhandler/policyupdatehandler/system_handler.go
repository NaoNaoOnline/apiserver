package policyupdatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter/policyemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type UpdateHandlerConfig struct {
	Cid []int64
	Emi policyemitter.Interface
	Log logger.Interface
	Prm permission.Interface
}

type UpdateHandler struct {
	cid []int64
	emi policyemitter.Interface
	log logger.Interface
	prm permission.Interface
}

func NewUpdateHandler(c UpdateHandlerConfig) *UpdateHandler {
	if len(c.Cid) == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cid must not be empty", c)))
	}
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Prm == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Prm must not be empty", c)))
	}

	var han *UpdateHandler
	{
		han = &UpdateHandler{
			cid: c.Cid,
			emi: c.Emi,
			log: c.Log,
			prm: c.Prm,
		}
	}

	return han
}
