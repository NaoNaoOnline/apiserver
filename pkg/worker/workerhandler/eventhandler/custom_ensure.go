package eventhandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *CustomHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var eve objectid.ID
	{
		eve = objectid.ID(tas.Meta.Get(objectlabel.EvntObject))
	}

	var des []*descriptionstorage.Object
	{
		des, err = h.des.SearchEvnt("", []objectid.ID{eve})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range des {
		{
			err := h.deleteLike(x.Desc, bud)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// If the given budget got exhausted, then that means there are still more
		// data structures to purge and the description itself should not be removed
		// just yet. Instead, we stop processing here and continue cleaning up on the
		// next execution. If we would just go ahead and remove the actual
		// description, then we would leave internal data structures behind that we
		// cannot easily find to cleanup, causing systemic state bloat.
		if bud.Break() {
			return nil
		}

		{
			err := h.deleteDesc(x.Desc, bud)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	// If the given budget got exhausted, then that means there are still more
	// data structures to purge and the event itself should not be removed just
	// yet. Instead, we stop processing here and continue cleaning up on the next
	// execution. If we would just go ahead and remove the actual event, then we
	// would leave internal data structures behind that we cannot easily find to
	// cleanup, causing systemic state bloat.
	if bud.Break() {
		return nil
	}

	{
		err := h.deleteEvnt(eve, bud)
		if err != nil {
			return tracer.Mask(err)
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

func (h *CustomHandler) deleteEvnt(inp objectid.ID, bud *budget.Budget) error {
	var err error

	var eve []*eventstorage.Object
	{
		eve, err = h.eve.SearchEvnt([]objectid.ID{inp})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		_, err := h.eve.DeleteEvnt(eve[:bud.Claim(len(eve))])
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
