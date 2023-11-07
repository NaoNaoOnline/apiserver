package descriptionhandler

import (
	"context"
	"time"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *description.UpdateI) (*description.UpdateO, error) {
	var err error

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	var out []objectstate.String

	//
	// Update by JSON-Patch
	//

	{
		var des []objectid.ID
		var pat [][]*descriptionstorage.Patch
		for _, x := range req.Object {
			if x.Intern != nil && x.Update != nil && x.Intern.Desc != "" {
				des = append(des, objectid.ID(x.Intern.Desc))
				pat = append(pat, inpPat(x.Update))
			}
		}

		if len(des) != 0 {
			var inp []*descriptionstorage.Object
			{
				inp, err = h.des.SearchDesc(use, des)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			//
			// Verify the given input.
			//

			{
				err = h.updateVrfyPtch(ctx, inp)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			//
			// Update the given resources.
			//

			{
				lis, err := h.des.UpdatePtch(inp, pat)
				if err != nil {
					return nil, tracer.Mask(err)
				}

				out = append(out, lis...)
			}
		}
	}

	//
	// Track external like
	//

	{
		var des []objectid.ID
		var inc []bool
		for _, x := range req.Object {
			if x.Intern != nil && x.Symbol != nil && x.Intern.Desc != "" && (x.Symbol.Like == "add" || x.Symbol.Like == "rem") {
				{
					des = append(des, objectid.ID(x.Intern.Desc))
				}

				if x.Symbol.Like == "add" {
					inc = append(inc, true)
				}

				if x.Symbol.Like == "rem" {
					inc = append(inc, false)
				}
			}
		}

		if len(des) != 0 {
			var inp []*descriptionstorage.Object
			{
				inp, err = h.des.SearchDesc(use, des)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			//
			// Verify the given input.
			//

			{
				err = h.updateVrfyLike(ctx, inp)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			//
			// Update the given resources.
			//

			{
				lis, err := h.des.UpdateLike(use, inp, inc)
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

	var res *description.UpdateO
	{
		res = &description.UpdateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &description.UpdateO_Object{
			Intern: &description.UpdateO_Object_Intern{
				Stts: x.String(),
			},
			Public: &description.UpdateO_Object_Public{},
		})
	}

	return res, nil
}

func (h *Handler) updateVrfyLike(ctx context.Context, inp descriptionstorage.Slicer) error {
	var err error

	var eve []*eventstorage.Object
	{
		eve, err = h.eve.SearchEvnt(inp.Evnt())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range eve {
		// Ensure descriptions cannot be liked if their events have already been
		// deleted.
		if !x.Dltd.IsZero() {
			return tracer.Mask(eventDeletedError)
		}

		// Ensure descriptions cannot be liked if their events have already
		// happened.
		if x.Happnd() {
			return tracer.Mask(eventAlreadyHappenedError)
		}
	}

	return nil
}

func (h *Handler) updateVrfyPtch(ctx context.Context, inp descriptionstorage.Slicer) error {
	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	for _, x := range inp {
		if use != x.User {
			return tracer.Mask(runtime.UserNotOwnerError)
		}

		// Ensure descriptions cannot be updated after 5 minutes of their creation.
		if x.Crtd.Add(5 * time.Minute).Before(time.Now().UTC()) {
			return tracer.Mask(descriptionUpdatePeriodError)
		}
	}

	return nil
}
