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

		// We need to figure out whether a record with equal SMA fields does already
		// exist. If the given record does already exist on another chain, then we
		// will find exactly one element, because there can only ever be one element
		// for any combination of kind / system / member. Since we call
		// Sorted.Search.Order with true, we will receive values and scores.
		// Therefore we expect a list of length 2, where the first item, the value,
		// is the SMA access level, and where the second item, the score, is the
		// object ID of the existing policy object.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(polSys(inp[i].Kind, inp[i].Syst, inp[i].Memb), 0, -1, true)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		if len(val) == 2 && val[0] == strconv.FormatInt(inp[i].Acce, 10) {
			inp[i], err = r.update(objectid.ID(val[1]), inp[i])
			if err != nil {
				return nil, tracer.Mask(err)
			}
		} else {
			inp[i], err = r.create(inp[i])
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}

func (r *Redis) create(inp *Object) (*Object, error) {
	var err error

	var now time.Time
	{
		now = time.Now().UTC()
	}

	{
		inp.Crtd = now
		inp.Plcy = objectid.Random(objectid.Time(now))
	}

	// Once we know the record is valid, we create the normalized key-value pair
	// so that we can search for policy objects using their IDs.
	{
		err = r.red.Simple().Create().Element(polObj(inp.Plcy), musStr(inp))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Create the record kind mappings for record kind search queries. With that
	// we can search for records of a given kind. That is, records of a
	// particular smart contract event.
	{
		err = r.red.Sorted().Create().Score(polKin(inp.Kind), inp.Plcy.String(), inp.Plcy.Float())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Create the policy member mappings for policy member search queries. With
	// that we can answer the question whether a given member is part of any
	// policy.
	{
		err = r.red.Sorted().Create().Score(polMem(inp.Kind, inp.Memb), inp.Plcy.String(), inp.Plcy.Float())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Create policy system mappings for policy system search queries. With that
	// we can search for members within a specific system and answer the
	// question whether they have a particular access level in that system.
	{
		err = r.red.Sorted().Create().Score(polSys(inp.Kind, inp.Syst, inp.Memb), strconv.FormatInt(inp.Acce, 10), inp.Plcy.Float())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return inp, nil
}

func (r *Redis) update(oid objectid.ID, inp *Object) (*Object, error) {
	var err error

	// In order to update the existing policy object we fetch it first.
	var exi []*Object
	{
		exi, err = r.SearchPlcy([]objectid.ID{oid})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Add the new chain specific information to the existing policy object.
	{
		exi[0].Blck = append(exi[0].Blck, inp.Blck...)
		exi[0].ChID = append(exi[0].ChID, inp.ChID...)
		exi[0].From = append(exi[0].From, inp.From...)
		exi[0].Hash = append(exi[0].Hash, inp.Hash...)
		exi[0].Time = append(exi[0].Time, inp.Time...)
	}

	// After merging the existing policy object with the new record we need to
	// verify whether the policy object is still valid. This may help prevent
	// issues caused by event indexing or duplicated chain IDs.
	{
		err := exi[0].Verify()
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Update the normalized key-value pair.
	{
		err = r.red.Simple().Create().Element(polObj(exi[0].Plcy), musStr(exi[0]))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return exi[0], nil
}
