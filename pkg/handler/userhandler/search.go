package userhandler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
	"github.com/NaoNaoOnline/apiserver/pkg/context/subjectclaim"
	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *user.SearchI) (*user.SearchO, error) {
	var out []*userstorage.Object

	if len(req.Object) == 0 || (len(req.Object) == 1 && req.Object[0].Intern.User == "") {
		var sub string
		{
			sub = subjectclaim.FromContext(ctx)
		}

		{
			obj, err := h.use.SearchSubj(sub)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, obj)
		}
	} else {
		for _, x := range req.Object {
			if x.Intern.User == "" {
				return nil, tracer.Mask(fmt.Errorf("multiple request objects must all contain user IDs"))
			}
		}

		var use []scoreid.String
		for _, x := range req.Object {
			use = append(use, scoreid.String(x.Intern.User))
		}

		{
			lis, err := h.use.SearchUser(use)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	var res *user.SearchO
	{
		res = &user.SearchO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &user.SearchO_Object{
			Intern: &user.SearchO_Object_Intern{
				Crtd: strconv.Itoa(int(x.Crtd.Unix())),
				User: x.User.String(),
			},
			Public: &user.SearchO_Object_Public{
				Imag: x.Imag,
				Name: x.Name,
			},
		})
	}

	return res, nil
}
