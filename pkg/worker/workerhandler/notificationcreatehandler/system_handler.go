package notificationcreatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/notificationstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Log logger.Interface
	Not notificationstorage.Interface
}

type SystemHandler struct {
	log logger.Interface
	not notificationstorage.Interface
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Not == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Not must not be empty", c)))
	}

	return &SystemHandler{
		log: c.Log,
		not: c.Not,
	}
}
