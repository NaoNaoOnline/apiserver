package policyhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han policy.API
}

func (w *wrapper) Create(ctx context.Context, req *policy.CreateI) (*policy.CreateO, error) {
	if len(req.Object) == 0 {
		return nil, tracer.Mask(queryObjectEmptyError)
	}

	for _, x := range req.Object {
		if x == nil {
			return nil, tracer.Mask(queryObjectEmptyError)
		}
	}

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *policy.DeleteI) (*policy.DeleteO, error) {
	if len(req.Object) == 0 {
		return nil, tracer.Mask(queryObjectEmptyError)
	}

	for _, x := range req.Object {
		if x == nil {
			return nil, tracer.Mask(queryObjectEmptyError)
		}
	}

	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *policy.SearchI) (*policy.SearchO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(queryObjectEmptyError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(queryObjectEmptyError)
			}
		}
	}

	{
		for _, x := range req.Object {
			if x.Symbol == nil || x.Symbol.Ltst == "" {
				return nil, tracer.Mask(searchLtstEmptyError)
			}
			if x.Public != nil && x.Public.Kind != "" && (x.Symbol.Ltst == "default" || x.Symbol.Ltst == "aggregated") {
				return nil, tracer.Mask(searchKindConflictError)
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *policy.UpdateI) (*policy.UpdateO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(queryObjectEmptyError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(queryObjectEmptyError)
			}
		}
	}

	{
		if userid.FromContext(ctx) == "" {
			return nil, tracer.Mask(userIDEmptyError)
		}

		if len(req.Object) != 1 && req.Object[0].Symbol.Sync != "default" {
			return nil, tracer.Mask(updateSyncInvalidError)
		}
	}

	return w.han.Update(ctx, req)
}
