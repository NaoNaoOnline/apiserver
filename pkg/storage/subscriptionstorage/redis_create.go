package subscriptionstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreateSubs(inp []*Object) ([]*Object, error) {
	var err error

	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// whether the creator addresses comply with the expected format.
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Ensure the given creator addresses meet our criteria of legitimate
		// content creators.
		var vld []bool
		{
			vld, err = r.VerifyAddr(inp[i].Crtr)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		if !vldAdd(vld) {
			return nil, tracer.Mask(subscriptionCrtrInvalidError)
		}

		// Lookup the user ID of the recipient using the subscriber address. There
		// should be exactly one user ID for any given address.
		var rec []objectid.ID
		{
			rec, _, err = r.wal.SearchAddr([]string{inp[i].Recv})
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		if len(rec) != 1 {
			return nil, tracer.Mask(runtime.ExecutionFailedError)
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

		// Create the receiver specific mappings for receiver specific search
		// queries. With that we can show the user all subscriptions they received.
		// That is, the subscriptions they use to access premium features.
		{
			err = r.red.Sorted().Create().Score(subRec(rec[0]), inp[i].Subs.String(), inp[i].Subs.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the payer specific mappings for payer specific search queries.
		// With that we can show the user all subscriptions they created. That is,
		// the subscriptions they paid for.
		{
			err = r.red.Sorted().Create().Score(subPay(inp[i].User), inp[i].Subs.String(), inp[i].Subs.Float())
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

func vldAdd(vld []bool) bool {
	for _, x := range vld {
		if !x {
			return false
		}
	}

	return true
}
