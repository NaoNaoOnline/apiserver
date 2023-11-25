package serverhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/descriptionhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/eventhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/labelhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/listhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/policyhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/rulehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/userhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/wallethandler"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Emi *emitter.Emitter
	Loc locker.Interface
	Log logger.Interface
	Prm permission.Interface
	Sto *storage.Storage
}

type Handler struct {
	des *descriptionhandler.Handler
	eve *eventhandler.Handler
	lab *labelhandler.Handler
	lis *listhandler.Handler
	pol *policyhandler.Handler
	rul *rulehandler.Handler
	use *userhandler.Handler
	wal *wallethandler.Handler
}

func New(c Config) *Handler {
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Loc == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Loc must not be empty", c)))
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

	var h *Handler
	{
		h = &Handler{
			des: descriptionhandler.NewHandler(descriptionhandler.HandlerConfig{Eve: c.Sto.Evnt(), Des: c.Sto.Desc(), Log: c.Log, Prm: c.Prm}),
			eve: eventhandler.NewHandler(eventhandler.HandlerConfig{Eve: c.Sto.Evnt(), Log: c.Log, Prm: c.Prm, Rul: c.Sto.Rule()}),
			lab: labelhandler.NewHandler(labelhandler.HandlerConfig{Lab: c.Sto.Labl(), Log: c.Log}),
			lis: listhandler.NewHandler(listhandler.HandlerConfig{Lis: c.Sto.List(), Log: c.Log}),
			pol: policyhandler.NewHandler(policyhandler.HandlerConfig{Emi: c.Emi.Plcy(), Loc: c.Loc, Log: c.Log, Prm: c.Prm}),
			rul: rulehandler.NewHandler(rulehandler.HandlerConfig{Lis: c.Sto.List(), Log: c.Log, Rul: c.Sto.Rule()}),
			use: userhandler.NewHandler(userhandler.HandlerConfig{Log: c.Log, Use: c.Sto.User()}),
			wal: wallethandler.NewHandler(wallethandler.HandlerConfig{Log: c.Log, Wal: c.Sto.Wllt()}),
		}
	}

	return h
}

func (h *Handler) Hand() []Interface {
	return []Interface{
		h.des,
		h.eve,
		h.lab,
		h.lis,
		h.pol,
		h.rul,
		h.use,
		h.wal,
	}
}
