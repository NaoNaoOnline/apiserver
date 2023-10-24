package eventhandler

import (
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

	{
		_, err = h.eve.DeleteWrkr(eve[:bud.Claim(len(eve))])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
