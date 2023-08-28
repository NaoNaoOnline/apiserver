package votehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Vot votestorage.Interface
}

type Handler struct {
	log logger.Interface
	vot votestorage.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Vot == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Vot must not be empty", c)))
	}

	return &Handler{
		log: c.Log,
		vot: c.Vot,
	}
}
