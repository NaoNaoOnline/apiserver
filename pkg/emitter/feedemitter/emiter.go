package feedemitter

import (
	"fmt"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"
)

type EmitterConfig struct {
	Log logger.Interface
	Res rescue.Interface
}

type Emitter struct {
	log logger.Interface
	res rescue.Interface
}

func NewEmitter(c EmitterConfig) *Emitter {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Res == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Res must not be empty", c)))
	}

	return &Emitter{
		log: c.Log,
		res: c.Res,
	}
}
