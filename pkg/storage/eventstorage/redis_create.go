package eventstorage

import (
	"net/url"
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
			err = r.validateCreate(inp[i])
			if err != nil {
				return nil, tracer.Mask(err)
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

func (r *Redis) validateCreate(inp *Object) error {
	if len(inp.Cate) > 5 {
		return tracer.Maskf(tooManyLabelsError, "allowed are up to 5 category labels")
	}

	if inp.Dura == 0 {
		return tracer.Mask(eventDurationEmptyError)
	}
	if inp.Dura < 0 {
		return tracer.Mask(eventDurationNegativeError)
	}

	if len(inp.Host) > 5 {
		return tracer.Maskf(tooManyLabelsError, "allowed are up to 5 host labels")
	}

	if !valLin(inp.Link) {
		return tracer.Mask(eventLinkInvalidError)
	}

	if inp.Time.IsZero() {
		return tracer.Maskf(eventTimeInvalidError, "time must not be empty")
	}
	if inp.Time.Compare(time.Now().UTC()) != +1 {
		return tracer.Maskf(eventTimeInvalidError, "time must be in the future")
	}

	if inp.User == "" {
		return tracer.Mask(userIDEmptyError)
	}

	{
		var key []string
		for _, x := range append(inp.Cate, inp.Host...) {
			key = append(key, labObj(x))
		}

		cou, err := r.red.Simple().Exists().Multi(key...)
		if err != nil {
			return tracer.Mask(err)
		}

		if cou != int64(len(key)) {
			return tracer.Maskf(labelNotFoundError, "%d labels do not exist", int64(len(key))-cou)
		}
	}

	return nil
}

func valLin(str string) bool {
	if str == "" {
		return false
	}

	poi, err := url.Parse(str)
	if err != nil {
		return false
	}

	if poi.Scheme != "https" {
		return false
	}

	return true
}
