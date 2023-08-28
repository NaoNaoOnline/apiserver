package votestorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp []*Object) ([]*Object, error) {
	var err error

	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// whether the mapped description and reaction objects do in fact exist,
		// since we must not create broken mappings.
		{
			err = r.validateCreate(inp[i])
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
		// search queries.
		{
			err = r.red.Sorted().Create().Element(votDes(inp[i].Desc), inp[i].Vote.String(), inp[i].Vote.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}

func (r *Redis) validateCreate(inp *Object) error {
	if inp.Desc == "" {
		return tracer.Mask(descriptionIDEmptyError)
	}

	if inp.Rctn == "" {
		return tracer.Mask(reactionIDEmptyError)
	}

	{
		exi, err := r.red.Simple().Exists().Multi(desObj(inp.Desc))
		if err != nil {
			return tracer.Mask(err)
		}

		if exi != 1 {
			return tracer.Maskf(descriptionNotFoundError, inp.Desc.String())
		}
	}

	if !r.rct.Exists(inp.Rctn) {
		return tracer.Maskf(reactionNotFoundError, inp.Rctn.String())
	}

	return nil
}
