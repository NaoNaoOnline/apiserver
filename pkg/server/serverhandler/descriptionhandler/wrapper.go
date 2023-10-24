package descriptionhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han *Handler
}

func (w *wrapper) Create(ctx context.Context, req *description.CreateI) (*description.CreateO, error) {
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

	{
		for _, x := range req.Object {
			if x.Public == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	{
		if userid.FromContext(ctx) == "" {
			return nil, tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *description.DeleteI) (*description.DeleteO, error) {
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

	{
		for _, x := range req.Object {
			if x.Intern == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && x.Intern.Desc == "" {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	{
		if userid.FromContext(ctx) == "" {
			return nil, tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *description.SearchI) (*description.SearchO, error) {
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

	{
		for _, x := range req.Object {
			if x.Public == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Public != nil && x.Public.Evnt == "" {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *description.UpdateI) (*description.UpdateO, error) {
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

	{
		for _, x := range req.Object {
			if x.Intern == nil && x.Update == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Update != nil && len(x.Update) == 0 {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern == nil && x.Update != nil {
				return nil, tracer.Mask(updateEmptyError)
			}
			if x.Intern != nil && x.Update == nil {
				return nil, tracer.Mask(updateEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && x.Intern.Desc == "" {
				return nil, tracer.Mask(updateEmptyError)
			}
		}

		for _, x := range req.Object {
			for _, y := range x.Update {
				if y == nil {
					return nil, tracer.Mask(updateEmptyError)
				}
			}
		}
	}

	{
		if userid.FromContext(ctx) == "" {
			return nil, tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	return w.han.Update(ctx, req)
}
