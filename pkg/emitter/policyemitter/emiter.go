package policyemitter

import (
	"fmt"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"
)

type EmitterConfig struct {
	Cid []int64
	Cnt []string
	Log logger.Interface
	Res rescue.Interface
	Rpc []string
}

type Emitter struct {
	cid []int64
	cnt []string
	log logger.Interface
	res rescue.Interface
	rpc []string
}

func NewEmitter(c EmitterConfig) *Emitter {
	if len(c.Cid) == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cid must not be empty", c)))
	}
	if len(c.Cnt) == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cnt must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Res == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Res must not be empty", c)))
	}
	if len(c.Rpc) == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rpc must not be empty", c)))
	}

	return &Emitter{
		cid: c.Cid,
		cnt: c.Cnt,
		log: c.Log,
		res: c.Res,
		rpc: c.Rpc,
	}
}
