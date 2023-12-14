package subscriptionstorage

import (
	"context"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
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

		var now time.Time
		{
			now = time.Now().UTC()
		}

		// Check whether the created subscription for the given receiver is
		// effectively a renewal.
		if subRen(exi, inp[i].Unix, now) {
			{
				r.log.Log(
					context.Background(),
					"level", "info",
					"message", "processing subscription renewal",
					objectlabel.StrgObject, object.String(r),
				)
			}

			{
				err = inp[i].VerifyUnix(VerifyRenw(now))
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}
		} else {
			if len(exi) == 0 {
				{
					r.log.Log(
						context.Background(),
						"level", "info",
						"message", "processing first subscription",
						objectlabel.StrgObject, object.String(r),
					)
				}

				// The very first subscription month is for free. So we add 1 month to
				// the current time that is being used to verify the given subscription
				// period.
				{
					err = inp[i].VerifyUnix(VerifyOnce(now.AddDate(0, 1, 0)))
					if err != nil {
						return nil, tracer.Mask(err)
					}
				}
			} else {
				{
					r.log.Log(
						context.Background(),
						"level", "info",
						"message", "processing single subscription",
						objectlabel.StrgObject, object.String(r),
					)
				}

				// Single subscriptions represent those that are neither renewals nor
				// the very first subscriptions. If this code is executed, then the user
				// already had a subscription, then had none, and now wants one again.
				{
					err = inp[i].VerifyUnix(VerifyOnce(now))
					if err != nil {
						return nil, tracer.Mask(err)
					}
				}
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

// subRen expresses whether the current subscription object is effectively a
// renewal based on the existing subscriptions for the given receiver, if any.
// Essentially a subscription for the immediate prior month must exist for the
// current subscription to be classified as renewal.
func subRen(exi []*Object, uni time.Time, now time.Time) bool {
	if len(exi) == 0 {
		return false
	}

	var sta time.Time
	var end time.Time
	{
		sta = timMon(now).AddDate(0, 1, -7)
		end = timMon(now).AddDate(0, 1, 0)
	}

	if now.Before(sta) || now.After(end) {
		return false
	}

	var des time.Time
	{
		des = uni.AddDate(0, -1, 0)
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
