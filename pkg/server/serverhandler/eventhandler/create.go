package eventhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *event.CreateI) (*event.CreateO, error) {
	var err error

	var inp []*eventstorage.Object
	for _, x := range req.Object {
		if x.Public != nil {
			inp = append(inp, &eventstorage.Object{
				Cate: inpLab(x.Public.Cate),
				Dura: inpDur(x.Public.Dura),
				Host: inpLab(x.Public.Host),
				Link: x.Public.Link,
				Time: inpTim(x.Public.Time),
				User: userid.FromContext(ctx),
			})
		}
	}

	//
	// Create the given resources.
	//

	var out []*eventstorage.Object
	{
		out, err = h.eve.CreateEvnt(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Create background tasks for the created resources.
	//

	{
		_, err = h.eve.CreateWrkr(out)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *event.CreateO
	{
		res = &event.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &event.CreateO_Object{
			Intern: &event.CreateO_Object_Intern{
				Crtd: outTim(x.Crtd),
				Evnt: x.Evnt.String(),
			},
		})
	}

	return res, nil
}
