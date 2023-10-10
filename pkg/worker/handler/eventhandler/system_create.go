package eventhandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *SystemHandler) Create() *task.Task {
	return &task.Task{
		Cron: &task.Cron{
			task.Aevery: "5 minutes",
		},
		Meta: &task.Meta{
			objectlabel.EvntAction: objectlabel.ActionDelete,
			objectlabel.EvntObject: "*",
			objectlabel.EvntOrigin: objectlabel.OriginSystem,
		},
	}
}
