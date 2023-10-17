package policyhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/cache/policycache"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *policy.SearchI) (*policy.SearchO, error) {
	var out []*policycache.Record

	//
	// Search policies by aggregation and delete events.
	//

	var def bool
	for _, x := range req.Object {
		if x.Symbol != nil && x.Symbol.Ltst == "default" {
			def = true
		}
	}

	if def {
		lis, err := h.prm.SearchRcrd()
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Construct RPC response.
	//

	var res *policy.SearchO
	{
		res = &policy.SearchO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &policy.SearchO_Object{
			Extern: outExt(x),
			Intern: &policy.SearchO_Object_Intern{
				User: x.User.String(),
			},
			Public: &policy.SearchO_Object_Public{
				Acce: strconv.FormatInt(x.Acce, 10),
				Memb: x.Memb,
				Syst: strconv.FormatInt(x.Syst, 10),
			},
		})
	}

	return res, nil
}
