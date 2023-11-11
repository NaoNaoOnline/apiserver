package eventhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type CustomHandlerConfig struct {
	Des descriptionstorage.Interface
	Eve eventstorage.Interface
	Lis liststorage.Interface
	Log logger.Interface
	Rul rulestorage.Interface
}

type CustomHandler struct {
	des descriptionstorage.Interface
	eve eventstorage.Interface
	lis liststorage.Interface
	log logger.Interface
	rul rulestorage.Interface
}

func NewCustomHandler(c CustomHandlerConfig) *CustomHandler {
	if c.Des == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Des must not be empty", c)))
	}
	if c.Eve == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Eve must not be empty", c)))
	}
	if c.Lis == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Lis must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Rul == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rul must not be empty", c)))
	}

	return &CustomHandler{
		des: c.Des,
		eve: c.Eve,
		lis: c.Lis,
		log: c.Log,
		rul: c.Rul,
	}
}
