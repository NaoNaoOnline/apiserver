package discordcreatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/worker/client/discordclient"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/template/eventtemplate"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Dis discordclient.Interface
	Log logger.Interface
	Tem *eventtemplate.Template
}

type SystemHandler struct {
	dis discordclient.Interface
	log logger.Interface
	tem *eventtemplate.Template
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Dis == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Dis must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Tem == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Tem must not be empty", c)))
	}

	return &SystemHandler{
		dis: c.Dis,
		log: c.Log,
		tem: c.Tem,
	}
}
