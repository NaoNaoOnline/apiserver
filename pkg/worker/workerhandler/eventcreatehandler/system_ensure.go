package eventcreatehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
)

func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	// TODO emit events for further processing
	return nil
}
