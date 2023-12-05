package notificationcreatehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *SystemCateHandler) Filter(tas *task.Task) bool {
	return tas.Meta.Has(map[string]string{
		objectlabel.EvntObject: "*",
		objectlabel.NotiAction: objectlabel.ActionCreate,
		objectlabel.NotiOrigin: objectlabel.OriginSystem,
	})
}
