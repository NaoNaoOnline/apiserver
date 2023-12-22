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

	var eob *eventstorage.Object
	{
		eob, err = h.searchEvnt(eid, bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// It might happen that events get created and at the time the task runs the
	// event got already deleted. In such a case we just stop processing here.
	if eob == nil {
		return nil
	}

	{
		err = h.fee.CreateEvnt(eob)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var lid []objectid.ID
	{
		lid, err = h.fee.SearchList(eob.Evnt, feed.PagAll())
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

func (h *SystemHandler) searchEvnt(eid objectid.ID, bud *budget.Budget) (*eventstorage.Object, error) {
	var err error

	var eob []*eventstorage.Object
	{
		eob, err = h.eve.SearchEvnt("", []objectid.ID{eid})
		if eventstorage.IsEventObjectNotFound(err) {
			return nil, nil
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(eob) == 0 {
		return nil, nil
	}

	return eob[0], nil
}
