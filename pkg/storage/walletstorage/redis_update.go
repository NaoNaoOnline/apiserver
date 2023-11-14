package walletstorage

import (
	"encoding/json"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) UpdatePtch(obj []*Object, pat PatchSlicer) ([]*Object, []objectstate.String, error) {
	var err error

	var out []*Object
	var sta []objectstate.String
	for i := range obj {
		// At first we need to validate the given JSON-Patches and ensure the
		// desired modifications are permitted at all.
		for _, x := range pat[i] {
			err := x.Verify()
			if err != nil {
				return nil, nil, tracer.Mask(err)
			}
		}

		var now time.Time
		{
			now = time.Now().UTC()
		}

		{
			if pat.AddLab(i, "*") {
				obj[i].Labl.Time = now
			}

			if pat.RemLab(i, "*") {
				obj[i].Labl.Time = now
			}
		}

		var dec jsonpatch.Patch
		{
			dec, err = jsonpatch.DecodePatch(musByt(pat[i]))
			if err != nil {
				return nil, nil, tracer.Mask(err)
			}
		}

		// Now apply the valid JSON-Patches to the internal wallet object.
		var byt []byte
		{
			byt, err = dec.Apply([]byte(musStr(obj[i])))
			if err != nil {
				return nil, nil, tracer.Mask(err)
			}
		}

		var upd *Object
		{
			upd = &Object{}
		}

		{
			err = json.Unmarshal(byt, upd)
			if err != nil {
				return nil, nil, tracer.Mask(err)
			}
		}

		// Verify the modified wallet object to ensure the applied changes are not
		// rendering it invalid.
		{
			err = upd.VerifyObct()
			if err != nil {
				return nil, nil, tracer.Mask(err)
			}
			err = upd.VerifyPtch()
			if err != nil {
				return nil, nil, tracer.Mask(err)
			}
		}

		// Once we know the modified wallet object is still valid after applying the
		// JSON-Patch, we update its normalized key-value pair. Note that we use the
		// resource IDs from the given input in order to construct the storage key
		// for the wallet object. This input data should come from our internal
		// storage. If we were to use the updated state to construct the storage
		// keys, and if the input validation were to fail for any reason, a
		// potential attack vector would open, because an attacker could choose to
		// overwrite any wallet object.
		{
			err = r.red.Simple().Create().Element(walObj(obj[i].User, obj[i].Wllt), musStr(upd))
			if err != nil {
				return nil, nil, tracer.Mask(err)
			}
		}

		{
			sta = append(sta, objectstate.Updated)
			out = append(out, upd)
		}
	}

	return out, sta, nil
}

func (r *Redis) UpdateSign(inp []*Object) ([]*Object, []objectstate.String, error) {
	var err error

	var sta []objectstate.String
	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// ensure whether the provided wallet signature is in fact valid. For
		// instance, we cannot update a wallet for an user that is not owned by that
		// user.
		{
			err = inp[i].VerifyObct()
			if err != nil {
				return nil, nil, tracer.Mask(err)
			}
			err = inp[i].VerifySign()
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
