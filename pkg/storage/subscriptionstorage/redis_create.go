package subscriptionstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreateSubs(inp []*Object) ([]*Object, error) {
	var err error

	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// whether the creator addresses comply with the expected format.
		{
			err := inp[i].VerifyObct()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Ensure the user IDs of the subscription payer and receiver do in fact
		// exist.
		var use []*userstorage.Object
		{
			use, err = r.use.SearchUser([]objectid.ID{inp[i].Payr, inp[i].Rcvr})
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		if len(use) != 2 {
			return nil, tracer.Mask(userNotFoundError)
		}

		var exi []*Object
		{
			exi, err = r.SearchRcvr([]objectid.ID{inp[i].Rcvr}, PagAll())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Check whether the created subscription for the given receiver is
		// effectively a renewal.
		if subRen(exi, inp[i]) {
			err = inp[i].VerifyUnix(VerifyRenw(time.Now().UTC()))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		} else {
			err = inp[i].VerifyUnix(VerifyOnce(time.Now().UTC()))
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

		// Create the payer specific mappings for payer specific search queries.
		// With that we can show the user all subscriptions they paid for.
		{
			err = r.red.Sorted().Create().Score(subPay(inp[i].Payr), inp[i].Subs.String(), inp[i].Subs.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the receiver specific mappings for receiver specific search
		// queries. With that we can show the user all subscriptions they received.
		{
			err = r.red.Sorted().Create().Score(subRec(inp[i].Rcvr), inp[i].Subs.String(), inp[i].Subs.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the user specific mappings for user specific search queries. With
		// that we can show the user all subscriptions they created. That is usually
		// the subscriptions they paid for.
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

// subRen expresses whether the current subscription object is effectively a
// renewal based on the existing subscriptions for the given receiver, if any.
// Essentially a subscription for the immediate prior month must exist for the
// current subscription to be classified as renewal.
func subRen(exi []*Object, cur *Object) bool {
	if len(exi) == 0 {
		return false
	}

	var des time.Time
	{
		des = cur.Unix.AddDate(0, -1, 0)
	}

	for _, x := range exi {
		if x.Unix.Equal(des) {
			return true
		}
	}

	return false
}

func vldAdd(vld []bool) bool {
	for _, x := range vld {
		if !x {
			return false
		}
	}

	return true
}
