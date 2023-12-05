package eventcreatehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

// Ensure receives a task for every event being created on the NaoNao platform
// and simply fans out more detailed tasks without any further pre-processing.
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

	if h.twi.Verify() {
		err = h.emi.Evnt().CreateTwtr(eid)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range append(eob[0].Bltn, eob[0].Cate...) {
		err = h.emi.Noti().Create(eid, x, "cate")
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range eob[0].Host {
		err = h.emi.Noti().Create(eid, x, "host")
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.emi.Noti().Create(eid, eob[0].User, "user")
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
