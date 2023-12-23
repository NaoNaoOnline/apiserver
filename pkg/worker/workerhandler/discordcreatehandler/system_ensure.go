package discordcreatehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/template/eventtemplate"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var eid objectid.ID
	{
		eid = objectid.ID(tas.Meta.Get(objectlabel.EvntObject))
	}

	// Try to generate the template and give the result to the discord client. It
	// might happen that events get created and at the time this task runs the
	// event got already deleted. Certain interruptions might happen. In such a
	// cases we just stop processing here.
	var tem string
	{
		tem, err = h.tem.Create(eid, eventtemplate.KindDiscord)
		if eventtemplate.IsCancel(err) {
			return nil
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.dis.Create(tem)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
