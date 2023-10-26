package descriptionhandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *CustomHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var des objectid.ID
	{
		des = objectid.ID(tas.Meta.Get(objectlabel.DescObject))
	}

	{
		err := h.deleteVote(des, bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err := h.deleteDesc(des, bud)
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
		des, err = h.des.SearchDesc([]objectid.ID{inp})
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

func (h *CustomHandler) deleteVote(inp objectid.ID, bud *budget.Budget) error {
	var err error

	var vot []*votestorage.Object
	{
		vot, err = h.vot.SearchDesc([]objectid.ID{inp})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		_, err := h.vot.Delete(vot[:bud.Claim(len(vot))])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}