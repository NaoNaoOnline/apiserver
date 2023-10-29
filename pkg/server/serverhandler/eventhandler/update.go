package eventhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *event.UpdateI) (*event.UpdateO, error) {
	var err error

	var out []objectstate.String

	//
	// Track external like
	//

	{
		var eve []objectid.ID
		for _, x := range req.Object {
			if x.Intern != nil && x.Symbol != nil && x.Intern.Evnt != "" && x.Symbol.Link == "add" {
				eve = append(eve, objectid.ID(x.Intern.Evnt))
			}
		}

		if len(eve) != 0 {
			var inp []*eventstorage.Object
			{
				inp, err = h.eve.SearchEvnt(eve)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			for i := range inp {
				// Ensure events cannot be clicked if they have already been deleted.
				if !inp[i].Dltd.IsZero() {
					return nil, tracer.Mask(eventDeletedError)
				}

				// Ensure events cannot be clicked if they have already happened.
				if inp[i].Happnd() {
					return nil, tracer.Mask(eventAlreadyHappenedError)
				}
			}

			{
				lis, err := h.eve.UpdateClck(inp)
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

	var res *event.UpdateO
	{
		res = &event.UpdateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &event.UpdateO_Object{
			Intern: &event.UpdateO_Object_Intern{
				Stts: x.String(),
			},
			Public: &event.UpdateO_Object_Public{},
		})
	}

	return res, nil
}
