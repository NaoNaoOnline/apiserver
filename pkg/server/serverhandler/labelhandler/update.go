package labelhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *label.UpdateI) (*label.UpdateO, error) {
	var err error

	var out []objectstate.String

	//
	// Update by JSON-Patch
	//

	{
		var lab []objectid.ID
		var pat [][]*labelstorage.Patch
		for _, x := range req.Object {
			if x.Intern != nil && x.Update != nil && x.Intern.Labl != "" {
				lab = append(lab, objectid.ID(x.Intern.Labl))
				pat = append(pat, inpPat(x.Update))
			}
		}

		if len(lab) != 0 {
			var inp []*labelstorage.Object
			{
				inp, err = h.lab.SearchLabl(lab)
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

			{
				lis, err := h.lab.UpdatePtch(inp, pat)
				if err != nil {
					return nil, tracer.Mask(err)
				}

				out = append(out, lis...)
			}
		}
	}

	//
	// Construct the RPC response.
	//

	var res *label.UpdateO
	{
		res = &label.UpdateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &label.UpdateO_Object{
			Intern: &label.UpdateO_Object_Intern{
				Stts: x.String(),
			},
			Public: &label.UpdateO_Object_Public{},
		})
	}

	return res, nil
}

func (h *Handler) updateVrfyPtch(ctx context.Context, obj labelstorage.Slicer, pat labelstorage.PatchSlicer) error {
	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	for i, x := range obj {
		if use != x.User.Data {
			return tracer.Mask(runtime.UserNotOwnerError)
		}

		for k := range x.Prfl {
			// Ensure label profiles can only be added if they do not already exist.
			// If RemLab returns false, given the existing label y, then it means that
			// a patch defines a label to be added that does not already exist.
			if pat.AddPro(i, k) && !pat.RemPro(i, k) {
				return tracer.Maskf(labelProfileAlreadyExistsError, k)
			}

			// Ensure label profiles can only be removed if they do already exist.
			if !generic.All(obj[i].ProPat(), pat.RemPat(i)) {
				return tracer.Maskf(labelProfileNotFoundError, k)
			}
		}
	}

	return nil
}
