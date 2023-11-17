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
	var lis objectid.ID
	{
		lis = objectid.ID(tas.Meta.Get(objectlabel.ListObject))
	}

	{
		err := h.deleteRule(lis, bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err := h.deleteList(lis, bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *CustomHandler) deleteList(inp objectid.ID, bud *budget.Budget) error {
	var err error

	var lis []*liststorage.Object
	{
		lis, err = h.lis.SearchList([]objectid.ID{inp})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		_, err := h.lis.DeleteList(lis[:bud.Claim(len(lis))])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *CustomHandler) deleteRule(inp objectid.ID, bud *budget.Budget) error {
	var err error

	var rul []*rulestorage.Object
	{
		rul, err = h.rul.SearchList([]objectid.ID{inp})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		_, err := h.rul.Delete(rul[:bud.Claim(len(rul))])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
