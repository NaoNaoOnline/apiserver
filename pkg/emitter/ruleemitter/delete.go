package ruleemitter

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (e *Emitter) Delete(rid objectid.ID) error {
	var tas *task.Task
	{
		tas = &task.Task{
			Meta: &task.Meta{
				objectlabel.RuleAction: objectlabel.ActionDelete,
				objectlabel.RuleObject: rid.String(),
				objectlabel.RuleOrigin: objectlabel.OriginCustom,
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
