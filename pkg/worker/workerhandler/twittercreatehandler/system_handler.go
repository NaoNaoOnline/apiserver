package twittercreatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/worker/client/twitterclient"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/template/eventtemplate"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Log logger.Interface
	Tem *eventtemplate.Template
	Twi twitterclient.Interface
}

type SystemHandler struct {
	log logger.Interface
	tem *eventtemplate.Template
	twi twitterclient.Interface
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Tem == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Tem must not be empty", c)))
	}
	if c.Twi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Twi must not be empty", c)))
	}

	return &SystemHandler{
		tem: c.Tem,
		log: c.Log,
		twi: c.Twi,
	}
}
