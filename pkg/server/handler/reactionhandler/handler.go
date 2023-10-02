package reactionhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/reactionstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Rct reactionstorage.Interface
}

type Handler struct {
	log logger.Interface
	rct reactionstorage.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Rct == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rct must not be empty", c)))
	}

	return &Handler{
		log: c.Log,
		rct: c.Rct,
	}
}
