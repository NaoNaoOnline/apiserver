package walletstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
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
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			inp[i].Addr = inp[i].Comadd().Hex()
			inp[i].Crtd = time.Now().UTC()
			inp[i].Last = inp[i].Crtd
			inp[i].Wllt = objectid.New(inp[i].Crtd)
		}

		// Once we know the wallet's signature is valid, we create the normalized
		// key-value pair so that we can search for wallet objects using their IDs.
		{
			err = r.red.Simple().Create().Element(walObj(inp[i].User, inp[i].Wllt), musStr(inp[i]))
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
		// that we can show the user all wallets they created. Note that we index
		// the user's wallets by their wallet addresses. With that we can search for
		// a user wallet only knowing their wallet address.
		{
			err = r.red.Sorted().Create().Score(walUse(inp[i].User), inp[i].Wllt.String(), inp[i].Wllt.Float(), keyfmt.Indx(inp[i].Addr))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}
