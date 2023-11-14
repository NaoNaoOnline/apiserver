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

	//
	// Verify the given input.
	//

	{
		err = h.updateVrfyPtch(ctx, inp, pat)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Update the given resources.
	//

	var out []objectstate.String
	{
		out, err = h.use.UpdatePtch(inp, pat)
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

func (h *Handler) updateVrfyPtch(ctx context.Context, inp []*userstorage.Object, pat userstorage.PatchSlicer) error {
	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	for i, x := range inp {
		if use != x.User {
			return tracer.Mask(runtime.UserNotOwnerError)
		}

		// Ensure user names can only be updated once within the past 7 days.
		if pat.RepNam(i) && !x.UpdNam() {
			return tracer.Mask(nameUpdatePeriodError)
		}
	}

	return nil
}
