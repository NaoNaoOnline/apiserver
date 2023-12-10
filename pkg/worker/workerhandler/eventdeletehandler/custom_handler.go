package eventdeletehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/feed"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type CustomHandlerConfig struct {
	Des descriptionstorage.Interface
	Eve eventstorage.Interface
	Fee feed.Interface
	Log logger.Interface
}

type CustomHandler struct {
	des descriptionstorage.Interface
	eve eventstorage.Interface
	fee feed.Interface
	log logger.Interface
}

func NewCustomHandler(c CustomHandlerConfig) *CustomHandler {
	if c.Des == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Des must not be empty", c)))
	}
	if c.Eve == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Eve must not be empty", c)))
	}
	if c.Fee == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Fee must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &CustomHandler{
		des: c.Des,
		eve: c.Eve,
		fee: c.Fee,
		log: c.Log,
	}
}
