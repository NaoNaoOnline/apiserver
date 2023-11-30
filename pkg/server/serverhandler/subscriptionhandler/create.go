package subscriptionhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/subscription"
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
				Recv: x.Public.Recv,
				Unix: inpUni(x.Public.Unix),
				User: userid.FromContext(ctx),
			})
		}
	}

	//
	// Create the given resources.
	//

	var out []*subscriptionstorage.Object
	{
		out, err = h.sub.CreateSubs(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Create background tasks for the created resources.
	//

	{
		_, err = h.sub.CreateWrkr(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
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
