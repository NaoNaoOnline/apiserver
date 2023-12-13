package userstorage

import (
	"encoding/json"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) UpdateObct(inp []*Object) ([]objectstate.String, error) {
	var err error

	var sta []objectstate.String
	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// whether the user name complies with the expected format.
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Once we know the user object is valid, we update the normalized key-value
		// pair so that we can reflect the user object's internal change.
		{
			err = r.red.Simple().Create().Element(useObj(inp[i].User), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			sta = append(sta, objectstate.Updated)
		}
	}

	return sta, nil
}

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
			if pat.RplHom(i) {
				obj[i].Home.Time = now
			}

			if pat.RplNam(i) {
				obj[i].Name.Time = now
			}

			if pat.RplNam(i) {
				obj[i].Prfl.Time = now
			}
		}

		var dec jsonpatch.Patch
		{
			dec, err = jsonpatch.DecodePatch(musByt(pat[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Now apply the valid JSON-Patches to the internal user object.
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

		// Verify the modified user object to ensure the applied changes are not
		// rendering it invalid.
		{
			err := upd.Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Once we know the modified user object is still valid after applying the
		// JSON-Patch, we update its normalized key-value pair. Note that we use the
		// resource IDs from the given input in order to construct the storage key
		// for the user object. This input data should come from our internal
		// storage. If we were to use the updated state to construct the storage
		// keys, and if the input validation were to fail for any reason, a
		// potential attack vector would open, because an attacker could choose to
		// overwrite any user object.
		{
			err = r.red.Simple().Create().Element(useObj(obj[i].User), musStr(upd))
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
