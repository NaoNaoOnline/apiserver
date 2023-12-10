package eventhandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

func pagingKind(req *event.SearchI) string {
	return req.Filter.Paging.Kind
}

func pagingPage(req *event.SearchI) [2]int {
	return [2]int{
		int(musNum(req.Filter.Paging.Strt)),
		int(musNum(req.Filter.Paging.Stop)),
	}
}

func pagingUnix(req *event.SearchI) [2]float64 {
	return [2]float64{
		objectid.Paging(musNum(req.Filter.Paging.Strt)).Float(),
		objectid.Paging(musNum(req.Filter.Paging.Stop)).Float(),
	}
}

func symbolList(req *event.SearchI) objectid.ID {
	for _, x := range req.Object {
		if x.Symbol != nil && x.Symbol.List != "" {
			return objectid.ID(x.Symbol.List)
		}
	}

	return ""
}
