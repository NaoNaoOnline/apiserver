package descriptionhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *description.CreateI) (*description.CreateO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	var inp *descriptionstorage.Object
	{
		inp = &descriptionstorage.Object{
			Evnt: objectid.String(req.Object[0].Public.Evnt),
			Text: req.Object[0].Public.Text,
			User: userid.FromContext(ctx),
		}
	}

	var out *descriptionstorage.Object
	{
		out, err = h.des.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *description.CreateO
	{
		res = &description.CreateO{
			Object: []*description.CreateO_Object{
				{
					Intern: &description.CreateO_Object_Intern{
						Crtd: strconv.Itoa(int(out.Crtd.Unix())),
						Desc: out.Desc.String(),
					},
				},
			},
		}
	}

	return res, nil
}
