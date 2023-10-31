package eventhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/limiter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *event.SearchI) (*event.SearchO, error) {
	var out []*eventstorage.Object

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
		lis, err := h.eve.SearchEvnt(evn)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search events by user.
	//

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

		lis, err := h.eve.SearchRule(rul)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search events by time.
	//

	var lts bool
	for _, x := range req.Object {
		if x.Symbol != nil && x.Symbol.Ltst == "default" {
			lts = true
		}
	}

	if lts {
		lis, err := h.eve.SearchLtst()
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search events by reactions.
	//

	if userid.FromContext(ctx) != "" {
		var rct bool
		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.Rctn == "default" {
				rct = true
			}
		}

		if rct {
			lis, err := h.eve.SearchLike(userid.FromContext(ctx))
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
