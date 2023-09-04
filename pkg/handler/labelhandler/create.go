package labelhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/NaoNaoOnline/apiserver/pkg/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *label.CreateI) (*label.CreateO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	var inp []*labelstorage.Object
	for _, x := range req.Object {
		inp = append(inp, &labelstorage.Object{
			Desc: x.Public.Desc,
			Disc: x.Public.Disc,
			Kind: x.Public.Kind,
			Name: x.Public.Name,
			Twit: x.Public.Twit,
			User: userid.FromContext(ctx),
		})
	}

	var out []*labelstorage.Object
	{
		out, err = h.lab.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *label.CreateO
	{
		res = &label.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &label.CreateO_Object{
			Intern: &label.CreateO_Object_Intern{
				Crtd: strconv.Itoa(int(x.Crtd.Unix())),
				Labl: x.Labl.String(),
			},
		})
	}

	return res, nil
}
