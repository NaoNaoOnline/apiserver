package listdeletehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *CustomHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var lid objectid.ID
	{
		lid = objectid.ID(tas.Meta.Get(objectlabel.ListObject))
	}

	{
		err := h.deleteRule(lid, bud)
		if err != nil {
			return tracer.Mask(err)
		}

		if bud.Break() {
			return nil
		}
	}

	{
		err := h.deleteList(lid, bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *CustomHandler) deleteList(lid objectid.ID, bud *budget.Budget) error {
	var err error

	var lob []*liststorage.Object
	{
		lob, err = h.lis.SearchList([]objectid.ID{lid})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Delete the aggregated feed.
	{
		err = h.fee.DeleteFeed(lid)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Delete the list object from list storage.
	{
		_, err := h.lis.DeleteList(lob[:bud.Claim(len(lob))])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *CustomHandler) deleteRule(lid objectid.ID, bud *budget.Budget) error {
	var err error

	var rob []*rulestorage.Object
	{
		rob, err = h.rul.SearchList([]objectid.ID{lid}, rulestorage.PagAll())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Delete all necessary cross-references between the deleted rules and all the
	// events they described.
	for _, x := range rob {
		err = h.fee.DeleteRule(x)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Only delete the rule objects from rule storage once the feed references are
	// cleaned up.
	{
		_, err := h.rul.DeleteRule(rob[:bud.Claim(len(rob))])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
