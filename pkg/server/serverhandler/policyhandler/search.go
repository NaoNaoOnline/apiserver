package policyhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/server/limiter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *policy.SearchI) (*policy.SearchO, error) {
	var out []*policystorage.Object

	//
	// Search policies by aggregation.
	//

	var def bool
	for _, x := range req.Object {
		if x.Symbol != nil && x.Symbol.Ltst == "dflt" {
			def = true
		}
	}

	if def {
		lis, err := h.prm.SearchActv()
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Construct the RPC response.
	//

	var res *policy.SearchO
	{
		res = &policy.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "policy",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
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
