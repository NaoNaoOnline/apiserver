package reactionstorage

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
		// whether the reaction does already exist, since our reaction names must be
		// unique.
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			exi, err := r.red.Sorted().Exists().Index(rctKin(inp[i].Kind), keyfmt.Indx(inp[i].Name))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			if exi {
				return nil, tracer.Maskf(reactionObjectAlreadyExistsError, keyfmt.Indx(inp[i].Name))
			}
		}

		{
			inp[i].Crtd = time.Now().UTC()
			inp[i].Rctn = objectid.New(inp[i].Crtd)
			inp[i].Name = keyfmt.Name(inp[i].Name)
		}

		// Once we know the reaction is unique, we create the normalized key-value
		// pair so that we can search for reaction objects using their IDs.
		{
			err = r.red.Simple().Create().Element(rctObj(inp[i].Rctn), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Now we create the reaction kind mappings for reaction kind search
		// queries. With that we can search for reactions of a given kind. Note that
		// we ensure the reaction names are unique by using the reaction name as
		// additional index within the redis sorted sets.
		{
			err = r.red.Sorted().Create().Index(rctKin(inp[i].Kind), inp[i].Rctn.String(), inp[i].Rctn.Float(), keyfmt.Indx(inp[i].Name))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the user specific mappings for user specific search queries. With
		// that we can show the user all reactions they created.
		{
			err = r.red.Sorted().Create().Score(rctUse(inp[i].User), inp[i].Rctn.String(), inp[i].Rctn.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}
