package subscriptionhandler

import (
	"context"
	"strconv"
	"strings"

	"github.com/NaoNaoOnline/apigocode/pkg/subscription"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/limiter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *subscription.SearchI) (*subscription.SearchO, error) {
	var out []*subscriptionstorage.Object

	//
	// Search subscriptions by user, payer.
	//

	{
		var pay objectid.ID
		for _, x := range req.Object {
			if x.Public != nil && x.Public.Payr != "" {
				pay = objectid.ID(x.Public.Payr)
			}
		}

		if pay != "" {
			if pay != userid.FromContext(ctx) {
				return nil, tracer.Mask(runtime.UserNotOwnerError)
			}

			lis, err := h.sub.SearchPayr([]objectid.ID{pay}, subscriptionstorage.PagAll())
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	//
	// Search subscriptions by user, receiver.
	//

	{
		var rec objectid.ID
		for _, x := range req.Object {
			if x.Public != nil && x.Public.Rcvr != "" {
				rec = objectid.ID(x.Public.Rcvr)
			}
		}

		if rec != "" {
			if rec != userid.FromContext(ctx) {
				return nil, tracer.Mask(runtime.UserNotOwnerError)
			}

			lis, err := h.sub.SearchRcvr([]objectid.ID{rec}, subscriptionstorage.PagAll())
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *subscription.SearchO
	{
		res = &subscription.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "subscription",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
		res.Object = append(res.Object, &subscription.SearchO_Object{
			Intern: &subscription.SearchO_Object_Intern{
				Crtd: outTim(x.Crtd),
				Fail: x.Fail,
				Stts: x.Stts.String(),
				Subs: x.Subs.String(),
				User: x.User.String(),
			},
			Public: &subscription.SearchO_Object_Public{
				Crtr: strings.Join(x.Crtr, ","),
				Payr: x.Payr.String(),
				Rcvr: x.Rcvr.String(),
				Unix: outTim(x.Unix),
			},
		})
	}

	return res, nil
}
