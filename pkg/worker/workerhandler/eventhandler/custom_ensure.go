package eventhandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
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

			if bud.Break() {
				return nil
			}
		}

		{
			err := h.deleteDesc(x.Desc, bud)
			if err != nil {
				return tracer.Mask(err)
			}

			if bud.Break() {
				return nil
			}
		}
	}

	{
		err := h.deleteRule(eve, bud)
		if err != nil {
			return tracer.Mask(err)
		}

		if bud.Break() {
			return nil
		}
	}

	{
		err := h.deleteLink(eve, bud)
		if err != nil {
			return tracer.Mask(err)
		}

		if bud.Break() {
			return nil
		}
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
		eve, err = h.eve.SearchEvnt("", []objectid.ID{inp})
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

func (h *CustomHandler) deleteRule(inp objectid.ID, bud *budget.Budget) error {
	var err error

	var rid []objectid.ID
	{
		rid, err = h.eve.SearchRule(inp)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// It might very well be that the event we want to delete was not added to any
	// list. In this case we just return here.
	if len(rid) == 0 {
		return nil
	}

	var rul rulestorage.Slicer
	{
		rul, err = h.rul.SearchRule(rid[:bud.Claim(len(rid))])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var del []*rulestorage.Object
	var upd []*rulestorage.Object
	for _, x := range rul {
		// For each rule object, remove the given event and then check whether the
		// rule is empty or not.
		{
			x.RemRes(inp)
		}

		if x.HasRes() {
			upd = append(upd, x)
		} else {
			del = append(del, x)
		}
	}

	// If there are rules that are empty after we cleaned them up, then delete
	// those rule objects from storage.
	if len(del) != 0 {
		_, err = h.rul.Delete(del)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// If there are rules that are not empty after we cleaned them up, then update
	// those rule objects in storage.
	if len(upd) != 0 {
		_, err = h.rul.Update(upd)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Remove the rule elements from the event reference list to ensure that
	// consecutive calls do not start all over from scratch.
	{
		_, err := h.eve.DeleteRule(inp, rul.Rule())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
