package policyhandler

import (
	"fmt"
	"strconv"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

// Create returns a system task template for periodically fetching policy
// records from the smart contract of the configured blockchain, see chain ID.
// The task template defines a unique policy key, based on the configured chain
// ID, to trigger the task defined by UpdateHandler. The workflow here intends
// to buffer all policy records from all chains within the memory implementation
// of the policy cache, and once all chains have been scraped, merge and update
// the policy cache in order for permission states to take full affect as
// defined by the Policy smart contracts on all deployed chains.
func (h *BufferHandler) Create() *task.Task {
	return &task.Task{
		Cron: &task.Cron{
			task.Aevery: "6 hours",
		},
		Meta: &task.Meta{
			objectlabel.PlcyAction: objectlabel.ActionBuffer,
			objectlabel.PlcyNetwrk: strconv.FormatInt(h.cid, 10),
			objectlabel.PlcyOrigin: objectlabel.OriginSystem,
		},
		Gate: &task.Gate{
			fmt.Sprintf(objectlabel.PlcyUnique, h.cid): task.Trigger,
		},
	}
}
