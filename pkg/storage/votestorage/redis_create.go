package votestorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp []*Object) ([]*Object, error) {
	var err error

	for i := range inp {
		// We need to validate the given input object and, amongst others, whether
		// the mapped description object does in fact exist, since we must not
		// create broken mappings. We do also need the event ID for user specific
		// mappings, and so we search for the description object and get the event
		// ID from there, catching two birds with one stone.
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var des *descriptionstorage.Object
		{
			des, err = r.searchDesc(inp[i].Desc)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var eve *eventstorage.Object
		{
			eve, err = r.searchEvnt(des.Evnt)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Ensure votes cannot be added to events that have already happened.
		{
			if eve.Happnd() {
				return nil, tracer.Mask(eventAlreadyHappenedError)
			}
		}

		// Ensure the reaction used by the user does in fact exist.
		{
			cou, err := r.red.Simple().Exists().Multi(rctObj(inp[i].Rctn))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			if cou != 1 {
				return nil, tracer.Maskf(reactionObjectNotFoundError, inp[i].Rctn.String())
			}
		}

		{
			inp[i].Crtd = time.Now().UTC()
			inp[i].Evnt = eve.Evnt
			inp[i].Vote = objectid.New(inp[i].Crtd)
		}

		// Ensure the user vote limit globally is respected.
		{
			cou, err := r.red.Sorted().Metric().Count(votUse(inp[i].User))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			if cou >= 100 {
				return nil, tracer.Mask(voteUserLimitError)
			}
		}

		// Ensure the user vote limit per event is respected.
		{
			cou, err := r.red.Sorted().Metric().Count(votEve(inp[i].User, inp[i].Evnt))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			if cou >= 5 {
				return nil, tracer.Mask(voteEventLimitError)
			}
		}

		// Once we know the mapped objects exist, we create the normalized key-value
		// pair so that we can search for vote objects using their IDs.
		{
			err = r.red.Simple().Create().Element(votObj(inp[i].Vote), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Now we create the description specific mappings for description specific
		// search queries. This allows us to search for all votes associated to a
		// description.
		{
			err = r.red.Sorted().Create().Score(votDes(inp[i].Desc), inp[i].Vote.String(), inp[i].Vote.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the reaction specific mappings for reaction specific search
		// queries. With that we can search for events that the user reacted to. For
		// user reaction indexing, we use the event ID as score. There might be many
		// votes created per event per user. So we use Sorted.Create.Value, which
		// allows us to use duplicated scores.
		{
			err = r.red.Sorted().Create().Value(votUse(inp[i].User), inp[i].Vote.String(), inp[i].Evnt.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Now we create the user/event specific mappings for user/event specific
		// search queries. This allows us to search for the amount of votes a user
		// made on an event.
		{
			err = r.red.Sorted().Create().Score(votEve(inp[i].User, inp[i].Evnt), inp[i].Vote.String(), inp[i].Vote.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}
