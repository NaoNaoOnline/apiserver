package descriptionhandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Ensure(tas *task.Task) error {
	var des objectid.ID
	{
		des = objectid.ID(tas.Meta.Get(objectlabel.DescObject))
	}

	var lim *objectid.Limiter
	{
		lim = objectid.NewLimiter()
	}

	{
		err := h.deleteVote(des, lim)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err := h.deleteDesc(des, lim)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *Handler) deleteDesc(inp objectid.ID, lim *objectid.Limiter) error {
	var err error

	var des []*descriptionstorage.Object
	{
		des, err = h.des.SearchDesc([]objectid.ID{inp})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		_, err := h.des.DeleteDesc(des[:lim.Limit(len(des))])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *Handler) deleteVote(inp objectid.ID, lim *objectid.Limiter) error {
	var err error

	var vot []*votestorage.Object
	{
		vot, err = h.vot.SearchDesc([]objectid.ID{inp})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		_, err := h.vot.Delete(vot[:lim.Limit(len(vot))])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
