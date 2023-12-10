package eventdeletehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/feed"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

// Ensure purges all event related data structures as defined by the provided
// task. If the given budget got exhausted intermittently, then that means there
// are still more data structures to purge and the description itself should not
// be removed just yet. Instead, we stop processing here and continue cleaning
// up on the next execution. If we would just go ahead and remove the actual
// description, then we would leave internal data structures behind that we
// cannot easily find to cleanup, causing systemic state bloat.
func (h *CustomHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var eid objectid.ID
	{
		eid = objectid.ID(tas.Meta.Get(objectlabel.EvntObject))
	}

	var des []*descriptionstorage.Object
	{
		des, err = h.des.SearchEvnt("", []objectid.ID{eid})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range des {
		{
			err = h.deleteLike(x.Desc, bud)
			if err != nil {
				return tracer.Mask(err)
			}

			if bud.Break() {
				return nil
			}
		}

		{
			err = h.deleteDesc(x.Desc, bud)
			if err != nil {
				return tracer.Mask(err)
			}

			if bud.Break() {
				return nil
			}
		}
	}

	{
		err = h.deleteLink(eid, bud)
		if err != nil {
			return tracer.Mask(err)
		}

		if bud.Break() {
			return nil
		}
	}

	var eob []*eventstorage.Object
	{
		eob, err = h.eve.SearchEvnt("", []objectid.ID{eid})
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

	{
		err = h.fee.DeleteEvnt(eob[0])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range lid {
		err = h.fee.CreateFeed(x)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		_, err = h.eve.DeleteEvnt(eob[:bud.Claim(len(eob))])
		if err != nil {
			return tracer.Mask(err)
		}

		if bud.Break() {
			return nil
		}
	}

	return nil
}

func (h *CustomHandler) deleteDesc(inp objectid.ID, bud *budget.Budget) error {
	var err error

	var des []*descriptionstorage.Object
	{
		des, err = h.des.SearchDesc("", []objectid.ID{inp})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		_, err := h.des.DeleteDesc(des[:bud.Claim(len(des))])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *CustomHandler) deleteLike(inp objectid.ID, bud *budget.Budget) error {
	var err error

	var use []objectid.ID
	{
		use, err = h.des.SearchLike(inp)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		_, err := h.des.DeleteLike(inp, use[:bud.Claim(len(use))])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *CustomHandler) deleteLink(inp objectid.ID, bud *budget.Budget) error {
	var err error

	var use []objectid.ID
	{
		use, err = h.eve.SearchLink(inp)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		_, err := h.eve.DeleteLink(inp, use[:bud.Claim(len(use))])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
