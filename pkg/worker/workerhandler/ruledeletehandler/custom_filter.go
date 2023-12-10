package ruledeletehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *CustomHandler) Filter(tas *task.Task) bool {
	return tas.Meta.Has(map[string]string{
		objectlabel.RuleAction: objectlabel.ActionDelete,
		objectlabel.RuleObject: "*",
		objectlabel.RuleOrigin: objectlabel.OriginCustom,
	})
}
