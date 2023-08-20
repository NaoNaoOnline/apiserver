package ratinghandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/ratingstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Rat ratingstorage.Interface
}

type Handler struct {
	log logger.Interface
	rat ratingstorage.Interface
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
