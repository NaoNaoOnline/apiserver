package twittercreatehandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *SystemHandler) Filter(tas *task.Task) bool {
	return tas.Meta.Has(map[string]string{
		objectlabel.EvntAction: objectlabel.ActionTicker,
		objectlabel.EvntObject: "*",
		objectlabel.EvntOrigin: objectlabel.OriginSystem,
	})
}
