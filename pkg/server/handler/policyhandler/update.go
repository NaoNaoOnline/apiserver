package policyhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *policy.UpdateI) (*policy.UpdateO, error) {
	//
	// Emit scrape tasks.
	//

	{
		err := h.emi.Scrape()
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct RPC response.
	//

	var res *policy.UpdateO
	{
		res = &policy.UpdateO{
			Object: []*policy.UpdateO_Object{
				{
					Intern: &policy.UpdateO_Object_Intern{
						Stts: objectstate.Started.String(),
					},
				},
			},
		}
	}

	return res, nil
}
