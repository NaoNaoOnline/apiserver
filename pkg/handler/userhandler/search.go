package userhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
	"github.com/NaoNaoOnline/apiserver/pkg/context/subjectclaim"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *user.SearchI) (*user.SearchO, error) {
	var out []*userstorage.Object

	if len(req.Object) == 0 {
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
		//
		// Validate the RPC integrity.
		//

		for _, x := range req.Object {
			if x.Intern.User != "" && (x.Public.Name != "") {
				return nil, tracer.Mask(searchUserConflictError)
			}
			if x.Public.Name != "" && (x.Intern.User != "") {
				return nil, tracer.Mask(searchNameConflictError)
			}
		}

		//
		// Search users by name.
		//

		var nam []string
		for _, x := range req.Object {
			if x.Public.Name != "" {
				nam = append(nam, x.Public.Name)
			}
		}

		if len(nam) != 0 {
			lis, err := h.use.SearchName(nam)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}

		//
		// Search users by ID.
		//

		var use []objectid.String
		for _, x := range req.Object {
			if x.Intern.User != "" {
				use = append(use, objectid.String(x.Intern.User))
			}
		}

		if len(use) != 0 {
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
