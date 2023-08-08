package eventhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Eve eventstorage.Interface
	Log logger.Interface
}

type Handler struct {
	eve eventstorage.Interface
	log logger.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Eve == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Eve must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &Handler{
		eve: c.Eve,
		log: c.Log,
	}
}
