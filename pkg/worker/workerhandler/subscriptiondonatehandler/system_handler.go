package subscriptiondonatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Eve eventstorage.Interface
	Log logger.Interface
	Sub subscriptionstorage.Interface
	Use userstorage.Interface
}

type SystemHandler struct {
	eve eventstorage.Interface
	log logger.Interface
	sub subscriptionstorage.Interface
	use userstorage.Interface
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Eve == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Eve must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Sub == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sub must not be empty", c)))
	}
	if c.Use == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Use must not be empty", c)))
	}

	return &SystemHandler{
		eve: c.Eve,
		log: c.Log,
		sub: c.Sub,
		use: c.Use,
	}
}
