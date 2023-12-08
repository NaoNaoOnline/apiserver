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
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/subscriptionhandler"
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
	han []Interface
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

	var han []Interface

	{
		han = append(han, descriptionhandler.NewHandler(descriptionhandler.HandlerConfig{
			Eve: c.Sto.Evnt(),
			Des: c.Sto.Desc(),
			Log: c.Log,
			Prm: c.Prm,
		}))
	}

	{
		han = append(han, eventhandler.NewHandler(eventhandler.HandlerConfig{
			Eve: c.Sto.Evnt(),
			Fee: c.Sto.Feed(),
			Log: c.Log,
			Prm: c.Prm,
			Rul: c.Sto.Rule(),
		}))
	}

	{
		han = append(han, labelhandler.NewHandler(labelhandler.HandlerConfig{
			Lab: c.Sto.Labl(),
			Log: c.Log,
		}))
	}

	{
		han = append(han, listhandler.NewHandler(listhandler.HandlerConfig{
			Lis: c.Sto.List(),
			Log: c.Log,
		}))
	}

	{
		han = append(han, policyhandler.NewHandler(policyhandler.HandlerConfig{
			Emi: c.Emi.Plcy(),
			Loc: c.Loc,
			Log: c.Log,
			Prm: c.Prm,
		}))
	}

	{
		han = append(han, rulehandler.NewHandler(rulehandler.HandlerConfig{
			Lis: c.Sto.List(),
			Log: c.Log,
			Rul: c.Sto.Rule(),
		}))
	}

	{
		han = append(han, subscriptionhandler.NewHandler(subscriptionhandler.HandlerConfig{
			Emi: c.Emi.Subs(),
			Loc: c.Loc,
			Log: c.Log,
			Sub: c.Sto.Subs(),
		}))
	}

	{
		han = append(han, userhandler.NewHandler(userhandler.HandlerConfig{
			Log: c.Log,
			Use: c.Sto.User(),
		}))
	}

	{
		han = append(han, wallethandler.NewHandler(wallethandler.HandlerConfig{
			Log: c.Log,
			Sub: c.Sto.Subs(),
			Wal: c.Sto.Wllt(),
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
