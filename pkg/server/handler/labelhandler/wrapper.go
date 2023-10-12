package labelhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han label.API
}

func (w *wrapper) Create(ctx context.Context, req *label.CreateI) (*label.CreateO, error) {
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

		// It is super important to validate the label kind for labels that can be
		// created via RPC, since we have natively supported and individually created
		// label kinds. Label kind bltn must not be created via RPC, since the system
		// itself is the only authority to manage those labels internally.
		for _, x := range req.Object {
			if x.Public != nil && x.Public.Kind != "cate" && x.Public.Kind != "host" {
				return nil, tracer.Mask(createKindInvalidError)
			}
		}
	}

	{
		if userid.FromContext(ctx) == "" {
			return nil, tracer.Mask(handler.UserIDEmptyError)
		}
	}

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *label.DeleteI) (*label.DeleteO, error) {
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

func (w *wrapper) Search(ctx context.Context, req *label.SearchI) (*label.SearchO, error) {
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

		for _, x := range req.Object {
			if x.Public != nil && x.Public.Kind == "" {
				return nil, tracer.Mask(handler.QueryObjectEmptyError)
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *label.UpdateI) (*label.UpdateO, error) {
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