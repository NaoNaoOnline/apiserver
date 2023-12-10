package eventhandler

import (
	"context"
	"strconv"
	"time"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/limiter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *event.SearchI) (*event.SearchO, error) {
	var out []*eventstorage.Object

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	//
	// Search events by ID.
	//

	{
		var eve []objectid.ID
		for _, x := range req.Object {
			if x.Intern != nil && x.Intern.Evnt != "" {
				eve = append(eve, objectid.ID(x.Intern.Evnt))
			}
		}

		if len(eve) != 0 {
			lis, err := h.eve.SearchEvnt(use, eve)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	//
	// Search events by user, created.
	//

	{
		var use []objectid.ID
		for _, x := range req.Object {
			if x.Intern != nil && x.Intern.User != "" {
				use = append(use, objectid.ID(x.Intern.User))
			}
		}

		if len(use) != 0 {
			lis, err := h.eve.SearchUser(use)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	//
	// Search events by label.
	//

	{
		var lab [][]objectid.ID
		for _, x := range req.Object {
			if x.Public != nil && x.Public.Cate != "" {
				lab = append(lab, inpLab(x.Public.Cate))
			}
			if x.Public != nil && x.Public.Host != "" {
				lab = append(lab, inpLab(x.Public.Host))
			}
		}

		for _, x := range lab {
			lis, err := h.eve.SearchLabl(x)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	//
	// Search events by list.
	//

	{
		var lis objectid.ID
		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.List != "" {
				lis = objectid.ID(x.Symbol.List)
			}
		}

		if lis != "" {
			var pag [2]int
			{
				pag = [2]int{
					int(musNum(req.Filter.Paging.Strt)),
					int(musNum(req.Filter.Paging.Stop)),
				}
			}

			// TODO restrict feed length for non-premium users
			eid, err := h.fee.SearchFeed(lis, pag)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			if len(eid) != 0 {
				var eob eventstorage.Slicer
				{
					eob, err = h.eve.SearchEvnt("", eid)
					if err != nil {
						return nil, tracer.Mask(err)
					}
				}

				out = append(out, eob...)
			}
		}
	}

	//
	// Search events by time, happened.
	//

	{
		var hap bool
		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.Time == "hpnd" {
				hap = true
			}
		}

		if hap {
			lis, err := h.eve.SearchHpnd()
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	//
	// Search events by time, pagination.
	//

	{
		var pag bool
		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.Time == "page" {
				pag = true
			}
		}

		if pag {
			min := time.Unix(musNum(req.Filter.Paging.Strt), 0)
			max := time.Unix(musNum(req.Filter.Paging.Stop), 0)

			lis, err := h.eve.SearchTime(min, max)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	//
	// Search events by time, upcoming.
	//

	{
		var upc bool
		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.Time == "upcm" {
				upc = true
			}
		}

		if upc {
			lis, err := h.eve.SearchUpcm()
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	//
	// Search events by likes, pagination.
	//

	{
		var lik string
		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.Like != "" {
				lik = x.Symbol.Like
			}
		}

		if lik != "" {
			var pag [2]int
			{
				pag = [2]int{
					int(musNum(req.Filter.Paging.Strt)),
					int(musNum(req.Filter.Paging.Stop)),
				}
			}

			lis, err := h.eve.SearchLike(objectid.ID(lik), pag)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *event.SearchO
	{
		res = &event.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "event",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
		// Events marked to be deleted cannot be searched anymore.
		if !x.Dltd.IsZero() {
			continue
		}

		res.Object = append(res.Object, &event.SearchO_Object{
			Extern: []*event.SearchO_Object_Extern{
				{
					Amnt: strconv.FormatInt(x.Clck.Data, 10),
					Kind: "link",
					User: x.Clck.User,
				},
			},
			Intern: &event.SearchO_Object_Intern{
				Crtd: outTim(x.Crtd),
				Evnt: x.Evnt.String(),
				User: x.User.String(),
			},
			Public: &event.SearchO_Object_Public{
				Cate: outLab(append(x.Cate, x.Bltn...)),
				Dura: outDur(x.Dura),
				Host: outLab(x.Host),
				Link: x.Link,
				Time: outTim(x.Time),
			},
		})
	}

	return res, nil
}
