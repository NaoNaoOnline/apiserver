package eventhandler

import (
	"context"
	"strconv"
	"time"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/isprem"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/limiter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *event.SearchI) (*event.SearchO, error) {
	var err error

	var out []*eventstorage.Object

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	var pre bool
	{
		pre = isprem.FromContext(ctx)
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
		var lid []objectid.ID
		{
			lid = symbolList(req)
		}

		if len(lid) != 0 {
			var kin string
			{
				kin = pagingKind(req)
			}

			// TODO restrict feed length for non-premium users to 1 week of history

			// kin=page is used to receive the full list of events described by the
			// provided paging range. This response contains all event information
			// with the returned event objects.
			if kin == "" || kin == "page" {
				for _, x := range lid {
					var eid []objectid.ID
					{
						eid, err = h.fee.SearchPage(x, pagingPage(req))
						if err != nil {
							return nil, tracer.Mask(err)
						}
					}

					if len(eid) != 0 {
						var eob []*eventstorage.Object
						{
							eob, err = h.eve.SearchEvnt(use, eid)
							if err != nil {
								return nil, tracer.Mask(err)
							}
						}

						{
							out = append(out, eob...)
						}
					}
				}
			}

			// kin=unix is used to receive the delta of events that a user has not
			// seen yet for any given list. This notification response has a different
			// shape as it does only contain the event ID and respective list ID in
			// the returned event objects. Note that only users with an active premium
			// subscription should be allowed to receive list notifications. So if a
			// client asks for the events a user might have missed, and if that user
			// does not have an active premium subscription, then return an error.
			// Further note that notifications are pull based and entirely managed on
			// the client side. Any client may just work around the restrictions set
			// here.
			if kin == "unix" {
				if pre {
					for _, x := range lid {
						var eid []objectid.ID
						{
							eid, err = h.fee.SearchUnix(x, pagingUnix(req))
							if err != nil {
								return nil, tracer.Mask(err)
							}
						}

						for _, y := range eid {
							out = append(out, &eventstorage.Object{
								Evnt: y,
								List: x,
							})
						}
					}
				} else {
					{
						return nil, tracer.Mask(searchFeedPremiumError)
					}
				}
			}
		}
	}

	//
	// Search events by time, pagination.
	//

	{
		var pag bool
		for _, x := range req.Object {
			if x.Symbol != nil && x.Symbol.Time == "dflt" {
				pag = true
			}
		}

		// TODO check for paging kind unix

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

		// Event object responses change shape when searching for the delta of a
		// feed. These notification responses do only contain event IDs and their
		// respective list IDs. All other responses should return all available
		// information for the requested event objects.
		var eob *event.SearchO_Object
		if x.List != "" {
			eob = &event.SearchO_Object{
				Intern: &event.SearchO_Object_Intern{
					Evnt: x.Evnt.String(),
					List: x.List.String(),
				},
			}
		} else {
			eob = &event.SearchO_Object{
				Extern: []*event.SearchO_Object_Extern{
					{
						Amnt: strconv.FormatInt(x.Mtrc.Data[objectlabel.EventMetricUser], 10),
						Kind: "link",
						User: x.Mtrc.User[objectlabel.EventMetricUser],
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
			}
		}

		{
			res.Object = append(res.Object, eob)
		}
	}

	return res, nil
}
