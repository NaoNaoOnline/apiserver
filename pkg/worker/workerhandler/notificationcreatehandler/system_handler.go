package notificationcreatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/notificationstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Eve eventstorage.Interface
	Log logger.Interface
	Not notificationstorage.Interface
	Rul rulestorage.Interface
}

type SystemHandler struct {
	eve eventstorage.Interface
	log logger.Interface
	not notificationstorage.Interface
	rul rulestorage.Interface
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Eve == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Eve must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Not == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Not must not be empty", c)))
	}
	if c.Rul == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rul must not be empty", c)))
	}

	return &SystemHandler{
		eve: c.Eve,
		log: c.Log,
		not: c.Not,
		rul: c.Rul,
	}
}
