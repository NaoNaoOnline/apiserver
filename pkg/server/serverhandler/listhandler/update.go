package listhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/list"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/isprem"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *list.UpdateI) (*list.UpdateO, error) {
	var err error

	var lis []objectid.ID
	var pat [][]*liststorage.Patch
	for _, x := range req.Object {
		if x.Intern != nil && x.Update != nil && x.Intern.List != "" {
			lis = append(lis, objectid.ID(x.Intern.List))
			pat = append(pat, inpPat(x.Update))
		}
	}

	var inp []*liststorage.Object
	{
		inp, err = h.lis.SearchList(lis)
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
		out, err = h.lis.UpdatePtch(inp, pat)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *list.UpdateO
	{
		res = &list.UpdateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &list.UpdateO_Object{
			Intern: &list.UpdateO_Object_Intern{
				Stts: x.String(),
			},
			Public: &list.UpdateO_Object_Public{},
		})
	}

	return res, nil
}

func (h *Handler) updateVrfyPtch(ctx context.Context, inp []*liststorage.Object, pat liststorage.PatchSlicer) error {
	var pre bool
	{
		pre = isprem.FromContext(ctx)
	}

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	for i, x := range inp {
		if use != x.User {
			return tracer.Mask(runtime.UserNotOwnerError)
		}

		// Ensure list notifications can only be updated by users having a premium
		// subscription. So if the given patches define the feed time to be
		// replaced, and if the user does not have a premium subscription, then
		// return an error.
		if pat.RplFee(i) && !pre {
			return tracer.Mask(updateFeedPremiumError)
		}
	}

	return nil
}
