package labelhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *label.SearchI) (*label.SearchO, error) {
	var err error

	uni := map[string]struct{}{}
	for _, x := range req.Object {
		uni[x.Intern.Kind] = struct{}{}
	}

	var kin []string
	for k := range uni {
		kin = append(kin, k)
	}

	var out []*labelstorage.Object
	{
		out, err = h.lab.Search(kin)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *label.SearchO
	{
		res = &label.SearchO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &label.SearchO_Object{
			Intern: &label.SearchO_Object_Intern{
				Crtd: strconv.Itoa(int(x.Crtd.Unix())),
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
