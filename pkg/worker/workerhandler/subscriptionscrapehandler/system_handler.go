package subscriptionscrapehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type ScrapeHandlerConfig struct {
	Cid int64
	Cnt string
	Log logger.Interface
	Rpc string
	Sub subscriptionstorage.Interface
}

type ScrapeHandler struct {
	cid int64
	cnt string
	log logger.Interface
	rpc string
	sub subscriptionstorage.Interface
}

func NewScrapeHandler(c ScrapeHandlerConfig) *ScrapeHandler {
	if c.Cid == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cid must not be empty", c)))
	}
	if c.Cnt == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cnt must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Rpc == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rpc must not be empty", c)))
	}
	if c.Sub == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sub must not be empty", c)))
	}

	var han *ScrapeHandler
	{
		han = &ScrapeHandler{
			cid: c.Cid,
			cnt: c.Cnt,
			log: c.Log,
			rpc: c.Rpc,
			sub: c.Sub,
		}
	}

	return han
}
