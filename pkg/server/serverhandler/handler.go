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
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/reactionhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/rulehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/userhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/votehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler/wallethandler"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Emi *emitter.Emitter
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
	rea *reactionhandler.Handler
	rul *rulehandler.Handler
	use *userhandler.Handler
	vot *votehandler.Handler
	wal *wallethandler.Handler
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

	var h *Handler
	{
		h = &Handler{
			des: descriptionhandler.NewHandler(descriptionhandler.HandlerConfig{Eve: c.Sto.Evnt(), Des: c.Sto.Desc(), Log: c.Log, Prm: c.Prm}),
			eve: eventhandler.NewHandler(eventhandler.HandlerConfig{Eve: c.Sto.Evnt(), Log: c.Log, Prm: c.Prm, Rul: c.Sto.Rule()}),
			lab: labelhandler.NewHandler(labelhandler.HandlerConfig{Lab: c.Sto.Labl(), Log: c.Log}),
			lis: listhandler.NewHandler(listhandler.HandlerConfig{Lis: c.Sto.List(), Log: c.Log}),
			pol: policyhandler.NewHandler(policyhandler.HandlerConfig{Emi: c.Emi.Plcy(), Log: c.Log, Prm: c.Prm}),
			rea: reactionhandler.NewHandler(reactionhandler.HandlerConfig{Log: c.Log, Rct: c.Sto.Rctn()}),
			rul: rulehandler.NewHandler(rulehandler.HandlerConfig{Lis: c.Sto.List(), Log: c.Log, Rul: c.Sto.Rule()}),
			use: userhandler.NewHandler(userhandler.HandlerConfig{Log: c.Log, Use: c.Sto.User()}),
			vot: votehandler.NewHandler(votehandler.HandlerConfig{Des: c.Sto.Desc(), Eve: c.Sto.Evnt(), Log: c.Log, Vot: c.Sto.Vote()}),
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
		h.rea,
		h.rul,
		h.use,
		h.vot,
		h.wal,
	}
}
