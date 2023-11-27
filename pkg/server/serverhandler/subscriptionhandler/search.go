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
	// Search subscriptions by user, created.
	//

	{
		var use objectid.ID
		for _, x := range req.Object {
			if x.Intern != nil && x.Intern.User != "" {
				use = objectid.ID(x.Intern.User)
			}
		}

		if use != "" {
			if use != userid.FromContext(ctx) {
				return nil, tracer.Mask(runtime.UserNotOwnerError)
			}

			lis, err := h.sub.SearchUser([]objectid.ID{use})
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
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				Fail: outPoi(x.Fail),
				Stts: x.Stts.String(),
				Subs: x.Subs.String(),
				User: x.User.String(),
			},
			Public: &subscription.SearchO_Object_Public{
				Crtr: strings.Join(x.Crtr, ","),
				Sbsc: x.Sbsc,
				Unix: strconv.FormatInt(x.Unix.Unix(), 10),
			},
		})
	}

	return res, nil
}
