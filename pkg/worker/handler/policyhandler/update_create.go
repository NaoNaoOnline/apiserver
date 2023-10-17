package policyhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

// Create returns a system task template for triggering a task that merges and
// updates all policy records buffered within the memory implementation of the
// policy cache. The workflow here intends to wait for all buffer tasks emitted
// by the BufferHandler to complete, and then to emit the update task that
// finalizes the policy synchronization cycle. The update task here gets
// triggered periodically based on the Task.Cron definition inside the
// BufferHandler's task template. Regardless, the update task can also be
// triggered on demand, which is the case for the policy.API/Update RPC handler,
// which emits all buffer tasks upon an authorized request, causing the policy
// records to be synchronized on demand.
func (h *UpdateHandler) Create() *task.Task {
	return &task.Task{
		Meta: &task.Meta{
			objectlabel.PlcyAction: objectlabel.ActionUpdate,
			objectlabel.PlcyOrigin: objectlabel.OriginSystem,
		},
		Gate: cidGat(h.cid),
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
