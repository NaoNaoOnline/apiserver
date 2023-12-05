package eventemitter

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (e *Emitter) CreateEvnt(eid objectid.ID) error {
	var tas *task.Task
	{
		tas = &task.Task{
			Meta: &task.Meta{
				objectlabel.EvntAction: objectlabel.ActionCreate,
				objectlabel.EvntObject: eid.String(),
				objectlabel.EvntOrigin: objectlabel.OriginSystem,
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

func (e *Emitter) CreateTwtr(eid objectid.ID) error {
	var tas *task.Task
	{
		tas = &task.Task{
			Meta: &task.Meta{
				objectlabel.EvntObject: eid.String(),
				objectlabel.TwtrAction: objectlabel.ActionCreate,
				objectlabel.TwtrOrigin: objectlabel.OriginSystem,
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
