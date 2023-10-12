package labelhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *label.SearchI) (*label.SearchO, error) {
	var err error

	var kin []string
	for _, x := range req.Object {
		kin = append(kin, x.Public.Kind)
	}

	var out []*labelstorage.Object
	{
		out, err = h.lab.SearchKind(generic.Uni(kin))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct RPC response.
	//

	var res *label.SearchO
	{
		res = &label.SearchO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &label.SearchO_Object{
			Intern: &label.SearchO_Object_Intern{
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				Labl: x.Labl.String(),
				User: x.User.String(),
			},
			Public: &label.SearchO_Object_Public{
				Desc: x.Desc,
				Disc: x.Disc,
				Kind: x.Kind,
				Name: x.Name,
				Twit: x.Twit,
			},
		})
	}

	return res, nil
}
