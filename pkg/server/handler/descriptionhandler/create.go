package descriptionhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *description.CreateI) (*description.CreateO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	var inp []*descriptionstorage.Object
	for _, x := range req.Object {
		inp = append(inp, &descriptionstorage.Object{
			Evnt: objectid.ID(x.Public.Evnt),
			Text: x.Public.Text,
			User: userid.FromContext(ctx),
		})
	}

	var out []*descriptionstorage.Object
	{
		out, err = h.des.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *description.CreateO
	{
		res = &description.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &description.CreateO_Object{
			Intern: &description.CreateO_Object_Intern{
				Crtd: strconv.Itoa(int(x.Crtd.Unix())),
				Desc: x.Desc.String(),
			},
		})
	}

	return res, nil
}
