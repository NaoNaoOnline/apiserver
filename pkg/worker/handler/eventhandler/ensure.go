package eventhandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var eve objectid.ID
	{
		eve = objectid.ID(tas.Meta.Get(objectlabel.EvntObject))
	}

	var des []*descriptionstorage.Object
	{
		des, err = h.des.SearchEvnt([]objectid.ID{eve})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range des {
		{
			err := h.deleteVote(x.Desc, bud)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err := h.deleteDesc(x.Desc, bud)
			if err != nil {
				return tracer.Mask(err)
			}
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

func (h *Handler) deleteDesc(inp objectid.ID, bud *budget.Budget) error {
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

func (h *Handler) deleteEvnt(inp objectid.ID, bud *budget.Budget) error {
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

func (h *Handler) deleteVote(inp objectid.ID, bud *budget.Budget) error {
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
