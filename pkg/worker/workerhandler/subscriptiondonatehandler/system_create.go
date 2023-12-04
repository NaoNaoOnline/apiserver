package subscriptiondonatehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *SystemHandler) Create() *task.Task {
	return &task.Task{
		Cron: &task.Cron{
			task.Aevery: "hour",
		},
		Meta: &task.Meta{
			objectlabel.SubsAction: objectlabel.ActionDonate,
			objectlabel.SubsOrigin: objectlabel.OriginSystem,
		},
		Sync: &task.Sync{
			objectlabel.SubsPaging: "0",
		},
	}
}
