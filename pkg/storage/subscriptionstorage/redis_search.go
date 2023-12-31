package subscriptionstorage

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/redigo/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchCrtr(uid []objectid.ID) ([]*walletstorage.Object, error) {
	var err error

	var wob walletstorage.Slicer
	for i := range uid {
		// Use the user storage to search for event IDs that the given user ID
		// reacted to in the form of a link click.
		var eid []objectid.ID
		{
			eid, err = r.use.SearchLink([]objectid.ID{uid[i]})
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// There might not be any event IDs, and so we do not proceed, but instead
		// continue with the next user ID, if any.
		if len(eid) == 0 {
			continue
		}

		// Find all event objects for the respective event IDs.
		var eob eventstorage.Slicer
		{
			eob, err = r.eve.SearchEvnt("", eid)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Filter the given event objects and only keep those matching our
		// requirements.
		var fil eventstorage.Slicer
		{
			// First, group the event objects by the users who created them.
			dic := map[objectid.ID]eventstorage.Slicer{}
			for _, x := range eob {
				dic[x.User] = append(dic[x.User], x)
			}

			for _, x := range dic {
				// If a user created more than X events, then continue. Otherwise the
				// respective user is not considered a legitimate content creator.
				if len(x) < r.mse {
					continue
				}

				// If the user generated more than Y link clicks, then continue.
				// Otherwise the respective user is not considered a legitimate content
				// creator.
				if x.Mtrc(objectlabel.EventMetricUser) < int64(r.msl) {
					continue
				}

				// At this point all our criteria are met for the given user to be
				// considered a legitimate content creator.
				{
					fil = append(fil, x...)
				}
			}
		}

		// There might not be any event objects, and so we do not proceed, but
		// instead continue with the next user ID, if any.
		if len(fil) == 0 {
			continue
		}

		// Sort the filtered list of events by link clicks in order to promote
		// content creators that generate engagement.
		sort.SliceStable(fil, func(i, j int) bool {
			return fil[i].Mtrc.Data[objectlabel.EventMetricUser] > fil[j].Mtrc.Data[objectlabel.EventMetricUser]
		})

		// For the remaining content creators, lookup their accounting wallets.
		for _, x := range generic.Uni(fil.User()) {
			var lis []*walletstorage.Object
			{
				lis, err = r.wal.SearchKind(x, []string{"eth"})
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			for _, y := range lis {
				if y.HasLab(objectlabel.WalletAccounting) {
					// Track the accounting wallet of each content creator.
					{
						wob = append(wob, y)
					}

					// We can break here since there should only ever be a single
					// accounting wallet.
					{
						break
					}
				}
			}
		}
	}

	return wob, nil
}

func (r *Redis) SearchCurr(uid objectid.ID) (*Object, error) {
	var err error

	// Search for all subscription objects for the given user ID, which is
	// expected to be the receiver of the subscription to search for.
	var sob []*Object
	{
		sob, err = r.SearchRcvr([]objectid.ID{uid}, PagAll())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// It might very well be that there is no subscription for any given user. In
	// that case we return nil.
	if len(sob) == 0 {
		return nil, nil
	}

	// Sort subscription objects by subscription period in descending order with
	// first priority. The newest subscriptions will end up at index zero.
	sort.SliceStable(sob, func(i, j int) bool {
		return sob[i].Unix.Unix() > sob[j].Unix.Unix()
	})

	var now time.Time
	{
		now = time.Now().UTC()
	}

	// Verify whether the latest subscription for the given user is in fact for
	// the current, or next month. We expect the current month for single
	// subscriptions. We expect the next month for first subscriptions and
	// subscription renewals. VerifyUnix together with VerifyOnce will fail if the
	// subscription timestamp does not refer to the first day of the given month.
	// So if that specific error is returned, we did not find the valid latest
	// subscription. And in that case we simply return nil.
	if len(sob) == 1 || subRen(sob, sob[0].Unix, now) {
		err = sob[0].VerifyUnix(VerifyOnce(now.AddDate(0, 1, 0)))
		if IsSubscriptionUnixInvalid(err) {
			return nil, nil
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	} else {
		err = sob[0].VerifyUnix(VerifyOnce(now))
		if IsSubscriptionUnixInvalid(err) {
			return nil, nil
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// The validation above worked out, which means for us that we have found the
	// valid latest subscription.
	return sob[0], nil
}

func (r *Redis) SearchPayr(use []objectid.ID, pag [2]int) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range use {
		// val will result in a list of all subscription IDs paid for by the given
		// user, if any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(subPay(x), pag[0], pag[1])
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// There might not be any values, and so we do not proceed, but instead
		// continue with the next user ID, if any.
		if len(val) == 0 {
			continue
		}

		{
			lis, err := r.SearchSubs(objectid.IDs(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}

func (r *Redis) SearchRcvr(use []objectid.ID, pag [2]int) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range use {
		// val will result in a list of all subscription IDs received by the given
		// user, if any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(subRec(x), pag[0], pag[1])
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// There might not be any values, and so we do not proceed, but instead
		// continue with the next user ID, if any.
		if len(val) == 0 {
			continue
		}

		{
			lis, err := r.SearchSubs(objectid.IDs(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}

func (r *Redis) SearchSubs(inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(inp, keyfmt.SubscriptionObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(subscriptionObjectNotFoundError, "%v", inp)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out []*Object
	for i := range jsn {
		var obj *Object
		{
			obj = &Object{}
		}

		if jsn[i] != "" {
			err = json.Unmarshal([]byte(jsn[i]), obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			out = append(out, obj)
		}
	}

	return out, nil
}
