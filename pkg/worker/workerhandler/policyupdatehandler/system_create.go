package policyupdatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

// Create returns a system task template for merging and updating all policy
// records buffered within a sorted set. The workflow here intends to wait for
// all scrape tasks to complete, which are emitted by the ScrapeHandler. Then,
// the task template defined here emits the policy update task that merges all
// buffered policy records. Once the triggered update task completes, it itself
// triggers a task to refresh each local copy of active permission states for
// all workers participating in the network. The broadcasted buffer task
// completes the policy synchronization cycle.
//
// Note that the update task here gets triggered periodically based on the many
// scheduled tasks defining Task.Cron inside the ScrapeHandler's task template.
// Regardless, the update task can also be triggered on demand, which is the
// case for the very first startup sequence of the very first worker in the
// network, and the policy.API/Update RPC handler, which emits all scrape tasks
// upon an authorized request, causing the policy records to be synchronized on
// demand.
func (h *UpdateHandler) Create() *task.Task {
	return &task.Task{
		Gate: cidGat(h.cid),
		Meta: &task.Meta{
			objectlabel.PlcyAction: objectlabel.ActionUpdate,
			objectlabel.PlcyOrigin: objectlabel.OriginSystem,
		},
	}
}

func cidGat(cid []int64) *task.Gate {
	var gat *task.Gate
	{
		gat = &task.Gate{}
	}

	for _, x := range cid {
		gat.Set(fmt.Sprintf(objectlabel.PlcyUnique, x), task.Waiting)
	}

	return gat
}
