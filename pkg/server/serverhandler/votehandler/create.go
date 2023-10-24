package votehandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/vote"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *vote.CreateI) (*vote.CreateO, error) {
	var err error

	var inp []*votestorage.Object
	for _, x := range req.Object {
		inp = append(inp, &votestorage.Object{
			Desc: objectid.ID(x.Public.Desc),
			Rctn: objectid.ID(x.Public.Rctn),
			User: userid.FromContext(ctx),
		})
	}

	for i, x := range inp {
		var des []*descriptionstorage.Object
		{
			des, err = h.des.SearchDesc([]objectid.ID{x.Desc})
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		if len(des) != 1 {
			return nil, tracer.Mask(runtime.ExecutionFailedError)
		}

		// Ensure votes cannot be added to descriptions that have already been
		// deleted.
		if !des[0].Dltd.IsZero() {
			return nil, tracer.Mask(descriptionDeletedError)
		}

		var eve []*eventstorage.Object
		{
			eve, err = h.eve.SearchEvnt([]objectid.ID{des[0].Evnt})
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Set the event ID for our internal reference. We do this in the handler
		// because the storage should not search the event again. So the storage
		// relies on us doing it here.
		{
			inp[i].Evnt = eve[0].Evnt
		}

		// Ensure votes cannot be added to events that have already been deleted.
		if !eve[0].Dltd.IsZero() {
			return nil, tracer.Mask(eventDeletedError)
		}

		// Ensure votes cannot be added to events that have already happened.
		if eve[0].Happnd() {
			return nil, tracer.Mask(eventAlreadyHappenedError)
		}
	}

	var out []*votestorage.Object
	{
		out, err = h.vot.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct RPC response.
	//

	var res *vote.CreateO
	{
		res = &vote.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &vote.CreateO_Object{
			Intern: &vote.CreateO_Object_Intern{
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				Vote: x.Vote.String(),
			},
		})
	}

	return res, nil
}
