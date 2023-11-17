package policyscrapehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *ScrapeHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	{
		err := h.prm.ScrapeRcrd(tas, bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
