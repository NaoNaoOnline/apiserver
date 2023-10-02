package descriptionhandler

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *Handler) Filter(tas *task.Task) bool {
	met := map[string]string{
		objectlabel.DescAction: objectlabel.ActionDelete,
		objectlabel.DescObject: "*",
	}

	return tas.Meta.Has(met)
}
