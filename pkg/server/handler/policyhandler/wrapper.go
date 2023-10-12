package policyhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han policy.API
}

func (w *wrapper) Create(ctx context.Context, req *policy.CreateI) (*policy.CreateO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(handler.QueryObjectEmptyError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(handler.QueryObjectEmptyError)
			}
		}
	}

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *policy.DeleteI) (*policy.DeleteO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(handler.QueryObjectEmptyError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(handler.QueryObjectEmptyError)
			}
		}
	}

	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *policy.SearchI) (*policy.SearchO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(handler.QueryObjectEmptyError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(handler.QueryObjectEmptyError)
			}
		}
	}

	{
		for _, x := range req.Object {
			if x.Public == nil && x.Symbol == nil {
				return nil, tracer.Mask(handler.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Public != nil && x.Public.Kind == "" {
				return nil, tracer.Mask(searchPublicEmptyError)
			}
			if x.Symbol != nil && x.Symbol.Ltst == "" {
				return nil, tracer.Mask(searchSymbolEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Public != nil && x.Symbol != nil && x.Public.Kind != "" && (x.Symbol.Ltst == "default" || x.Symbol.Ltst == "aggregated") {
				return nil, tracer.Mask(searchKindConflictError)
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *policy.UpdateI) (*policy.UpdateO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(handler.QueryObjectEmptyError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(handler.QueryObjectEmptyError)
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
			return nil, tracer.Mask(handler.UserIDEmptyError)
		}
	}

	return w.han.Update(ctx, req)
}
