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

	for _, x := range inp {
		var mod bool
		{
			mod, err = h.prm.ExistsAcce(permission.SystemDesc, use, permission.AccessDelete)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		if mod {
			continue
		}

		if use != x.User {
			return nil, tracer.Mask(runtime.UserNotOwnerError)
		}

		// Ensure descriptions cannot be deleted after 5 minutes of their creation.
		if x.Crtd.Add(5 * time.Minute).Before(time.Now().UTC()) {
			return nil, tracer.Mask(descriptionDeletePeriodError)
		}

		var des []*descriptionstorage.Object
		{
			des, err = h.des.SearchEvnt(use, []objectid.ID{x.Evnt})
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Ensure the only description of an event cannot be deleted.
		if len(des) == 1 {
			return nil, tracer.Mask(descriptionRequirementError)
		}

		var eve []*eventstorage.Object
		{
			eve, err = h.eve.SearchEvnt([]objectid.ID{x.Evnt})
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Ensure descriptions cannot be removed from events that have already been
		// deleted.
		if !eve[0].Dltd.IsZero() {
			return nil, tracer.Mask(eventDeletedError)
		}

		// Ensure descriptions cannot be removed from events that have already
		// happened.
		if eve[0].Happnd() {
			return nil, tracer.Mask(eventAlreadyHappenedError)
		}
	}

	var out []objectstate.String
	{
		out, err = h.des.DeleteWrkr(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct RPC response.
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
