package rulecreatehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

// TODO
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

	{
		err = h.fee.CreateRule(rob[0])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.fee.CreateFeed(rob[0].List)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
