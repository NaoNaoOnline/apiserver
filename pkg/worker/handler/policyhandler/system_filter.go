package policyhandler

import (
	"strconv"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *SystemHandler) Filter(tas *task.Task) bool {
	return tas.Meta.Has(map[string]string{
		objectlabel.PlcyAction: objectlabel.ActionUpdate,
		objectlabel.PlcyNetwrk: strconv.FormatInt(h.cid, 10),
		objectlabel.PlcyObject: "*",
		objectlabel.PlcyOrigin: "*",
	})
}
