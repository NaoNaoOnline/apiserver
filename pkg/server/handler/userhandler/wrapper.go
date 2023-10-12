package userhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han user.API
}

func (w *wrapper) Create(ctx context.Context, req *user.CreateI) (*user.CreateO, error) {
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
			if x.Public == nil {
				return nil, tracer.Mask(handler.QueryObjectEmptyError)
			}
		}
	}

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *user.DeleteI) (*user.DeleteO, error) {
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

func (w *wrapper) Search(ctx context.Context, req *user.SearchI) (*user.SearchO, error) {
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
			if x.Intern == nil && x.Public == nil && x.Symbol == nil {
				return nil, tracer.Mask(handler.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && (x.Public != nil || x.Symbol != nil) {
				return nil, tracer.Mask(searchInternConflictError)
			}
			if x.Public != nil && (x.Intern != nil || x.Symbol != nil) {
				return nil, tracer.Mask(searchPublicConflictError)
			}
			if x.Symbol != nil && (x.Intern != nil || x.Public != nil) {
				return nil, tracer.Mask(searchSymbolConflictError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && x.Intern.User == "" {
				return nil, tracer.Mask(searchInternEmptyError)
			}
			if x.Public != nil && x.Public.Name == "" {
				return nil, tracer.Mask(searchPublicEmptyError)
			}
			if x.Symbol != nil && x.Symbol.User == "" {
				return nil, tracer.Mask(searchSymbolEmptyError)
			}
			if x.Symbol != nil && x.Symbol.User != "self" {
				return nil, tracer.Mask(searchSymbolInvalidError)
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *user.UpdateI) (*user.UpdateO, error) {
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

	return w.han.Update(ctx, req)
}
