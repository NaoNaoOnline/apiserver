package wallethandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/wallet"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han *Handler
}

func (w *wrapper) Create(ctx context.Context, req *wallet.CreateI) (*wallet.CreateO, error) {
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

	{
		if userid.FromContext(ctx) == "" {
			return nil, tracer.Mask(handler.UserIDEmptyError)
		}
	}

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *wallet.DeleteI) (*wallet.DeleteO, error) {
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
			if x.Intern == nil {
				return nil, tracer.Mask(handler.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && x.Intern.Wllt == "" {
				return nil, tracer.Mask(handler.QueryObjectEmptyError)
			}
		}
	}

	{
		if userid.FromContext(ctx) == "" {
			return nil, tracer.Mask(handler.UserIDEmptyError)
		}
	}

	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *wallet.SearchI) (*wallet.SearchO, error) {
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
			if x.Intern == nil && x.Public == nil {
				return nil, tracer.Mask(handler.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && x.Intern.Wllt == "" {
				return nil, tracer.Mask(searchInternEmptyError)
			}
			if x.Public != nil && x.Public.Kind == "" {
				return nil, tracer.Mask(searchPublicEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && (x.Public != nil) {
				return nil, tracer.Mask(searchInternConflictError)
			}
			if x.Public != nil && (x.Intern != nil) {
				return nil, tracer.Mask(searchPublicConflictError)
			}
		}
	}

	{
		if userid.FromContext(ctx) == "" {
			return nil, tracer.Mask(handler.UserIDEmptyError)
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *wallet.UpdateI) (*wallet.UpdateO, error) {
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
			if x.Intern == nil && x.Public == nil {
				return nil, tracer.Mask(handler.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && x.Intern.Wllt == "" {
				return nil, tracer.Mask(updateInternEmptyError)
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
