package descriptionhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type CustomHandlerConfig struct {
	Des descriptionstorage.Interface
	Log logger.Interface
}

type CustomHandler struct {
	des descriptionstorage.Interface
	log logger.Interface
}

func NewCustomHandler(c CustomHandlerConfig) *CustomHandler {
	if c.Des == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Des must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &CustomHandler{
		des: c.Des,
		log: c.Log,
	}
}
