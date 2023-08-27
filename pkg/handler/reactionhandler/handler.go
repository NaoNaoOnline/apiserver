package reactionhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/reactionstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Rat reactionstorage.Interface
}

type Handler struct {
	log logger.Interface
	rat reactionstorage.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Rat == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rat must not be empty", c)))
	}

	return &Handler{
		log: c.Log,
		rat: c.Rat,
	}
}
