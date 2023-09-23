package walletstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Update(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// ensure whether the provided wallet signature is in fact valid. For
		// instance, we cannot update a wallet for an user that is not owned by that
		// user.
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			inp[i].Last = time.Now().UTC()
		}

		// Once we know the wallet's signature is valid, we update the normalized
		// key-value pair so that we can reflect the wallet object's timestamp
		// change.
		{
			err = r.red.Simple().Create().Element(walObj(inp[i].User, inp[i].Wllt), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			out = append(out, objectstate.Updated)
		}
	}

	return out, nil
}
