package listhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/list"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/isprem"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *list.CreateI) (*list.CreateO, error) {
	var err error

	var inp []*liststorage.Object
	for _, x := range req.Object {
		if x.Public != nil {
			inp = append(inp, &liststorage.Object{
				Desc: objectfield.String{
					Data: x.Public.Desc,
				},
				User: userid.FromContext(ctx),
			})
		}
	}

	//
	// Verify the given input.
	//

	{
		err = h.createVrfy(ctx, inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Create the given resources.
	//

	var out []*liststorage.Object
	{
		out, err = h.lis.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *list.CreateO
	{
		res = &list.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &list.CreateO_Object{
			Intern: &list.CreateO_Object_Intern{
				Crtd: outTim(x.Crtd),
				List: x.List.String(),
			},
		})
	}

	return res, nil
}

func (h *Handler) createVrfy(ctx context.Context, obj liststorage.Slicer) error {
	var err error

	var pre bool
	{
		pre = isprem.FromContext(ctx)
	}

	for _, x := range obj {
		var amn int64
		{
			amn, err = h.lis.SearchAmnt(x.User)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			if !pre && amn >= 1 {
				return tracer.Mask(createListPremiumError)
			}
			if pre && amn >= 50 {
				return tracer.Mask(createListLimitError)
			}
		}
	}

	return nil
}
