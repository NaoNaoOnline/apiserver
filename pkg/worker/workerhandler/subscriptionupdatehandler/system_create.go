package subscriptionupdatehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

// Create returns a system task template for finishing the verification cycles
// of specific subscription objects. The workflow here intends to wait for all
// scrape tasks to complete, which are emitted by the ScrapeHandler. Then, the
// task template defined here triggers a task that, given the verification
// process was successful, marks the given subscription to be valid and deletes
// the subscription specific distributed lock.
//
// Note that the update task here is triggered on demand when a subscription is
// created, and optionally, when a user wants to verify their subscription
// status in case the creation process failed to reconcile intermittendly.
func (h *UpdateHandler) Create() *task.Task {
	return &task.Task{
		Gate: cidGat(h.cid),
		Meta: &task.Meta{
			objectlabel.SubsAction: objectlabel.ActionUpdate,
			objectlabel.SubsOrigin: objectlabel.OriginSystem,
		},
		Sync: &task.Sync{
			objectlabel.SubsObject: "n/a",
			objectlabel.SubsVerify: "n/a",
		},
	}
}

func cidGat(cid []int64) *task.Gate {
	var gat *task.Gate
	{
		gat = &task.Gate{}
	}

	for _, x := range cid {
		gat.Set(fmt.Sprintf(objectlabel.SubsUnique, x), task.Waiting)
	}

	return gat
}
