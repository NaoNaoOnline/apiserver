package subscriptionhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/subscription"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *subscription.UpdateI) (*subscription.UpdateO, error) {
	var err error

	//
	// Search for existing subscription for the current month.
	//

	var sub *subscriptionstorage.Object
	{
		sub, err = h.sub.SearchCurr(userid.FromContext(ctx))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if sub == nil {
		return nil, tracer.Mask(runtime.ExecutionFailedError)
	}

	var sta objectstate.String
	var des string
	if sub.Stts == objectstate.Success {
		des = ""
		sta = sub.Stts
	} else {
		var key string
		{
			key = fmt.Sprintf(objectlabel.SubsLocker, sub.Subs.String())
		}

		var cur string
		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.Pntr != "" {
				cur = x.Symbol.Pntr
			}
		}

		var exi bool
		{
			des, exi, err = h.loc.Exists(key)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			if cur == "" && exi {
				return nil, tracer.Mask(updateSyncLockError)
			}

			if cur == "" && !exi {
				// Create distributed lock.
				{
					des, err = h.loc.Create(key)
					if err != nil {
						return nil, tracer.Mask(err)
					}
				}

				// Emit scrape tasks.
				{
					err = h.emi.Scrape(sub.Subs)
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
	}

	//
	// Construct the RPC response.
	//

	var res *subscription.UpdateO
	{
		res = &subscription.UpdateO{
			Object: []*subscription.UpdateO_Object{
				{
					Intern: &subscription.UpdateO_Object_Intern{
						Stts: sta.String(),
					},
					Symbol: &subscription.UpdateO_Object_Symbol{
						Pntr: des,
					},
				},
			},
		}
	}

	return res, nil
}
