package policyemitter

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (e *Emitter) Buffer() error {
	var tas *task.Task
	{
		tas = &task.Task{
			Meta: &task.Meta{
				objectlabel.PlcyAction: objectlabel.ActionBuffer,
				objectlabel.PlcyOrigin: objectlabel.OriginSystem,
			},
			Node: &task.Node{
				task.Method: task.MthdAll,
			},
		}
	}

	{
		err := e.res.Create(tas)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
