package workerhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/descriptionhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/eventhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/listhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/policyhandler"
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
	dch Interface
	ech Interface
	esh Interface
	lch Interface
	pbh Interface
	puh Interface
	psh []Interface
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

	var psh []Interface
	for i := range c.Rpc {
		psh = append(psh, policyhandler.NewScrapeHandler(policyhandler.ScrapeHandlerConfig{
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
			dch: descriptionhandler.NewCustomHandler(descriptionhandler.CustomHandlerConfig{Des: c.Sto.Desc(), Log: c.Log}),
			ech: eventhandler.NewCustomHandler(eventhandler.CustomHandlerConfig{Eve: c.Sto.Evnt(), Des: c.Sto.Desc(), Lis: c.Sto.List(), Log: c.Log, Rul: c.Sto.Rule()}),
			esh: eventhandler.NewSystemHandler(eventhandler.SystemHandlerConfig{Eve: c.Sto.Evnt(), Log: c.Log}),
			lch: listhandler.NewCustomHandler(listhandler.CustomHandlerConfig{Lis: c.Sto.List(), Log: c.Log, Rul: c.Sto.Rule()}),
			pbh: policyhandler.NewBufferHandler(policyhandler.BufferHandlerConfig{Log: c.Log, Prm: c.Prm}),
			puh: policyhandler.NewUpdateHandler(policyhandler.UpdateHandlerConfig{Cid: c.Cid, Emi: c.Emi.Plcy(), Log: c.Log, Prm: c.Prm}),
			psh: psh,
		}
	}

	return h
}

func (h *Handler) Hand() []Interface {
	return append([]Interface{
		h.dch,
		h.ech,
		h.esh,
		h.lch,
		h.pbh,
		h.puh,
	},
		h.psh...,
	)
}
