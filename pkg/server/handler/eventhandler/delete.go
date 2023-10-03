package eventhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, req *event.DeleteI) (*event.DeleteO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	var eve []objectid.ID
	for _, x := range req.Object {
		eve = append(eve, objectid.ID(x.Intern.Evnt))
	}

	var obj []*eventstorage.Object
	{
		obj, err = h.eve.SearchEvnt(eve)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for _, x := range obj {
		if userid.FromContext(ctx) != x.User {
			return nil, tracer.Mask(userNotOwnerError)
		}
	}

	var out []objectstate.String
	{
		out, err = h.eve.DeleteWrkr(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

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
