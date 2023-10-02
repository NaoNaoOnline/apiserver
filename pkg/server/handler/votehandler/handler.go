package votehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Eve eventstorage.Interface
	Log logger.Interface
	Vot votestorage.Interface
}

type Handler struct {
	eve eventstorage.Interface
	log logger.Interface
	vot votestorage.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Eve == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Eve must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Vot == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Vot must not be empty", c)))
	}

	return &Handler{
		eve: c.Eve,
		log: c.Log,
		vot: c.Vot,
	}
}
