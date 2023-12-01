package subscriptionstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) VerifyAddr(add []string) ([]bool, error) {
	var err error

	// Use the wallet storage to search for the respective user IDs, given a list
	// of wallet addresses.
	var uid []objectid.ID
	{
		uid, _, err = r.wal.SearchAddr(add)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Just verify using VerifyUser, given our user IDs.
	var vld []bool
	{
		vld, err = r.VerifyUser(uid)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return vld, nil
}

func (r *Redis) VerifyUser(uid []objectid.ID) ([]bool, error) {
	var err error

	var vld []bool
	for i := range uid {
		// Use the paird user ID and wallet ID to lookup the matching wallet object.
		var wob walletstorage.Slicer
		{
			wob, err = r.wal.SearchKind(uid[i], []string{"eth"})
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// If the given user owns a wallet object designated for accounting, then
		// continue. Otherwise the respective user ID is not considered a legitimate
		// content creator.
		if len(wob.Labl(objectlabel.WalletAccounting)) == 0 {
			vld = append(vld, false)
			continue
		}

		// Use the event storage to search for the events created by the current
		// wallet owner. The given creator address must be owned by the same user
		// who created useful content.
		var eob eventstorage.Slicer
		{
			eob, err = r.eve.SearchUser([]objectid.ID{uid[i]})
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// If the wallet owner created more than X events, then continue. Otherwise
		// the respective wallet address is not considered a legitimate content
		// creator.
		if len(eob) < minEve {
			vld = append(vld, false)
			continue
		}

		// If the wallet owner generated more than Y link clicks, then continue.
		// Otherwise the respective wallet address is not considered a legitimate
		// content creator.
		if eob.Clck() < minLin {
			vld = append(vld, false)
			continue
		}

		// At this point all our criteria are met for the given wallet address to be
		// considered a legitimate content creator.
		{
			vld = append(vld, true)
		}
	}

	return vld, nil
}
