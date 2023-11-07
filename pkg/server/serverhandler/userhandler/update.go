package userhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *user.UpdateI) (*user.UpdateO, error) {
	var err error

	var use []objectid.ID
	var pat [][]*userstorage.Patch
	for _, x := range req.Object {
		if x.Intern != nil && x.Update != nil && x.Intern.User != "" {
			use = append(use, objectid.ID(x.Intern.User))
			pat = append(pat, inpPat(x.Update))
		}
	}

	var inp []*userstorage.Object
	{
		inp, err = h.use.SearchUser(use)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for _, x := range inp {
		if userid.FromContext(ctx) != x.User {
			return nil, tracer.Mask(runtime.UserNotOwnerError)
		}
	}

	var out []objectstate.String
	{
		out, err = h.use.Update(inp, pat)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *user.UpdateO
	{
		res = &user.UpdateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &user.UpdateO_Object{
			Intern: &user.UpdateO_Object_Intern{
				Stts: x.String(),
			},
			Public: &user.UpdateO_Object_Public{},
		})
	}

	return res, nil
}
