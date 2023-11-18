package twittercreatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/client/twitterclient"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Des descriptionstorage.Interface
	Eve eventstorage.Interface
	Lab labelstorage.Interface
	Log logger.Interface
	Twi twitterclient.Interface
}

type SystemHandler struct {
	des descriptionstorage.Interface
	eve eventstorage.Interface
	lab labelstorage.Interface
	log logger.Interface
	twi twitterclient.Interface
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Des == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Des must not be empty", c)))
	}
	if c.Eve == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Eve must not be empty", c)))
	}
	if c.Lab == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Lab must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Twi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Twi must not be empty", c)))
	}

	return &SystemHandler{
		des: c.Des,
		eve: c.Eve,
		lab: c.Lab,
		log: c.Log,
		twi: c.Twi,
	}
}
