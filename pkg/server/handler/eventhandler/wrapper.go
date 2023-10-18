package eventhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han *Handler
}

func (w *wrapper) Create(ctx context.Context, req *event.CreateI) (*event.CreateO, error) {
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

func (w *wrapper) Delete(ctx context.Context, req *event.DeleteI) (*event.DeleteO, error) {
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
	}

	{
		if userid.FromContext(ctx) == "" {
			return nil, tracer.Mask(handler.UserIDEmptyError)
		}
	}

	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *event.SearchI) (*event.SearchO, error) {
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
			if x.Intern != nil && (x.Intern.Evnt == "" && x.Intern.User == "") {
				return nil, tracer.Mask(searchInternEmptyError)
			}
			if x.Public != nil && (x.Public.Cate == "" && x.Public.Host == "") {
				return nil, tracer.Mask(searchPublicEmptyError)
			}
			if x.Symbol != nil && (x.Symbol.Ltst == "" && x.Symbol.Rctn == "") {
				return nil, tracer.Mask(searchSymbolEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && (x.Public != nil || x.Symbol != nil) {
				return nil, tracer.Mask(queryObjectConflictError)
			}
			if x.Public != nil && (x.Intern != nil || x.Symbol != nil) {
				return nil, tracer.Mask(queryObjectConflictError)
			}
			if x.Symbol != nil && (x.Intern != nil || x.Public != nil) {
				return nil, tracer.Mask(queryObjectConflictError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && (x.Intern.Evnt != "" && x.Intern.User != "") {
				return nil, tracer.Mask(searchInternConflictError)
			}
			if x.Symbol != nil && (x.Symbol.Ltst != "" && x.Symbol.Rctn != "") {
				return nil, tracer.Mask(searchSymbolConflictError)
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *event.UpdateI) (*event.UpdateO, error) {
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
