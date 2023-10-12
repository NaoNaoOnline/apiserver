package userhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/subjectclaim"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *user.CreateI) (*user.CreateO, error) {
	var err error

	var inp *userstorage.Object
	{
		inp = &userstorage.Object{
			Subj: []string{subjectclaim.FromContext(ctx)},
			Imag: req.Object[0].Public.Imag,
			Name: req.Object[0].Public.Name,
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
	// Construct RPC response.
	//

	var res *user.CreateO
	{
		res = &user.CreateO{
			Object: []*user.CreateO_Object{
				{
					Intern: &user.CreateO_Object_Intern{
						Crtd: strconv.Itoa(int(out.Crtd.Unix())),
						User: out.User.String(),
					},
				},
			},
		}
	}

	return res, nil
}
