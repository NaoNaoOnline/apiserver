package listhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/list"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han *Handler
}

func (w *wrapper) Create(ctx context.Context, req *list.CreateI) (*list.CreateO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *list.DeleteI) (*list.DeleteO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *list.SearchI) (*list.SearchO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *list.UpdateI) (*list.UpdateO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
		}

		for _, x := range req.Object {
			if x == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	return w.han.Update(ctx, req)
}
