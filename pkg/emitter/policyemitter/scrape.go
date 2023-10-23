package policyemitter

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
					objectlabel.PlcyAction: objectlabel.ActionScrape,
					objectlabel.PlcyChanid: fmt.Sprintf("%d", e.cid[i]),
					objectlabel.PlcyCntrct: e.cnt[i],
					objectlabel.PlcyOrigin: objectlabel.OriginCustom,
					objectlabel.PlcyRpcUrl: e.rpc[i],
				},
				Gate: &task.Gate{
					fmt.Sprintf(objectlabel.PlcyUnique, e.cid[i]): task.Trigger,
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
