package policyhandler

import (
	"strconv"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *SystemHandler) Create() *task.Task {
	return &task.Task{
		Cron: &task.Cron{
			// TODO set to 6 hours
			task.Aevery: "minute",
		},
		Meta: &task.Meta{
			objectlabel.PlcyAction: objectlabel.ActionUpdate,
			objectlabel.PlcyNetwrk: strconv.FormatInt(h.cid, 10),
			objectlabel.PlcyObject: "*",
			objectlabel.PlcyOrigin: objectlabel.OriginSystem,
		},
		Sync: &task.Sync{
			objectlabel.PlcyLatest: "0,0,0,0",
		},
	}
}
