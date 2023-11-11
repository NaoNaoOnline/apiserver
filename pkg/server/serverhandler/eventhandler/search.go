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

	var evn []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Evnt != "" {
			evn = append(evn, objectid.ID(x.Intern.Evnt))
		}
	}

	if len(evn) != 0 {
		lis, err := h.eve.SearchEvnt(use, evn)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search events by user.
	//

	var usr []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.User != "" {
			usr = append(usr, objectid.ID(x.Intern.User))
		}
	}

	if len(usr) != 0 {
		lis, err := h.eve.SearchUser(usr)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search events by label.
	//

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

	//
	// Search events by list.
	//

	var lid objectid.ID
	for _, x := range req.Object {
		if x.Symbol != nil && x.Symbol.List != "" {
			lid = objectid.ID(x.Symbol.List)
		}
	}

	if lid != "" {
		rul, err := h.rul.SearchList([]objectid.ID{lid})
		if err != nil {
			return nil, tracer.Mask(err)
		}

		eve, err := h.eve.SearchList(rul)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, eve...)
	}

	//
	// Search events by time, happened.
	//

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

	//
	// Search events by time, pagination.
	//

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

	//
	// Search events by time, upcoming.
	//

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

	//
	// Search events by reactions, pagination.
	//

	if userid.FromContext(ctx) != "" {
		var rct bool
		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.Rctn == "page" {
				rct = true
			}
		}

		if rct {
			min := musNum(req.Filter.Paging.Strt)
			max := musNum(req.Filter.Paging.Stop)

			lis, err := h.eve.SearchLike(userid.FromContext(ctx), int(min), int(max))
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
			"message", "search response got truncated",
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
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				Evnt: x.Evnt.String(),
				User: x.User.String(),
			},
			Public: &event.SearchO_Object_Public{
				Cate: outLab(x.Cate),
				Dura: outDur(x.Dura),
				Host: outLab(x.Host),
				Link: x.Link,
				Time: outTim(x.Time),
			},
		})
	}

	return res, nil
}
