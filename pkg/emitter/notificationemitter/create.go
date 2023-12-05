package notificationemitter

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (e *Emitter) Create(eid objectid.ID, oid objectid.ID, kin string) error {
	var tas *task.Task
	{
		tas = &task.Task{
			Meta: &task.Meta{
				objectlabel.EvntObject: eid.String(),
				objectlabel.NotiAction: objectlabel.ActionCreate,
				objectlabel.NotiOrigin: objectlabel.OriginSystem,
			},
			Sync: &task.Sync{
				task.Paging: "0",
			},
		}
	}

	{
		if kin == "cate" {
			tas.Meta.Set(objectlabel.CateObject, oid.String())
		}

		if kin == "host" {
			tas.Meta.Set(objectlabel.HostObject, oid.String())
		}

		if kin == "user" {
			tas.Meta.Set(objectlabel.UserObject, oid.String())
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
