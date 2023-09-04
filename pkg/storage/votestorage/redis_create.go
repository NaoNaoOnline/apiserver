package votestorage

import (
	"encoding/json"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/redigo/pkg/sorted"
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
		var jsn string
		{
			jsn, err = r.red.Simple().Search().Value(desObj(inp[i].Desc))
			if simple.IsNotFound(err) {
				return nil, tracer.Maskf(descriptionNotFoundError, inp[i].Desc.String())
			} else if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var obj *descriptionstorage.Object
		{
			obj = &descriptionstorage.Object{}
		}

		{
			err = json.Unmarshal([]byte(jsn), obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// We need to validate the given input object and, amongst others, whether
		// the mapped reaction object does in fact exist, since we must not create
		// broken mappings.
		{
			err = r.validateCreate(obj.Evnt, inp[i])
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			inp[i].Crtd = time.Now().UTC()
			inp[i].Vote = objectid.New(inp[i].Crtd)
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
		// votes created per event per user. So we use Sorted.Create.Score, which
		// allows us to use unique scores.
		{
			err = r.red.Sorted().Create().Score(eveVot(inp[i].User), obj.Evnt.String(), obj.Evnt.Float())
			if sorted.IsAlreadyExistsError(err) {
				// fall through
			} else if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Now we create the event/user specific mappings for event/user specific
		// search queries. This allows us to search for the amount of votes a user
		// made on an event.
		{
			err = r.red.Sorted().Create().Score(votUse(obj.Evnt, inp[i].User), inp[i].Vote.String(), inp[i].Vote.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}

func (r *Redis) validateCreate(eve objectid.String, inp *Object) error {
	if inp.Desc == "" {
		return tracer.Mask(descriptionIDEmptyError)
	}

	if inp.Rctn == "" {
		return tracer.Mask(reactionIDEmptyError)
	}

	{
		res, err := r.red.Sorted().Search().Order(votUse(eve, inp.User), 0, -1)
		if err != nil {
			return tracer.Mask(err)
		}

		if len(res) >= 5 {
			return tracer.Mask(voteLimitError)
		}
	}

	if !r.rct.Exists(inp.Rctn) {
		return tracer.Maskf(reactionNotFoundError, inp.Rctn.String())
	}

	return nil
}
