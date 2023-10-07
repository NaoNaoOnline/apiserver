package descriptionhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type CustomHandlerConfig struct {
	Des descriptionstorage.Interface
	Log logger.Interface
	Vot votestorage.Interface
}

type CustomHandler struct {
	des descriptionstorage.Interface
	log logger.Interface
	vot votestorage.Interface
}

func NewCustomHandler(c CustomHandlerConfig) *CustomHandler {
	if c.Des == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Des must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Vot == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Vot must not be empty", c)))
	}

	return &CustomHandler{
		des: c.Des,
		log: c.Log,
		vot: c.Vot,
	}
}
