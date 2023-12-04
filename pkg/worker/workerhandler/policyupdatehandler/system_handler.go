package policyupdatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter/policyemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Cid []int64
	Emi policyemitter.Interface
	Log logger.Interface
	Prm permission.Interface
}

type SystemHandler struct {
	cid []int64
	emi policyemitter.Interface
	log logger.Interface
	prm permission.Interface
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
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

	var han *SystemHandler
	{
		han = &SystemHandler{
			cid: c.Cid,
			emi: c.Emi,
			log: c.Log,
			prm: c.Prm,
		}
	}

	return han
}
