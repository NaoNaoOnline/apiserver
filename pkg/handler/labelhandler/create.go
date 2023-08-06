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

	var use string
	{
		use = userid.FromContext(ctx)
	}

	if use == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	var inp *labelstorage.Object
	{
		inp = &labelstorage.Object{
			Desc: req.Object[0].Public.Desc,
			Disc: req.Object[0].Public.Disc,
			Kind: req.Object[0].Public.Kind,
			Name: req.Object[0].Public.Name,
			Twit: req.Object[0].Public.Twit,
			User: use,
		}
	}

	var out *labelstorage.Object
	{
		out, err = h.lab.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *label.CreateO
	{
		res = &label.CreateO{
			Object: []*label.CreateO_Object{
				{
					Intern: &label.CreateO_Object_Intern{
						Crtd: strconv.Itoa(int(out.Crtd.Unix())),
						Labl: out.Labl.String(),
					},
				},
			},
		}
	}

	return res, nil
}
