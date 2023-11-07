package userhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/subjectclaim"
	"github.com/NaoNaoOnline/apiserver/pkg/server/limiter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *user.SearchI) (*user.SearchO, error) {
	var out []*userstorage.Object

	//
	// Search users by ID.
	//

	var use []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.User != "" {
			use = append(use, objectid.ID(x.Intern.User))
		}
	}

	if len(use) != 0 {
		lis, err := h.use.SearchUser(use)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search users by name.
	//

	var nam []string
	for _, x := range req.Object {
		if x.Public != nil && x.Public.Name != "" {
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
	// Search users by subject claim.
	//

	var slf bool
	for _, x := range req.Object {
		if x.Symbol != nil && x.Symbol.User == "self" {
			slf = true
		}
	}

	if slf {
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
	}

	var res *user.SearchO
	{
		res = &user.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response got truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "user",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
		// Users marked to be deleted cannot be searched anymore.
		if !x.Dltd.IsZero() {
			continue
		}

		res.Object = append(res.Object, &user.SearchO_Object{
			Intern: &user.SearchO_Object_Intern{
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				User: x.User.String(),
			},
			Public: &user.SearchO_Object_Public{
				Home: x.Home,
				Imag: x.Imag,
				Name: x.Name,
			},
		})
	}

	return res, nil
}
