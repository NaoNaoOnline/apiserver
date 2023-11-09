package descriptionhandler

import (
	"context"
	"time"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, req *description.DeleteI) (*description.DeleteO, error) {
	var err error

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	var des []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Desc != "" {
			des = append(des, objectid.ID(x.Intern.Desc))
		}
	}

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

	var mod bool
	{
		mod, err = h.prm.ExistsAcce(permission.SystemDesc, use, permission.AccessDelete)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if mod {
		// Skip all validity checks for moderators and go straight ahead to
		// deletion.
	} else {
		err = h.deleteVrfy(ctx, inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Delete the given resources.
	//

	var out []objectstate.String
	{
		out, err = h.des.DeleteWrkr(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *description.DeleteO
	{
		res = &description.DeleteO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &description.DeleteO_Object{
			Intern: &description.DeleteO_Object_Intern{
				Stts: x.String(),
			},
			Public: &description.DeleteO_Object_Public{},
		})
	}

	return res, nil
}

func (h *Handler) deleteVrfy(ctx context.Context, inp descriptionstorage.Slicer) error {
	var err error

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	for _, x := range inp {
		if use != x.User {
			return tracer.Mask(runtime.UserNotOwnerError)
		}

		// Ensure descriptions cannot be deleted after 5 minutes of their creation.
		if x.Crtd.Add(5 * time.Minute).Before(time.Now().UTC()) {
			return tracer.Mask(descriptionDeletePeriodError)
		}

		var des []*descriptionstorage.Object
		{
			des, err = h.des.SearchEvnt(use, []objectid.ID{x.Evnt})
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// Ensure the only description of an event cannot be deleted.
		if len(des) == 1 {
			return tracer.Mask(descriptionRequirementError)
		}
	}

	var eve []*eventstorage.Object
	{
		eve, err = h.eve.SearchEvnt("", inp.Evnt())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range eve {
		// Ensure descriptions cannot be removed from events that have already been
		// deleted.
		if !x.Dltd.IsZero() {
			return tracer.Mask(eventDeletedError)
		}

		// Ensure descriptions cannot be removed from events that have already
		// happened.
		if x.Happnd() {
			return tracer.Mask(eventAlreadyHappenedError)
		}
	}

	return nil
}
