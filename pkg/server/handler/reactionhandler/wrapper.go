package reactionhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/reaction"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han reaction.API
}

func (w *wrapper) Create(ctx context.Context, req *reaction.CreateI) (*reaction.CreateO, error) {
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

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *reaction.DeleteI) (*reaction.DeleteO, error) {
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

	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *reaction.SearchI) (*reaction.SearchO, error) {
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
			if x.Public == nil || x.Public.Kind == "" {
				return nil, tracer.Mask(searchKindEmptyError)
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *reaction.UpdateI) (*reaction.UpdateO, error) {
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

	return w.han.Update(ctx, req)
}
