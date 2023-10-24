package userhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Use userstorage.Interface
}

type Handler struct {
	log logger.Interface
	use userstorage.Interface
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
