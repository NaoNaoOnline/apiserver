package votehandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/vote"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han *Handler
}

func (w *wrapper) Create(ctx context.Context, req *vote.CreateI) (*vote.CreateO, error) {
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

func (w *wrapper) Delete(ctx context.Context, req *vote.DeleteI) (*vote.DeleteO, error) {
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
			if x.Intern == nil || x.Intern.Vote == "" {
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

func (w *wrapper) Search(ctx context.Context, req *vote.SearchI) (*vote.SearchO, error) {
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
			if x.Public != nil && x.Public.Desc == "" {
				return nil, tracer.Mask(votePublicEmptyError)
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *vote.UpdateI) (*vote.UpdateO, error) {
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
