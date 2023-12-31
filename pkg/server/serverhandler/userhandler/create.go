package userhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/subjectclaim"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *user.CreateI) (*user.CreateO, error) {
	var err error

	var inp *userstorage.Object
	{
		inp = &userstorage.Object{
			Imag: req.Object[0].Public.Imag,
			Name: objectfield.String{
				Data: req.Object[0].Public.Name,
			},
			Prfl: objectfield.MapStr{
				Data: map[string]string{},
			},
			Sclm: []string{subjectclaim.FromContext(ctx)},
		}
	}

	var out *userstorage.Object
	{
		out, err = h.use.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *user.CreateO
	{
		res = &user.CreateO{
			Object: []*user.CreateO_Object{
				{
					Intern: &user.CreateO_Object_Intern{
						Crtd: outTim(out.Crtd),
						User: out.User.String(),
					},
				},
			},
		}
	}

	return res, nil
}
