package eventcreatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter/eventemitter"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Emi eventemitter.Interface
	Log logger.Interface
}

type SystemHandler struct {
	emi eventemitter.Interface
	log logger.Interface
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &SystemHandler{
		emi: c.Emi,
		log: c.Log,
	}
}
