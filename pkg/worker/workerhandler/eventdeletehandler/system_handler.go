package eventdeletehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Eve eventstorage.Interface
	Log logger.Interface
}

type SystemHandler struct {
	eve eventstorage.Interface
	log logger.Interface
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Eve == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Eve must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &SystemHandler{
		eve: c.Eve,
		log: c.Log,
	}
}
