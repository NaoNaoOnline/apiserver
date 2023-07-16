package label

import (
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/pbf/label"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
}

type Handler struct {
	label.UnimplementedAPIServer

	log logger.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}

	return &Handler{
		log: c.Log,
	}
}
