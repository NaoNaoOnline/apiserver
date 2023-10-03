package votehandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/vote"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, req *vote.DeleteI) (*vote.DeleteO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	var vot []objectid.ID
	for _, x := range req.Object {
		vot = append(vot, objectid.ID(x.Intern.Vote))
	}

	var inp []*votestorage.Object
	{
		inp, err = h.vot.SearchVote(vot)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for _, x := range inp {
		if userid.FromContext(ctx) != x.User {
			return nil, tracer.Mask(userNotOwnerError)
		}

		var des []*descriptionstorage.Object
		{
			des, err = h.des.SearchDesc([]objectid.ID{x.Desc})
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Ensure votes cannot be removed from descriptions that have already been
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

		// Ensure votes cannot be removed from events that have already been
		// deleted.
		if !eve[0].Dltd.IsZero() {
			return nil, tracer.Mask(eventDeletedError)
		}

		// Ensure votes cannot be removed from events that have already happened.
		if eve[0].Happnd() {
			return nil, tracer.Mask(eventAlreadyHappenedError)
		}
	}

	var out []objectstate.String
	{
		out, err = h.vot.Delete(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *vote.DeleteO
	{
		res = &vote.DeleteO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &vote.DeleteO_Object{
			Intern: &vote.DeleteO_Object_Intern{
				Stts: x.String(),
			},
			Public: &vote.DeleteO_Object_Public{},
		})
	}

	return res, nil
}
