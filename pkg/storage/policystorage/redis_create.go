package policystorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp []*Object) ([]*Object, error) {
	var err error

	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// whether the record represents a valid SMA.
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var now time.Time
		{
			now = time.Now().UTC()
		}

		{
			inp[i].Crtd = now
			inp[i].Plcy = objectid.Random(objectid.Time(now))
		}

		// Once we know the record is valid, we create the normalized key-value pair
		// so that we can search for policy objects using their IDs.
		{
			err = r.red.Simple().Create().Element(polObj(inp[i].Plcy), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Now we create the record kind mappings for record kind search queries.
		// With that we can search for records of a given kind. That is, records of
		// a particular smart contract event.
		{
			err = r.red.Sorted().Create().Score(polKin(inp[i].Kind), inp[i].Plcy.String(), inp[i].Plcy.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}
