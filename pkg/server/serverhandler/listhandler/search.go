package listhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/list"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/isprem"
	"github.com/NaoNaoOnline/apiserver/pkg/server/limiter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *list.SearchI) (*list.SearchO, error) {
	var out []*liststorage.Object

	var pre bool
	{
		pre = isprem.FromContext(ctx)
	}

	//
	// Search lists by user, created.
	//

	{
		var use objectid.ID
		for _, x := range req.Object {
			if x.Intern != nil && x.Intern.User != "" {
				use = objectid.ID(x.Intern.User)
			}
		}

		// Since creating multiple lists is a premium feature it is important to
		// only return the first list for users that created multiple lists in the
		// past as a premium subscriber and have now been downgraded. So users with
		// an expired premium subscription should only see their first list as
		// incentive to upgrade again.
		var pag [2]int
		if !pre {
			pag = liststorage.PagFir()
		} else {
			pag = liststorage.PagAll()
		}

		{
			lis, err := h.lis.SearchUser(use, pag)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *list.SearchO
	{
		res = &list.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "list",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
		// Lists marked to be deleted cannot be searched anymore.
		if !x.Dltd.IsZero() {
			continue
		}

		res.Object = append(res.Object, &list.SearchO_Object{
			Intern: &list.SearchO_Object_Intern{
				Crtd: outTim(x.Crtd),
				Feed: &list.SearchO_Object_Intern_Feed{
					Time: outTim(x.Feed.Time),
				},
				List: x.List.String(),
				User: x.User.String(),
			},
			Public: &list.SearchO_Object_Public{
				Desc: x.Desc.Data,
				Feed: outTim(x.Feed.Data),
			},
		})
	}

	return res, nil
}
