package descriptionhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *description.CreateI) (*description.CreateO, error) {
	var err error

	var inp []*descriptionstorage.Object
	for _, x := range req.Object {
		if x.Public != nil {
			inp = append(inp, &descriptionstorage.Object{
				Evnt: objectid.ID(x.Public.Evnt),
				Text: objectfield.String{
					Data: x.Public.Text,
				},
				User: userid.FromContext(ctx),
			})
		}
	}

	//
	// Verify the given input.
	//

	{
		err = h.createVrfy(ctx, inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Create the given resources.
	//

	var out []*descriptionstorage.Object
	{
		out, err = h.des.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

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

func (h *Handler) createVrfy(ctx context.Context, inp descriptionstorage.Slicer) error {
	var err error

	var eve []*eventstorage.Object
	{
		eve, err = h.eve.SearchEvnt("", inp.Evnt())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range eve {
		// Ensure descriptions cannot be added to events that have already been
		// deleted.
		if !x.Dltd.IsZero() {
			return tracer.Mask(eventDeletedError)
		}

		// Ensure descriptions cannot be added to events that have already happened.
		if x.Happnd() {
			return tracer.Mask(eventAlreadyHappenedError)
		}
	}

	return nil
}
