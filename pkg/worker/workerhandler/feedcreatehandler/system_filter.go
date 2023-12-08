package feedcreatehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *SystemHandler) Filter(tas *task.Task) bool {
	return tas.Meta.Has(map[string]string{
		objectlabel.EvntObject: "*",
		objectlabel.FeedAction: objectlabel.ActionCreate,
		objectlabel.FeedOrigin: objectlabel.OriginSystem,
	})
}
