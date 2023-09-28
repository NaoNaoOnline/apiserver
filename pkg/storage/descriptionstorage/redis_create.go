package descriptionstorage

import (
	"encoding/json"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp []*Object) ([]*Object, error) {
	var err error

	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// ensure whether the event mapped to the description does already exist. For
		// instance, we cannot create a description for an event that is not there.
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// We want to verify whether the associated event has already happened. For
		// that we have to fetch the event object, so we can access its time
		// information.
		var jsn []string
		{
			jsn, err = r.red.Simple().Search().Multi(eveObj(inp[i].Evnt))
			if simple.IsNotFound(err) {
				return nil, tracer.Maskf(eventObjectNotFoundError, inp[i].Evnt.String())
			} else if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var obj *eventstorage.Object
		{
			obj = &eventstorage.Object{}
		}

		{
			err = json.Unmarshal([]byte(jsn[0]), obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Ensure descriptions cannot be added to events that have already happened.
		if obj.Happnd() {
			return nil, tracer.Mask(eventAlreadyHappenedError)
		}

		{
			inp[i].Crtd = time.Now().UTC()
			inp[i].Desc = objectid.New(inp[i].Crtd)
		}

		// Once we know the associated event exists, we create the normalized
		// key-value pair for the description object, so that we can search for
		// description objects using their IDs.
		{
			err = r.red.Simple().Create().Element(desObj(inp[i].Desc), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Now we create the event and user specific mappings for event and user
		// specific search queries.
		{
			err = r.red.Sorted().Create().Score(desEve(inp[i].Evnt), inp[i].Desc.String(), inp[i].Desc.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}

			err = r.red.Sorted().Create().Score(desUse(inp[i].User), inp[i].Desc.String(), inp[i].Desc.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}
