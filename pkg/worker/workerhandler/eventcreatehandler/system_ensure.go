package eventcreatehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/feed"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var eid objectid.ID
	{
		eid = objectid.ID(tas.Meta.Get(objectlabel.EvntObject))
	}

	var eob []*eventstorage.Object
	{
		eob, err = h.eve.SearchEvnt("", []objectid.ID{eid})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.fee.CreateEvnt(eob[0])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var lid []objectid.ID
	{
		lid, err = h.fee.SearchList(eob[0].Evnt, feed.PagAll())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range lid[:bud.Claim(len(lid))] {
		err = h.fee.CreateFeed(x)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
