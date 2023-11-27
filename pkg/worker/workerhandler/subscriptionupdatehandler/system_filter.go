package subscriptionupdatehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *UpdateHandler) Filter(tas *task.Task) bool {
	return tas.Meta.Has(map[string]string{
		objectlabel.SubsAction: objectlabel.ActionUpdate,
		objectlabel.SubsObject: "*",
		objectlabel.SubsOrigin: objectlabel.OriginSystem,
	})
}
