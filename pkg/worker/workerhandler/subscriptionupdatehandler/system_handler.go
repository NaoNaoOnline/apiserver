package subscriptionupdatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type UpdateHandlerConfig struct {
	Cid []int64
	Loc locker.Interface
	Log logger.Interface
	Sub subscriptionstorage.Interface
	Use userstorage.Interface
	Wal walletstorage.Interface
}

type UpdateHandler struct {
	cid []int64
	loc locker.Interface
	log logger.Interface
	sub subscriptionstorage.Interface
	use userstorage.Interface
	wal walletstorage.Interface
}

func NewUpdateHandler(c UpdateHandlerConfig) *UpdateHandler {
	if len(c.Cid) == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cid must not be empty", c)))
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
	if c.Use == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Use must not be empty", c)))
	}
	if c.Wal == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Wal must not be empty", c)))
	}

	var han *UpdateHandler
	{
		han = &UpdateHandler{
			cid: c.Cid,
			loc: c.Loc,
			log: c.Log,
			sub: c.Sub,
			use: c.Use,
			wal: c.Wal,
		}
	}

	return han
}
