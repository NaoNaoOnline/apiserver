package descriptionhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *description.SearchI) (*description.SearchO, error) {
	var evn []objectid.ID
	for _, x := range req.Object {
		if x.Public != nil && x.Public.Evnt != "" {
			evn = append(evn, objectid.ID(x.Public.Evnt))
		}
	}

	var out []*descriptionstorage.Object
	{
		lis, err := h.des.SearchEvnt(evn)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Construct RPC response.
	//

	var res *description.SearchO
	{
		res = &description.SearchO{}
	}

	for _, x := range out {
		// Descriptions marked to be deleted cannot be searched anymore.
		if !x.Dltd.IsZero() {
			continue
		}

		res.Object = append(res.Object, &description.SearchO_Object{
			Intern: &description.SearchO_Object_Intern{
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				Desc: x.Desc.String(),
				User: x.User.String(),
			},
			Public: &description.SearchO_Object_Public{
				Evnt: x.Evnt.String(),
				Text: x.Text,
			},
		})
	}

	return res, nil
}
