package subscriptionstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

const (
	minEve = 3
	minLin = 3
)

func (r *Redis) VerifyAddr(add []string) ([]bool, error) {
	var err error

	// Use the wallet storage to search for the respective user and wallet IDs,
	// given a list of wallet addresses.
	var uid []objectid.ID
	var wid []objectid.ID
	{
		uid, wid, err = r.wal.SearchAddr(add)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Some validation to be super defensive.
	{
		if len(uid) != len(add) {
			return nil, tracer.Maskf(runtime.ExecutionFailedError, "%v", add)
		}
	}

	var vld []bool
	for i := range uid {
		// Use the paird user ID and wallet ID to lookup the matching wallet object.
		var wob []*walletstorage.Object
		{
			wob, err = r.wal.SearchWllt(uid[i], []objectid.ID{wid[i]})
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Some validation to be super defensive.
		{
			if len(wob) != 1 {
				return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d", len(wob))
			}
			if wob[0].Addr.Data != add[i] {
				return nil, tracer.Maskf(runtime.ExecutionFailedError, "%s != %s", wob[0].Addr.Data, add[i])
			}
		}

		// If the wallet object is labelled to be used for accounting, then
		// continue. Otherwise the respective wallet address is not considered a
		// legitimate content creator.
		if !wob[0].HasLab(objectlabel.WalletAccounting) {
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
