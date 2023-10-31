package eventhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, req *event.DeleteI) (*event.DeleteO, error) {
	var err error

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	var eve []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Evnt != "" {
			eve = append(eve, objectid.ID(x.Intern.Evnt))
		}
	}

	var obj []*eventstorage.Object
	{
		obj, err = h.eve.SearchEvnt(eve)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var mod bool
	{
		mod, err = h.prm.ExistsAcce(permission.SystemEvnt, use, permission.AccessDelete)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for _, x := range obj {
		// Skip all validity checks for moderators and go straight ahead to
		// deletion.
		if mod {
			break
		}

		if use != x.User {
			return nil, tracer.Mask(runtime.UserNotOwnerError)
		}
	}

	var out []objectstate.String
	{
		out, err = h.eve.DeleteWrkr(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *event.DeleteO
	{
		res = &event.DeleteO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &event.DeleteO_Object{
			Intern: &event.DeleteO_Object_Intern{
				Stts: x.String(),
			},
			Public: &event.DeleteO_Object_Public{},
		})
	}

	return res, nil
}
