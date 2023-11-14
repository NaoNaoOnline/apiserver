package policyhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han *Handler
}

func (w *wrapper) Create(ctx context.Context, req *policy.CreateI) (*policy.CreateO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
		}

		if len(req.Object) > 100 {
			return nil, tracer.Mask(runtime.QueryObjectLimitError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *policy.DeleteI) (*policy.DeleteO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
		}

		if len(req.Object) > 100 {
			return nil, tracer.Mask(runtime.QueryObjectLimitError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *policy.SearchI) (*policy.SearchO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
		}

		if len(req.Object) > 100 {
			return nil, tracer.Mask(runtime.QueryObjectLimitError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	{
		for _, x := range req.Object {
			if x.Symbol == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.Ltst == "" {
				return nil, tracer.Mask(searchSymbolEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.Ltst != "default" {
				return nil, tracer.Mask(searchLtstInvalidError)
			}
		}
	}

	{
		if userid.FromContext(ctx) == "" {
			return nil, tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	var err error

	var exi bool
	{
		exi, err = w.han.prm.ExistsMemb(userid.FromContext(ctx))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// We want the search result to be empty if the caller is not a policy member.
	// We do not want the call to fail because the overview of policy records is
	// integrated into the policy section on the settings page in the webclient.
	// Every user looking at their settings page will make a request to get the
	// list of available policies. It is ok not to get any result for the majority
	// of users. That is why we do not fail, but instead return an empty list,
	// together with an explanation for the empty response.
	if !exi {
		var res *policy.SearchO
		{
			res = &policy.SearchO{
				Reason: []*policy.SearchO_Reason{
					{
						Desc: runtime.PolicyMemberError.Desc,
						Kind: runtime.PolicyMemberError.Kind,
					},
				},
			}
		}

		return res, nil
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *policy.UpdateI) (*policy.UpdateO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
		}

		if len(req.Object) > 100 {
			return nil, tracer.Mask(runtime.QueryObjectLimitError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	{
		if len(req.Object) != 1 {
			return nil, tracer.Mask(updateSyncInvalidError)
		}

		for _, x := range req.Object {
			if x.Symbol == nil || x.Symbol.Sync != "default" {
				return nil, tracer.Mask(updateSyncInvalidError)
			}
		}
	}

	{
		if userid.FromContext(ctx) == "" {
			return nil, tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	var err error

	var exi bool
	{
		exi, err = w.han.prm.ExistsMemb(userid.FromContext(ctx))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if !exi {
		return nil, tracer.Mask(runtime.PolicyMemberError)
	}

	return w.han.Update(ctx, req)
}
