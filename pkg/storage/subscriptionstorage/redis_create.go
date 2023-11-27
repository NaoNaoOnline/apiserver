package subscriptionstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreateSubs(inp []*Object) ([]*Object, error) {
	var err error

	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// whether the subscription addresses comply with the expected format.
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
			inp[i].Stts = objectstate.Created
			inp[i].Subs = objectid.Random(objectid.Time(now))
		}

		// Once we know the subscription input is valid, we create the normalized
		// key-value pair so that we can search for subscription objects using their
		// IDs.
		{
			err = r.red.Simple().Create().Element(subObj(inp[i].Subs), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the subscriber specific mappings for subscriber specific search
		// queries. With that we can show the subscriber all subscriptions they had.
		// That is, the subscriptions they enjoy to access premium features.
		{
			err = r.red.Sorted().Create().Score(subAdd(inp[i].Sbsc), inp[i].Subs.String(), inp[i].Subs.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the user specific mappings for user specific search queries. With
		// that we can show the user all subscriptions they created. That is, the
		// subscriptions they paid for.
		{
			err = r.red.Sorted().Create().Score(subUse(inp[i].User), inp[i].Subs.String(), inp[i].Subs.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}

func (r *Redis) CreateWrkr(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		{
			err = r.emi.Scrape(inp[i].Subs)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			out = append(out, objectstate.Started)
		}
	}

	return out, nil
}
