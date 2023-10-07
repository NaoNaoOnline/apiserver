package eventhandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var eve []*eventstorage.Object
	{
		eve, err = h.eve.SearchHpnd()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range eve[:bud.Claim(len(eve))] {
		var tas *task.Task
		{
			tas = &task.Task{
				Meta: &task.Meta{
					objectlabel.EvntAction: objectlabel.ActionDelete,
					objectlabel.EvntObject: x.Evnt.String(),
					objectlabel.EvntOrigin: objectlabel.OriginCustom,
				},
			}
		}

		{
			err = h.res.Create(tas)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
