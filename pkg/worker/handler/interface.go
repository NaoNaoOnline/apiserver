package handler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
)

type Interface interface {
	// Ensure executes the handler specific business logic in order to complete
	// the given task, if possible.
	Ensure(*task.Task, *budget.Budget) error
	// Filter expresses whether the handler wants to process the given task.
	Filter(*task.Task) bool
}
