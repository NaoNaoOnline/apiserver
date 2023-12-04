package subscriptionscrapehandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *SystemHandler) Filter(tas *task.Task) bool {
	return tas.Meta.Has(map[string]string{
		objectlabel.SubsAction: objectlabel.ActionScrape,
		objectlabel.SubsChanid: fmt.Sprintf("%d", h.cid),
		objectlabel.SubsCntrct: h.cnt,
		objectlabel.SubsOrigin: "*",
		objectlabel.SubsRpcUrl: h.rpc,
	})
}
