package workerhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/client/twitterclient"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/descriptiondeletehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/eventcreatehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/eventdeletehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/listdeletehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/policybufferhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/policyscrapehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/policyupdatehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/subscriptiondonatehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/subscriptionscrapehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/subscriptionupdatehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler/twittercreatehandler"
	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Cid are the chain IDs for all deployed chains.
	Cid []int64
	Emi *emitter.Emitter
	Loc locker.Interface
	Log logger.Interface
	// Pcn are the policy contract addresses for all deployed chains.
	Pcn []string
	Prm permission.Interface
	// Rpc are the RPC endpoints for all deployed chains.
	Rpc []string
	// Scn are the subscription contract addresses for all deployed chains.
	Scn []string
	Sto *storage.Storage
	Twi twitterclient.Interface
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
	if c.Twi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Twi must not be empty", c)))
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
			Twi: c.Twi,
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
		han = append(han, policybufferhandler.NewSystemHandler(policybufferhandler.SystemHandlerConfig{
			Log: c.Log,
			Prm: c.Prm,
		}))
	}

	for i := range c.Cid {
		han = append(han, policyscrapehandler.NewSystemHandler(policyscrapehandler.SystemHandlerConfig{
			Cid: c.Cid[i],
			Cnt: c.Pcn[i],
			Log: c.Log,
			Prm: c.Prm,
			Rpc: c.Rpc[i],
		}))
	}

	{
		han = append(han, policyupdatehandler.NewSystemHandler(policyupdatehandler.SystemHandlerConfig{
			Cid: c.Cid,
			Emi: c.Emi.Plcy(),
			Log: c.Log,
			Prm: c.Prm,
		}))
	}

	{
		han = append(han, subscriptiondonatehandler.NewSystemHandler(subscriptiondonatehandler.SystemHandlerConfig{
			Eve: c.Sto.Evnt(),
			Log: c.Log,
			Sub: c.Sto.Subs(),
			Use: c.Sto.User(),
		}))
	}

	for i := range c.Cid {
		han = append(han, subscriptionscrapehandler.NewSystemHandler(subscriptionscrapehandler.SystemHandlerConfig{
			Cid: c.Cid[i],
			Cnt: c.Scn[i],
			Log: c.Log,
			Rpc: c.Rpc[i],
			Sub: c.Sto.Subs(),
		}))
	}

	{
		han = append(han, subscriptionupdatehandler.NewSystemHandler(subscriptionupdatehandler.SystemHandlerConfig{
			Cid: c.Cid,
			Loc: c.Loc,
			Log: c.Log,
			Sub: c.Sto.Subs(),
			Use: c.Sto.User(),
			Wal: c.Sto.Wllt(),
		}))
	}

	{
		han = append(han, twittercreatehandler.NewSystemHandler(twittercreatehandler.SystemHandlerConfig{
			Des: c.Sto.Desc(),
			Eve: c.Sto.Evnt(),
			Lab: c.Sto.Labl(),
			Log: c.Log,
			Twi: c.Twi,
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
