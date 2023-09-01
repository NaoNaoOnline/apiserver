package descriptionstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp *Object) (*Object, error) {
	var err error

	// At first we need to validate the given input object and, amongst others,
	// ensure whether the event mapped to the description does already exist. For
	// instance, we cannot create a description for an event that is not there.
	{
		err = r.validateCreate(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		inp.Crtd = time.Now().UTC()
		inp.Desc = objectid.New(inp.Crtd)
	}

	// Once we know the associated event exists, we create the normalized
	// key-value pair for the description object, so that we can search for
	// description objects using their IDs.
	{
		err = r.red.Simple().Create().Element(desObj(inp.Desc), musStr(inp))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Now we create the event and user specific mappings for event and user
	// specific search queries.
	{
		err = r.red.Sorted().Create().Score(desEve(inp.Evnt), inp.Desc.String(), inp.Desc.Float())
		if err != nil {
			return nil, tracer.Mask(err)
		}

		err = r.red.Sorted().Create().Score(desUse(inp.User), inp.Desc.String(), inp.Desc.Float())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return inp, nil
}

func (r *Redis) validateCreate(inp *Object) error {
	if inp.Text == "" {
		return tracer.Mask(descriptionTextEmptyError)
	}

	if inp.User == "" {
		return tracer.Mask(userIDEmptyError)
	}

	{
		cou, err := r.red.Simple().Exists().Multi(eveObj(inp.Evnt))
		if err != nil {
			return tracer.Mask(err)
		}

		if cou != 1 {
			return tracer.Mask(eventNotFoundError)
		}
	}

	return nil
}
