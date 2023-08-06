package labelstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp *Object) (*Object, error) {
	var err error

	{
		inp.Crtd = time.Now().UTC()
		inp.Labl = scoreid.New(inp.Crtd)
	}

	if inp.Kind != "cate" && inp.Kind != "host" {
		return nil, tracer.Mask(invalidInputError)
	}

	// At first we need to validate whether the label does already exist, since
	// our label names must be unique.
	{
		exi, err := r.red.Sorted().Exists().Index(keyKin(inp.Kind), inp.Name)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if exi {
			return nil, tracer.Mask(alreadyExistsError)
		}
	}

	// Once we know the label is unique, we create the normalized key-value pair
	// so that we can search for label objects using their IDs.
	{
		err = r.red.Simple().Create().Element(keyObj(inp.Labl), musStr(inp))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Now we create the global and user specific mappings for global and user
	// specific search queries. For the global queries we ensure the label names
	// are unique by using the label name as additional index within the redis
	// sorted sets.
	{
		err = r.red.Sorted().Create().Element(keyKin(inp.Kind), inp.Labl.String(), inp.Labl.Float(), inp.Name)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		err = r.red.Sorted().Create().Element(keyUse(inp.User), inp.Labl.String(), inp.Labl.Float())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return inp, nil
}
