package listhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/list"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
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
