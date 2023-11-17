package eventdeletehandler

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

const (
	ninetyDays = time.Hour * 24 * 90
)

func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var now time.Time
	{
		now = time.Now().UTC()
	}

	var min time.Time
	var max time.Time
	{
		min = time.Time{}
		max = now.Add(-ninetyDays)
	}

	var eve []*eventstorage.Object
	{
		eve, err = h.eve.SearchTime(min, max)
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
