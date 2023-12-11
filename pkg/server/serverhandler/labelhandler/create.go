package labelhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *label.CreateI) (*label.CreateO, error) {
	var err error

	var inp []*labelstorage.Object
	for _, x := range req.Object {
		if x.Public != nil {
			inp = append(inp, &labelstorage.Object{
				Kind: x.Public.Kind,
				Name: objectfield.String{
					Data: x.Public.Name,
				},
				Prfl: objectfield.MapStr{},
				User: objectfield.ID{
					Data: userid.FromContext(ctx),
				},
			})
		}
	}

	var out []*labelstorage.Object
	{
		out, err = h.lab.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *label.CreateO
	{
		res = &label.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &label.CreateO_Object{
			Intern: &label.CreateO_Object_Intern{
				Crtd: outTim(x.Crtd),
				Labl: x.Labl.String(),
			},
		})
	}

	return res, nil
}
