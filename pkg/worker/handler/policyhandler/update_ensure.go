package policyhandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *UpdateHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	{
		err := h.pol.Update()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
