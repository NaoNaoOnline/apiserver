package labelstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp *Object) (*Object, error) {
	var err error

	// At first we need to validate the given input object and, amongst others,
	// whether the label does already exist, since our label names must be unique.
	{
		err = r.validateCreate(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		inp.Crtd = time.Now().UTC()
		inp.Labl = scoreid.New(inp.Crtd)
	}

	// Once we know the label is unique, we create the normalized key-value pair
	// so that we can search for label objects using their IDs.
	{
		err = r.red.Simple().Create().Element(labObj(inp.Labl), musStr(inp))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Now we create the global and user specific mappings for global and user
	// specific search queries. For the global queries we ensure the label names
	// are unique by using the label name as additional index within the redis
	// sorted sets.
	{
		err = r.red.Sorted().Create().Element(labKin(inp.Kind), inp.Labl.String(), inp.Labl.Float(), inp.Name)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		err = r.red.Sorted().Create().Element(labUse(inp.User), inp.Labl.String(), inp.Labl.Float())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return inp, nil
}

func (r *Redis) validateCreate(inp *Object) error {
	if inp.Kind != "cate" && inp.Kind != "host" {
		return tracer.Mask(invalidLabelKindError)
	}

	if inp.Name == "" {
		return tracer.Mask(labelNameEmptyError)
	}

	if inp.User == "" {
		return tracer.Mask(userIDEmptyError)
	}

	{
		exi, err := r.red.Sorted().Exists().Index(labKin(inp.Kind), inp.Name)
		if err != nil {
			return tracer.Mask(err)
		}

		if exi {
			return tracer.Mask(labelAlreadyExistsError)
		}
	}

	return nil
}
