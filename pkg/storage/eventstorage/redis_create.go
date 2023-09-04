package eventstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp []*Object) ([]*Object, error) {
	var err error

	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// ensure whether the labels mapped to the event do already exist. For
		// instance, we cannot create an event for a label that is not there.
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			var key []string
			for _, x := range append(inp[i].Cate, inp[i].Host...) {
				key = append(key, labObj(x))
			}

			cou, err := r.red.Simple().Exists().Multi(key...)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			if cou != int64(len(key)) {
				return nil, tracer.Maskf(labelObjectNotFoundError, "%d labels do not exist", int64(len(key))-cou)
			}
		}

		{
			inp[i].Crtd = time.Now().UTC()
			inp[i].Evnt = objectid.New(inp[i].Crtd)
		}

		// Once we know the associated labels exist, we create the normalized
		// key-value pair for the event object, so that we can search for event
		// objects using their IDs.
		{
			err = r.red.Simple().Create().Element(eveObj(inp[i].Evnt), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the time specific mappings for time specific search queries. With
		// that we can search for events that are happening right now. For event
		// time indexing, we use the event time as score. There might be many events
		// happening at the same time. So we use Sorted.Create.Value, which allows
		// us to use duplicated scores.
		{
			err = r.red.Sorted().Create().Value(keyfmt.EventTime, inp[i].Evnt.String(), float64(inp[i].Time.Unix()))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the label specific mappings for label specific search queries.
		// With that we can search for events associated to certain labels.
		for _, x := range append(inp[i].Cate, inp[i].Host...) {
			err = r.red.Sorted().Create().Score(eveLab(x), inp[i].Evnt.String(), inp[i].Evnt.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the user specific mappings for user specific search queries. With
		// that we can show the user all events they created.
		{
			err = r.red.Sorted().Create().Score(eveUse(inp[i].User), inp[i].Evnt.String(), inp[i].Evnt.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}
