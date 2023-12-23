package eventhandler

import (
	"context"
	"time"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
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
				Mtrc: objectfield.MapInt{
					Data: map[string]int64{},
				},
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

	go func() {
		// We defer the task creation by 5 seconds because we want both the event
		// object and the description object to be created in most cases when the
		// "event create" task is being executed. The event and description
		// resources are decoupled and the description object gets created as a very
		// last step. It can happen that the processing of the "event create" task
		// has already happened while the description object was not yet created.
		// There is a systemic race condition here. Delaying by some arbitrary
		// amount of time like 5 seconds should remedy the issue in most cases. One
		// alternative would be to create the "event create" task once the first
		// description got created on an event, but this is quite ugly as well and
		// not bullet proof either.
		{
			time.Sleep(5 * time.Second)
		}

		{
			_, err = h.eve.CreateWrkr(out)
			if err != nil {
				h.log.Log(ctx, "level", "error", "message", err.Error())
			}
		}
	}()

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
