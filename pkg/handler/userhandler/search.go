package userhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
	"github.com/NaoNaoOnline/apiserver/pkg/context/subjectclaim"
	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *user.SearchI) (*user.SearchO, error) {
	var err error

	var use scoreid.String
	{
		use = scoreid.String(req.Object[0].Intern.User)
	}

	var sub string
	if use == "" {
		sub = subjectclaim.FromContext(ctx)
	}

	var out *userstorage.Object
	{
		out, err = h.use.Search(sub, use)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *user.SearchO
	{
		res = &user.SearchO{
			Object: []*user.SearchO_Object{
				{
					Intern: &user.SearchO_Object_Intern{
						Crtd: strconv.Itoa(int(out.Crtd.Unix())),
						User: out.User.String(),
					},
					Public: &user.SearchO_Object_Public{
						Imag: out.Imag,
						Name: out.Name,
					},
				},
			},
		}
	}

	return res, nil
}
