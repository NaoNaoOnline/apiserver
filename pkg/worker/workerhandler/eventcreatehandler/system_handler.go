package eventcreatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter/eventemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/client/twitterclient"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Emi eventemitter.Interface
	Log logger.Interface
	Twi twitterclient.Interface
}

type SystemHandler struct {
	emi eventemitter.Interface
	log logger.Interface
	twi twitterclient.Interface
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Twi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Twi must not be empty", c)))
	}

	return &SystemHandler{
		emi: c.Emi,
		log: c.Log,
		twi: c.Twi,
	}
}
