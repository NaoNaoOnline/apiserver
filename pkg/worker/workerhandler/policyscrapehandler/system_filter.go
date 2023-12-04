package policyscrapehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *SystemHandler) Filter(tas *task.Task) bool {
	return tas.Meta.Has(map[string]string{
		objectlabel.PlcyAction: objectlabel.ActionScrape,
		objectlabel.PlcyChanid: fmt.Sprintf("%d", h.cid),
		objectlabel.PlcyCntrct: h.cnt,
		objectlabel.PlcyOrigin: "*",
		objectlabel.PlcyRpcUrl: h.rpc,
	})
}
