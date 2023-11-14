package walletstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreateXtrn(inp []*Object) ([]*Object, error) {
	var err error

	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// ensure whether the provided wallet signature is in fact valid. For
		// instance, we cannot create a wallet for an user that is not owned by that
		// user.
		{
			err = inp[i].VerifyObct()
			if err != nil {
				return nil, tracer.Mask(err)
			}
			err = inp[i].VerifySign()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var now time.Time
		{
			now = time.Now().UTC()
		}

		// We want to ensure that the given signature got generated recently. That
		// way we can prevent somebody from reusing an old signature that may not
		// even belong to the caller. For Object.Verify not to be overloaded with
		// current time dynamics, we test for signature recency in a separate step
		// here. Note this test must also be done in Redis.Update.
		if !inp[i].Mestim().Add(5 * time.Minute).After(now) {
			return nil, tracer.Mask(walletSignTooOldError)
		}

		// Ensure the user wallet limit globally is respected.
		{
			cou, err := r.red.Sorted().Metric().Count(walUse(inp[i].User))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			if cou >= 5 {
				return nil, tracer.Mask(walletUserLimitError)
			}
		}

		{
			inp[i].Addr = objectfield.String{
				Data: inp[i].Comadd().Hex(),
				Time: now,
			}
			inp[i].Crtd = now
			inp[i].Wllt = objectid.Random(objectid.Time(now))
		}

		// Once we know the wallet's signature is valid, we create the normalized
		// key-value pair so that we can search for wallet objects using their IDs.
		{
			err = r.red.Simple().Create().Element(walObj(inp[i].User, inp[i].Wllt), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the wallet address mappings for wallet address search queries.
		{
			err = r.red.Simple().Create().Element(walAdd(inp[i].Addr.Data), inp[i].User.String())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Now we create the wallet kind mappings for wallet kind search queries.
		// With that we can search for wallets of a given kind.
		{
			err = r.red.Sorted().Create().Score(walKin(inp[i].User, inp[i].Kind), inp[i].Wllt.String(), inp[i].Wllt.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the user specific mappings for user specific search queries. With
		// that we can show the user all wallets they created.
		{
			err = r.red.Sorted().Create().Score(walUse(inp[i].User), inp[i].Wllt.String(), inp[i].Wllt.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}
