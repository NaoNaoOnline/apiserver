package eventcreatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter"
	"github.com/NaoNaoOnline/apiserver/pkg/feed"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/client/twitterclient"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Emi *emitter.Emitter
	Eve eventstorage.Interface
	Fee feed.Interface
	Log logger.Interface
	Twi twitterclient.Interface
}

type SystemHandler struct {
	emi *emitter.Emitter
	eve eventstorage.Interface
	fee feed.Interface
	log logger.Interface
	twi twitterclient.Interface
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
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
	if c.Twi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Twi must not be empty", c)))
	}

	return &SystemHandler{
		emi: c.Emi,
		eve: c.Eve,
		fee: c.Fee,
		log: c.Log,
		twi: c.Twi,
	}
}
