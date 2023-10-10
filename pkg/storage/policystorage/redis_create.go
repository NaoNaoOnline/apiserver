package policystorage

import (
	"strconv"
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

		// Create the record kind mappings for record kind search queries. With that
		// we can search for records of a given kind. That is, records of a
		// particular smart contract event.
		{
			err = r.red.Sorted().Create().Score(polKin(inp[i].Kind), inp[i].Plcy.String(), inp[i].Plcy.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the policy member mappings for policy member search queries. With
		// that we can answer the question whether a given member is part of any
		// policy.
		{
			err = r.red.Sorted().Create().Score(polMem(inp[i].Kind, inp[i].Memb), inp[i].Plcy.String(), inp[i].Plcy.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create policy system mappings for policy system search queries. With that
		// we can search for members within a specific system and answer the
		// question whether they have a particular access level in that system.
		{
			err = r.red.Sorted().Create().Score(polSys(inp[i].Kind, inp[i].Syst, inp[i].Memb), strconv.FormatInt(inp[i].Acce, 10), inp[i].Plcy.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}
