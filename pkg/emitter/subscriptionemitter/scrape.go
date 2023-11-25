package subscriptionemitter

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (e *Emitter) Scrape() error {
	for i := range e.cid {
		var tas *task.Task
		{
			tas = &task.Task{
				Meta: &task.Meta{
					objectlabel.SubsAction: objectlabel.ActionScrape,
					objectlabel.SubsChanid: fmt.Sprintf("%d", e.cid[i]),
					objectlabel.SubsCntrct: e.cnt[i],
					objectlabel.SubsOrigin: objectlabel.OriginCustom,
					objectlabel.SubsRpcUrl: e.rpc[i],
				},
			}
		}

		{
			err := e.res.Create(tas)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
