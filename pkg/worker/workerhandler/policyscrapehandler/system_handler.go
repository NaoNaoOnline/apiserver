package policyscrapehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Cid int64
	Cnt string
	Log logger.Interface
	Prm permission.Interface
	Rpc string
}

type SystemHandler struct {
	cid int64
	cnt string
	log logger.Interface
	prm permission.Interface
	rpc string
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Cid == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cid must not be empty", c)))
	}
	if c.Cnt == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cnt must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Prm == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Prm must not be empty", c)))
	}
	if c.Rpc == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rpc must not be empty", c)))
	}

	var han *SystemHandler
	{
		han = &SystemHandler{
			cid: c.Cid,
			cnt: c.Cnt,
			log: c.Log,
			prm: c.Prm,
			rpc: c.Rpc,
		}
	}

	return han
}
