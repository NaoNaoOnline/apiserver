package policyupdatehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	{
		err := h.prm.UpdateActv()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err := h.emi.Buffer()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
