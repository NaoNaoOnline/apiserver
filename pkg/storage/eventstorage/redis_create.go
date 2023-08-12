package eventstorage

import (
	"net/url"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp *Object) (*Object, error) {
	var err error

	// At first we need to validate the given input object and, amongst others,
	// ensure whether the labels mapped to the event do already exist. For
	// instance, we cannot create an event for a label that is not there.
	{
		err = r.validateCreate(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		inp.Crtd = time.Now().UTC()
		inp.Evnt = scoreid.New(inp.Crtd)
	}

	// Once we know the associated labels exist, we create the normalized
	// key-value pair for the event object, so that we can search for event
	// objects using their IDs.
	{
		err = r.red.Simple().Create().Element(eveObj(inp.Evnt), musStr(inp))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Now we create the label and user specific mappings for label and user
	// specific search queries.
	for _, x := range append(inp.Cate, inp.Host) {
		err = r.red.Sorted().Create().Element(eveLab(x), inp.Evnt.String(), inp.Evnt.Float())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		err = r.red.Sorted().Create().Element(eveUse(inp.User), inp.Evnt.String(), inp.Evnt.Float())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return inp, nil
}

func (r *Redis) validateCreate(inp *Object) error {
	if inp.Dura == 0 {
		return tracer.Mask(eventDurationEmptyError)
	}

	if !valLin(inp.Link) {
		return tracer.Mask(eventLinkInvalidError)
	}

	if inp.Time.IsZero() {
		return tracer.Maskf(eventTimeInvalidError, "time must not be empty")
	}
	if inp.Time.Compare(time.Now()) != +1 {
		return tracer.Maskf(eventTimeInvalidError, "time must be in the future")
	}

	if inp.User == "" {
		return tracer.Mask(userIDEmptyError)
	}

	{
		var key []string
		{
			key = append(key, labObj(inp.Host))
		}

		for _, x := range inp.Cate {
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
