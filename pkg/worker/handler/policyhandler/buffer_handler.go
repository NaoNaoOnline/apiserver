package policyhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type BufferHandlerConfig struct {
	Log logger.Interface
	Prm permission.Interface
}

type BufferHandler struct {
	log logger.Interface
	prm permission.Interface
}

func NewBufferHandler(c BufferHandlerConfig) *BufferHandler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Prm == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Prm must not be empty", c)))
	}

	var han *BufferHandler
	{
		han = &BufferHandler{
			log: c.Log,
			prm: c.Prm,
		}
	}

	return han
}
