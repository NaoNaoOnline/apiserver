package descriptionhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/limiter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *description.SearchI) (*description.SearchO, error) {
	var out []*descriptionstorage.Object

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	var evn []objectid.ID
	for _, x := range req.Object {
		if x.Public != nil && x.Public.Evnt != "" {
			evn = append(evn, objectid.ID(x.Public.Evnt))
		}
	}

	//
	// Search the given resources.
	//

	{
		lis, err := h.des.SearchEvnt(use, evn)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Construct the RPC response.
	//

	var res *description.SearchO
	{
		res = &description.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "description",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
		// Descriptions marked to be deleted cannot be searched anymore.
		if !x.Dltd.IsZero() {
			continue
		}

		res.Object = append(res.Object, &description.SearchO_Object{
			Extern: []*description.SearchO_Object_Extern{
				{
					Amnt: strconv.FormatInt(x.Mtrc.Data[objectlabel.DescriptionMetricUser], 10),
					Kind: "like",
					User: x.Mtrc.User[objectlabel.DescriptionMetricUser],
				},
			},
			Intern: &description.SearchO_Object_Intern{
				Crtd: outTim(x.Crtd),
				Desc: x.Desc.String(),
				User: x.User.String(),
			},
			Public: &description.SearchO_Object_Public{
				Evnt: x.Evnt.String(),
				Text: x.Text.Data,
			},
		})
	}

	return res, nil
}
