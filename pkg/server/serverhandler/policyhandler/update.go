package policyhandler

import (
	"context"
	"strconv"
	"time"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *policy.UpdateI) (*policy.UpdateO, error) {
	var err error

	//
	// Emit scrape tasks.
	//

	var poi time.Time
	{
		poi, err = h.prm.SearchTime()
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cur string
	for _, x := range req.Object {
		if x.Symbol != nil && x.Symbol.Pntr != "" {
			cur = x.Symbol.Pntr
		}
	}

	var des string
	{
		des = strconv.FormatInt(poi.Unix(), 10)
	}

	var sta objectstate.String
	{
		if cur == "" {
			var exi bool
			{
				exi, err = h.prm.ExistsLock()
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			if exi {
				return nil, tracer.Mask(updateSyncLockError)
			}

			{
				err = h.emi.Scrape()
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			{
				err = h.prm.CreateLock()
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			sta = objectstate.Started
		}

		if cur != "" && cur == des {
			sta = objectstate.Waiting
		}

		if cur != "" && cur != des {
			des = ""
			sta = objectstate.Updated
		}
	}

	//
	// Construct the RPC response.
	//

	var res *policy.UpdateO
	{
		res = &policy.UpdateO{
			Object: []*policy.UpdateO_Object{
				{
					Intern: &policy.UpdateO_Object_Intern{
						Stts: sta.String(),
					},
					Symbol: &policy.UpdateO_Object_Symbol{
						Pntr: des,
					},
				},
			},
		}
	}

	return res, nil
}
