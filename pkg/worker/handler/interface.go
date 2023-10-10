package handler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
)

type Interface interface {
	// Create provides a task template that should be created upon process
	// initiation. The task template should describe scheduled tasks that this
	// handler is going to process. If the handler does not want to bootstrap any
	// task template, nil is returned.
	Create() *task.Task
	// Ensure executes the handler specific business logic in order to complete
	// the given task, if possible.
	Ensure(*task.Task, *budget.Budget) error
	// Filter expresses whether the handler wants to process the given task.
	Filter(*task.Task) bool
}
