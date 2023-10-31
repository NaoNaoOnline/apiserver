package rulehandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/rule"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/limiter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *rule.SearchI) (*rule.SearchO, error) {
	var err error

	var ids []objectid.ID
	for _, x := range req.Object {
		if x.Public != nil && x.Public.List != "" {
			ids = append(ids, objectid.ID(x.Public.List))
		}
	}

	var obj []*liststorage.Object
	{
		obj, err = h.lis.SearchList(ids)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for _, x := range obj {
		if userid.FromContext(ctx) != x.User {
			return nil, tracer.Mask(runtime.UserNotOwnerError)
		}
	}

	var out []*rulestorage.Object
	{
		lis, err := h.rul.SearchList(ids)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Construct the RPC response.
	//

	var res *rule.SearchO
	{
		res = &rule.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response got truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "rule",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
		// Rules marked to be deleted cannot be searched anymore.
		if !x.Dltd.IsZero() {
			continue
		}

		res.Object = append(res.Object, &rule.SearchO_Object{
			Intern: &rule.SearchO_Object_Intern{
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				Rule: x.Rule.String(),
				User: x.User.String(),
			},
			Public: &rule.SearchO_Object_Public{
				Excl: outIDs(x.Excl),
				Incl: outIDs(x.Incl),
				Kind: x.Kind,
				List: x.List.String(),
			},
		})
	}

	return res, nil
}
