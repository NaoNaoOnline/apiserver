package ruledeletehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *CustomHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var rid objectid.ID
	{
		rid = objectid.ID(tas.Meta.Get(objectlabel.RuleObject))
	}

	var rob []*rulestorage.Object
	{
		rob, err = h.rul.SearchRule([]objectid.ID{rid})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Delete all necessary cross-references between the deleted rule and all the
	// events it described.
	{
		err = h.fee.DeleteRule(rob[0])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Only delete the rule object from rule storage once the feed references are
	// cleaned up.
	{
		_, err := h.rul.DeleteRule(rob)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Generate the associated feed, excluding the deleted rule. Afer this step
	// the associated custom list will show all relevant events.
	{
		err = h.fee.CreateFeed(rob[0].List)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
