package liststorage

import (
	"encoding/json"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) UpdatePtch(obj []*Object, pat PatchSlicer) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range obj {
		// At first we need to validate the given JSON-Patches and ensure the
		// desired modifications are permitted at all.
		for _, x := range pat[i] {
			err := x.Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var now time.Time
		{
			now = time.Now().UTC()
		}

		{
			if pat.RplDes(i) {
				obj[i].Desc.Time = now
			}

			if pat.RplFee(i) {
				obj[i].Feed.Time = now
				// Since the apischema defines time values as strings of unix seconds,
				// and since we accept those values transparently for the JSON patches,
				// we have to transform unix seconds into formatted time strings. If we
				// do not do that, then the JSON patch will fail because Feed.Data is of
				// type time.Time, and that type requires a time formatted string.
				pat[i] = pat.UniTim(i)
			}
		}

		var dec jsonpatch.Patch
		{
			dec, err = jsonpatch.DecodePatch(musByt(pat[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Now apply the valid JSON-Patches to the internal list object.
		var byt []byte
		{
			byt, err = dec.Apply([]byte(musStr(obj[i])))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var upd *Object
		{
			upd = &Object{}
		}

		{
			err = json.Unmarshal(byt, upd)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Verify the modified list object to ensure the applied changes are not
		// rendering it invalid.
		{
			err := upd.Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Once we know the modified list object is still valid after applying the
		// JSON-Patch, we update its normalized key-value pair. Note that we use the
		// resource IDs from the given input in order to construct the storage key
		// for the list object. This input data should come from our internal
		// storage. If we were to use the updated state to construct the storage
		// keys, and if the input validation were to fail for any reason, a
		// potential attack vector would open, because an attacker could choose to
		// overwrite any list object.
		{
			err = r.red.Simple().Create().Element(lisObj(obj[i].List), musStr(upd))
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
