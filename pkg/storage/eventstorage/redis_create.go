package eventstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/redigo/pkg/sorted"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreateEvnt(inp []*Object) ([]*Object, error) {
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

		// Check if the labels of the event we want to create do even exist.
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
				return nil, tracer.Maskf(labelObjectNotFoundError, "%d of these labels do not exist on event %v %v", int64(len(key))-cou, inp[i].Evnt, key)
			}
		}

		// Check whether we can assign the natively supported platform label based
		// on the event link.
		{
			key, err := r.red.Sorted().Search().Index(keyfmt.LabelSystem, keyfmt.Indx(inp[i].Pltfrm()))
			if sorted.IsNotFound(err) {
				// fall through
			} else if err != nil {
				return nil, tracer.Mask(err)
			} else {
				inp[i].Bltn = append(inp[i].Bltn, objectid.ID(key))
			}
		}

		// Make sure the hosts of the event being created are not already indexed to
		// be online during the new expected event duration. Here we need to look
		// for all events within the following "minus-four-hours" time range,
		// because events can be 4 hours long.
		//
		//     < minus four hours > < new event start > < new event end >
		//
		{
			var min time.Time
			var max time.Time
			{
				min = inp[i].Time.Add(-(time.Hour * 4))
				max = inp[i].Time.Add(inp[i].Dura)
			}

			var obj []*Object
			{
				obj, err = r.SearchTime(min, max)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			// If any of inp[i].Host can be found in obj[j].Host, and if both event
			// durations overlap on the timeline, then we return an error, because
			// neither host for the event we want to create can be on two events
			// simultaneously.
			if inp[i].Ovrlap(obj) {
				return nil, tracer.Mask(hostParticipationConflictError)
			}
		}

		var now time.Time
		{
			now = time.Now().UTC()
		}

		{
			inp[i].Crtd = now
			inp[i].Evnt = objectid.Random(objectid.Time(now))
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

		var amn int64
		{
			amn, err = r.red.Sorted().Metric().Count(eveUse(inp[i].User))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Track the given user as content creator by incrementing the amount of
		// events added by 1. With this globally maintained sorted set we can search
		// for all content creators and donate premium subscriptions to those that
		// meet certain criteria of what we consider legitimate content creators. It
		// is important to only maintain users in this particular sorted set that
		// have created a minimum amount of events within our specified rolling time
		// window. This time window is the event retention in Redis, which is
		// currently set to 3 months. Note that we need to add 1 to our check
		// manually here since we are processing the additional event that is being
		// added right now.

		if (amn + 1) == int64(r.mse) {
			err = r.red.Sorted().Create().Score(keyfmt.EventCreator, inp[i].User.String(), float64(r.mse))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		if (amn + 1) > int64(r.mse) {
			_, err = r.red.Sorted().Floats().Score(keyfmt.EventCreator, inp[i].User.String(), +1.0)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the time specific mappings for time specific search queries. With
		// that we can search for events that are happening right now. For event
		// time indexing, we use the event time as score. There might be many events
		// happening at the same time. So we use Sorted.Create.Score, which allows
		// us to use duplicated scores.
		{
			err = r.red.Sorted().Create().Score(keyfmt.EventTime, inp[i].Evnt.String(), float64(inp[i].Time.Unix()))
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

func (r *Redis) CreateWrkr(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		{
			err = r.emi.CreateEvnt(inp[i].Evnt)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			out = append(out, objectstate.Started)
		}
	}

	return out, nil
}
