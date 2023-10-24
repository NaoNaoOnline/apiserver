package wallethandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Wal walletstorage.Interface
}

type Handler struct {
	log logger.Interface
	wal walletstorage.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Wal == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Wal must not be empty", c)))
	}

	return &Handler{
		log: c.Log,
		wal: c.Wal,
	}
}
