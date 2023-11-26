package subscriptionupdatehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
)

func (h *UpdateHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	// TODO mark subscription object in redis as valid
	// TODO delete subscription lock

	return nil
}
