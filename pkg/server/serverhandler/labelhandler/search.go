package labelhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/limiter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *label.SearchI) (*label.SearchO, error) {
	var out []*labelstorage.Object

	//
	// Search labels by user, created.
	//

	{
		var use objectid.ID
		for _, x := range req.Object {
			if x.Intern != nil && x.Intern.User != "" {
				use = objectid.ID(x.Intern.User)
			}
		}

		{
			lis, err := h.lab.SearchUser(use)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	//
	// Search labels by kind.
	//

	{
		var kin []string
		for _, x := range req.Object {
			if x.Public != nil && x.Public.Kind != "" {
				kin = append(kin, x.Public.Kind)
			}
		}

		{
			lis, err := h.lab.SearchKind(generic.Uni(kin))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *label.SearchO
	{
		res = &label.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response got truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "label",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
		// Labels marked to be deleted cannot be searched anymore.
		if !x.Dltd.IsZero() {
			continue
		}

		res.Object = append(res.Object, &label.SearchO_Object{
			Intern: &label.SearchO_Object_Intern{
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				Labl: x.Labl.String(),
				User: x.User.Data.String(),
			},
			Public: &label.SearchO_Object_Public{
				Desc: x.Desc.Data,
				Kind: x.Kind,
				Name: x.Name.Data,
			},
		})
	}

	return res, nil
}
