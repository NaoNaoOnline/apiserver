package listdeletehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/feed"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type CustomHandlerConfig struct {
	Fee feed.Interface
	Lis liststorage.Interface
	Log logger.Interface
	Rul rulestorage.Interface
}

type CustomHandler struct {
	fee feed.Interface
	lis liststorage.Interface
	log logger.Interface
	rul rulestorage.Interface
}

func NewCustomHandler(c CustomHandlerConfig) *CustomHandler {
	if c.Fee == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Fee must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Lis == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Lis must not be empty", c)))
	}
	if c.Rul == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rul must not be empty", c)))
	}

	return &CustomHandler{
		fee: c.Fee,
		log: c.Log,
		lis: c.Lis,
		rul: c.Rul,
	}
}
