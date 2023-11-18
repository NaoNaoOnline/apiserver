package eventcreatehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

// Ensure receives a task for every event being created on the NaoNao platform
// and simply fans out more detailed tasks without any further pre-processing.
func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	if h.twi.Verify() {
		err = h.emi.CreateTwtr(objectid.ID(tas.Meta.Get(objectlabel.EvntObject)))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
