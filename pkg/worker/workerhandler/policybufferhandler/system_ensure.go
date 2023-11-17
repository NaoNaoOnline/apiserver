package policybufferhandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *BufferHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	{
		_, err := h.prm.BufferActv()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
