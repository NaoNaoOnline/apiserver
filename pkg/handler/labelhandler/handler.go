package labelhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Lab labelstorage.Interface
}

type Handler struct {
	log logger.Interface
	lab labelstorage.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Lab == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Lab must not be empty", c)))
	}

	return &Handler{
		log: c.Log,
		lab: c.Lab,
	}
}
