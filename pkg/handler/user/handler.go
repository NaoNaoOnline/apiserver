package user

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/user"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Use user.Interface
}

type Handler struct {
	log logger.Interface
	use user.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Use == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Use must not be empty", c)))
	}

	return &Handler{
		log: c.Log,
		use: c.Use,
	}
}
