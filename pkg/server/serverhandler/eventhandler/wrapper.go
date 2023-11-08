package eventhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han *Handler
}

func (w *wrapper) Create(ctx context.Context, req *event.CreateI) (*event.CreateO, error) {
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
	}

	{
		if userid.FromContext(ctx) == "" {
			return nil, tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *event.DeleteI) (*event.DeleteO, error) {
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
			if x.Intern == nil {
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

func (w *wrapper) Search(ctx context.Context, req *event.SearchI) (*event.SearchO, error) {
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
			if x.Intern == nil && x.Public == nil && x.Symbol == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && (x.Intern.Evnt == "" && x.Intern.User == "") {
				return nil, tracer.Mask(searchInternEmptyError)
			}
			if x.Public != nil && (x.Public.Cate == "" && x.Public.Host == "") {
				return nil, tracer.Mask(searchPublicEmptyError)
			}
			if x.Symbol != nil && (x.Symbol.List == "" && x.Symbol.Rctn == "" && x.Symbol.Time == "") {
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
			if x.Symbol != nil && (x.Symbol.List != "" && x.Symbol.Rctn != "") {
				return nil, tracer.Mask(searchSymbolConflictError)
			}
			if x.Symbol != nil && (x.Symbol.Time != "" && x.Symbol.Rctn != "") {
				return nil, tracer.Mask(searchSymbolConflictError)
			}
			if x.Symbol != nil && (x.Symbol.List != "" && x.Symbol.Time != "") {
				return nil, tracer.Mask(searchSymbolConflictError)
			}
		}

		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.Rctn != "" && x.Symbol.Rctn != "page" {
				return nil, tracer.Mask(searchSymbolRctnError)
			}
			if x.Symbol != nil && x.Symbol.Time != "" && x.Symbol.Time != "hpnd" && x.Symbol.Time != "page" && x.Symbol.Time != "upcm" {
				return nil, tracer.Mask(searchSymbolTimeError)
			}
		}

		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.Rctn == "page" && (req.Filter == nil || req.Filter.Paging == nil || req.Filter.Paging.Strt == "" || req.Filter.Paging.Stop == "" || musNum(req.Filter.Paging.Strt) < 0 || musNum(req.Filter.Paging.Stop) < -1) {
				return nil, tracer.Mask(searchSymbolPageError)
			}
			if x.Symbol != nil && x.Symbol.Time == "page" && (req.Filter == nil || req.Filter.Paging == nil || req.Filter.Paging.Strt == "" || req.Filter.Paging.Stop == "" || musNum(req.Filter.Paging.Strt) <= 0 || musNum(req.Filter.Paging.Stop) <= 0) {
				return nil, tracer.Mask(searchSymbolPageError)
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *event.UpdateI) (*event.UpdateO, error) {
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
			if x.Intern == nil && x.Symbol == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern == nil && x.Symbol != nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
			if x.Intern != nil && x.Symbol == nil {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Intern != nil && x.Intern.Evnt == "" {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
			if x.Symbol != nil && x.Symbol.Link == "" {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}

		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.Link != "add" {
				return nil, tracer.Mask(updateSymbolInvalidError)
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
