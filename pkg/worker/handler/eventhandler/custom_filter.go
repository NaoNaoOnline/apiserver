package eventhandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *CustomHandler) Filter(tas *task.Task) bool {
	met := map[string]string{
		objectlabel.EvntAction: objectlabel.ActionDelete,
		objectlabel.EvntObject: "*",
		objectlabel.EvntOrigin: objectlabel.OriginCustom,
	}

	return tas.Meta.Has(met)
}
