package subscriptionhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/subscription"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *subscription.CreateI) (*subscription.CreateO, error) {
	var err error

	var inp []*subscriptionstorage.Object
	for _, x := range req.Object {
		if x.Public != nil {
			inp = append(inp, &subscriptionstorage.Object{
				Crtr: inpCrt(x.Public.Crtr),
				Payr: objectid.ID(x.Public.Payr),
				Rcvr: objectid.ID(x.Public.Rcvr),
				Unix: inpUni(x.Public.Unix),
				User: userid.FromContext(ctx),
			})
		}
	}

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

	// Create the given resources and create the background tasks for the created
	// resources, if they do not already exist. The logic below makes the
	// subscription.API/Create endpoint idempotent. We want this particular
	// behaviour for the creation process of subscriptions because there are
	// multiple onchain and offchain steps the user has to go through. In case of
	// failure and retry, an already existing subscription object should not
	// produce an error. Further, we do not want to create multiple subscription
	// objects for the same month.

	var out []*subscriptionstorage.Object
	if sub == nil {
		out, err = h.sub.CreateSubs(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	} else {
		out = append(out, sub)
	}

	//
	// Construct the RPC response.
	//

	var res *subscription.CreateO
	{
		res = &subscription.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &subscription.CreateO_Object{
			Intern: &subscription.CreateO_Object_Intern{
				Crtd: outTim(x.Crtd),
				Subs: x.Subs.String(),
			},
		})
	}

	return res, nil
}
