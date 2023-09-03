package descriptionhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *description.SearchI) (*description.SearchO, error) {
	var evn []objectid.String
	for _, x := range req.Object {
		evn = append(evn, objectid.String(x.Public.Evnt))
	}

	var out []*descriptionstorage.Object
	{
		lis, err := h.des.SearchEvnt(evn)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	var res *description.SearchO
	{
		res = &description.SearchO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &description.SearchO_Object{
			Intern: &description.SearchO_Object_Intern{
				Crtd: strconv.Itoa(int(x.Crtd.Unix())),
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
