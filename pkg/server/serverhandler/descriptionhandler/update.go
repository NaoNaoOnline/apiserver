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

			for _, x := range inp {
				if use == x.User {
					return nil, tracer.Mask(runtime.UserNotOwnerError)
				}
				// Ensure descriptions cannot be updated after 5 minutes of their creation.
				if x.Crtd.Add(5 * time.Minute).Before(time.Now().UTC()) {
					return nil, tracer.Mask(descriptionUpdatePeriodError)
				}
			}

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
			if x.Intern != nil && x.Symbol != nil && x.Intern.Desc != "" && (x.Symbol.Xtrn == "like" || x.Symbol.Xtrn == "ulik") {
				{
					des = append(des, objectid.ID(x.Intern.Desc))
				}

				if x.Symbol.Xtrn == "like" {
					inc = append(inc, true)
				}

				if x.Symbol.Xtrn == "ulik" {
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

			for _, x := range inp {
				var eve []*eventstorage.Object
				{
					eve, err = h.eve.SearchEvnt([]objectid.ID{x.Evnt})
					if err != nil {
						return nil, tracer.Mask(err)
					}
				}

				if len(eve) != 1 {
					return nil, tracer.Mask(runtime.ExecutionFailedError)
				}

				// Ensure descriptions cannot be liked if their events have already been
				// deleted.
				if !eve[0].Dltd.IsZero() {
					return nil, tracer.Mask(eventDeletedError)
				}

				// Ensure descriptions cannot be liked if their events have already
				// happened.
				if eve[0].Happnd() {
					return nil, tracer.Mask(eventAlreadyHappenedError)
				}
			}

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
	// Construct RPC response.
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
