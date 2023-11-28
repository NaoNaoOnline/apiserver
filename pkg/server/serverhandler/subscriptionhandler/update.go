package subscriptionhandler

import (
	"context"
	"fmt"
	"time"

	"github.com/NaoNaoOnline/apigocode/pkg/subscription"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *subscription.UpdateI) (*subscription.UpdateO, error) {
	var err error

	//
	// Emit scrape tasks.
	//

	var sob []*subscriptionstorage.Object
	{
		sob, err = h.sub.SearchRecv([]objectid.ID{userid.FromContext(ctx)}, subscriptionstorage.PagLat())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		err = sob[0].VerifyUnix(subscriptionstorage.VerifyOnce(time.Now().UTC()))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if sob[0].Stts == objectstate.Success {
			return nil, tracer.Mask(updateStatusSuccessError)
		}
	}

	var key string
	{
		key = fmt.Sprintf(objectlabel.SubsLocker, sob[0].Subs.String())
	}

	var cur string
	for _, x := range req.Object {
		if x.Symbol != nil && x.Symbol.Pntr != "" {
			cur = x.Symbol.Pntr
		}
	}

	var des string
	var exi bool
	{
		des, exi, err = h.loc.Exists(key)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var sta objectstate.String
	{
		if cur == "" && exi {
			return nil, tracer.Mask(updateSyncLockError)
		}

		if cur == "" && !exi {
			{
				des, err = h.loc.Create(key)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			{
				err = h.emi.Scrape(sob[0].Subs)
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
