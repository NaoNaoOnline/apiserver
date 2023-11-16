package labelhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han *Handler
}

func (w *wrapper) Create(ctx context.Context, req *label.CreateI) (*label.CreateO, error) {
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
			if x.Public == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
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
			return nil, tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *label.DeleteI) (*label.DeleteO, error) {
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

func (w *wrapper) Search(ctx context.Context, req *label.SearchI) (*label.SearchO, error) {
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
			if x.Intern == nil && x.Public == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && (x.Intern.Labl == "" && x.Intern.User == "") {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
			if x.Intern != nil && (x.Intern.Labl != "" && x.Intern.User != "") {
				return nil, tracer.Mask(searchInternConflictError)
			}
			if x.Public != nil && (x.Public.Name != "" && x.Public.Kind == "") {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Public != nil && x.Public.Kind == "" {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *label.UpdateI) (*label.UpdateO, error) {
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
			if x.Intern == nil && x.Update == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern == nil || x.Update == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && x.Intern.Labl == "" {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
			if x.Update != nil && len(x.Update) == 0 {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			for _, y := range x.Update {
				if y == nil {
					return nil, tracer.Mask(runtime.QueryObjectEmptyError)
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
