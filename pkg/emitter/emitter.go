package emitter

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter/descriptionemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/eventemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/listemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/policyemitter"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Cid []int64
	Cnt []string
	Log logger.Interface
	Res rescue.Interface
	Rpc []string
}

type Emitter struct {
	des descriptionemitter.Interface
	eve eventemitter.Interface
	lis listemitter.Interface
	pol policyemitter.Interface
}

func New(c Config) *Emitter {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Res == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Res must not be empty", c)))
	}

	var e *Emitter
	{
		e = &Emitter{
			des: descriptionemitter.NewEmitter(descriptionemitter.EmitterConfig{Log: c.Log, Res: c.Res}),
			eve: eventemitter.NewEmitter(eventemitter.EmitterConfig{Log: c.Log, Res: c.Res}),
			lis: listemitter.NewEmitter(listemitter.EmitterConfig{Log: c.Log, Res: c.Res}),
			pol: policyemitter.NewEmitter(policyemitter.EmitterConfig{Cid: c.Cid, Cnt: c.Cnt, Log: c.Log, Res: c.Res, Rpc: c.Rpc}),
		}
	}

	return e
}

func (e *Emitter) Desc() descriptionemitter.Interface {
	return e.des
}

func (e *Emitter) Evnt() eventemitter.Interface {
	return e.eve
}

func (e *Emitter) List() listemitter.Interface {
	return e.lis
}

func (e *Emitter) Plcy() policyemitter.Interface {
	return e.pol
}
