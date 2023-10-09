package policyhandler

import (
	"context"
	"strconv"
	"strings"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *policy.SearchI) (*policy.SearchO, error) {
	//
	// Validate the RPC integrity.
	//

	for _, x := range req.Object {
		if x.Public.Kind != "" && (x.Symbol.Ltst == "default" || x.Symbol.Ltst == "aggregated") {
			return nil, tracer.Mask(searchKindConflictError)
		}
	}

	var out []*policystorage.Object

	//
	// Search policies by aggregation and delete events.
	//

	var def bool
	for _, x := range req.Object {
		if x.Symbol.Ltst == "default" {
			def = true
		}
	}

	if def {
		agg, del, err := h.pol.SearchAggr()
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, agg...)
		out = append(out, del...)
	}

	//
	// Search policys by aggregation only.
	//

	var agg bool
	for _, x := range req.Object {
		if x.Symbol.Ltst == "aggregated" {
			agg = true
		}
	}

	if agg {
		agg, _, err := h.pol.SearchAggr()
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, agg...)
	}

	//
	// Search policys by label.
	//

	var pxy bool
	for _, x := range req.Object {
		if x.Symbol.Ltst == "proxy" {
			pxy = true
		}
	}

	if pxy {
		var kin []string
		for _, x := range req.Object {
			if x.Public.Kind == "" {
				continue
			}

			{
				kin = append(kin, strings.Split(x.Public.Kind, ",")...)
			}
		}

		if len(kin) == 0 {
			kin = []string{"CreateMember", "CreateSystem", "DeleteMember", "DeleteSystem"}
		}

		lis, err := h.pol.SearchKind(kin)
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
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				Plcy: x.Plcy.String(),
			},
			Public: &policy.SearchO_Object_Public{
				Acce: strconv.FormatInt(x.Acce, 10),
				Kind: x.Kind,
				Memb: x.Memb,
				Syst: strconv.FormatInt(x.Syst, 10),
			},
		})
	}

	return res, nil
}
