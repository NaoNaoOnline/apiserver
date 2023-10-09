package descriptionhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *description.CreateI) (*description.CreateO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	var inp []*descriptionstorage.Object
	for _, x := range req.Object {
		inp = append(inp, &descriptionstorage.Object{
			Evnt: objectid.ID(x.Public.Evnt),
			Text: x.Public.Text,
			User: userid.FromContext(ctx),
		})
	}

	for _, x := range inp {
		var eve []*eventstorage.Object
		{
			eve, err = h.eve.SearchEvnt([]objectid.ID{x.Evnt})
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Ensure descriptions cannot be added to events that have already been
		// deleted.
		if !eve[0].Dltd.IsZero() {
			return nil, tracer.Mask(eventDeletedError)
		}

		// Ensure descriptions cannot be added to events that have already happened.
		if eve[0].Happnd() {
			return nil, tracer.Mask(eventAlreadyHappenedError)
		}
	}

	var out []*descriptionstorage.Object
	{
		out, err = h.des.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *description.CreateO
	{
		res = &description.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &description.CreateO_Object{
			Intern: &description.CreateO_Object_Intern{
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				Desc: x.Desc.String(),
			},
		})
	}

	return res, nil
}
