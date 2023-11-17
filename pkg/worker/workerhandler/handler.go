package workerhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/descriptiondeletehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/eventcreatehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/eventdeletehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/listdeletehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/policybufferhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/policyscrapehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/policyupdatehandler"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Cid []int64
	Cnt []string
	Emi *emitter.Emitter
	Log logger.Interface
	Prm permission.Interface
	Rpc []string
	Sto *storage.Storage
}

type Handler struct {
	han []Interface
}

func New(c Config) *Handler {
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Prm == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Prm must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	var han []Interface

	{
		han = append(han, descriptiondeletehandler.NewCustomHandler(descriptiondeletehandler.CustomHandlerConfig{
			Des: c.Sto.Desc(),
			Log: c.Log,
		}))
	}

	{
		han = append(han, eventcreatehandler.NewSystemHandler(eventcreatehandler.SystemHandlerConfig{
			Emi: c.Emi.Evnt(),
			Log: c.Log,
		}))
	}

	{
		han = append(han, eventdeletehandler.NewCustomHandler(eventdeletehandler.CustomHandlerConfig{
			Eve: c.Sto.Evnt(),
			Des: c.Sto.Desc(),
			Lis: c.Sto.List(),
			Log: c.Log,
			Rul: c.Sto.Rule(),
		}))
	}

	{
		han = append(han, eventdeletehandler.NewSystemHandler(eventdeletehandler.SystemHandlerConfig{
			Eve: c.Sto.Evnt(),
			Log: c.Log,
		}))
	}

	{
		han = append(han, listdeletehandler.NewCustomHandler(listdeletehandler.CustomHandlerConfig{
			Lis: c.Sto.List(),
			Log: c.Log,
			Rul: c.Sto.Rule(),
		}))
	}

	{
		han = append(han, policybufferhandler.NewBufferHandler(policybufferhandler.BufferHandlerConfig{
			Log: c.Log,
			Prm: c.Prm,
		}))
	}

	{
		han = append(han, policyupdatehandler.NewUpdateHandler(policyupdatehandler.UpdateHandlerConfig{
			Cid: c.Cid,
			Emi: c.Emi.Plcy(),
			Log: c.Log,
			Prm: c.Prm,
		}))
	}

	for i := range c.Rpc {
		han = append(han, policyscrapehandler.NewScrapeHandler(policyscrapehandler.ScrapeHandlerConfig{
			Cid: c.Cid[i],
			Cnt: c.Cnt[i],
			Log: c.Log,
			Prm: c.Prm,
			Rpc: c.Rpc[i],
		}))
	}

	var h *Handler
	{
		h = &Handler{
			han: han,
		}
	}

	return h
}

func (h *Handler) Hand() []Interface {
	return h.han
}
