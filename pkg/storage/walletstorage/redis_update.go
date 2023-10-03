package walletstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Update(inp []*Object) ([]*Object, []objectstate.String, error) {
	var err error

	var sta []objectstate.String
	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// ensure whether the provided wallet signature is in fact valid. For
		// instance, we cannot update a wallet for an user that is not owned by that
		// user.
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, nil, tracer.Mask(err)
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
		// here. Note this test must also be done in Redis.Create.
		if !inp[i].Mestim().Add(5 * time.Minute).After(now) {
			return nil, nil, tracer.Mask(walletSignTooOldError)
		}

		{
			inp[i].Addr.Time = now
		}

		// Once we know the wallet's signature is valid, we update the normalized
		// key-value pair so that we can reflect the wallet object's timestamp
		// change.
		{
			err = r.red.Simple().Create().Element(walObj(inp[i].User, inp[i].Wllt), musStr(inp[i]))
			if err != nil {
				return nil, nil, tracer.Mask(err)
			}
		}

		{
			sta = append(sta, objectstate.Updated)
		}
	}

	return inp, sta, nil
}
