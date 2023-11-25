package subscriptionhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter/subscriptionemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Emi subscriptionemitter.Interface
	Loc locker.Interface
	Log logger.Interface
	Sub subscriptionstorage.Interface
}

type Handler struct {
	emi subscriptionemitter.Interface
	loc locker.Interface
	log logger.Interface
	sub subscriptionstorage.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Loc == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Loc must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Sub == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sub must not be empty", c)))
	}

	return &Handler{
		emi: c.Emi,
		loc: c.Loc,
		log: c.Log,
		sub: c.Sub,
	}
}
