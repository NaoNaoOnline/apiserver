package descriptionstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Update(obj []*Object, pat [][]*Patch) ([]objectstate.String, error) {
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

		// Now apply the valid JSON-Patches to the internal description object.
		var byt []byte
		{
			byt, err = musPat(pat[i]).Apply([]byte(musStr(obj[i])))
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

		// Verify the modified description object to ensure the applied changes are
		// not rendering it invalid.
		{
			err := upd.Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Once we know the modified description object is still valid after applying
		// the JSON-Patch, we update its normalized key-value pair.
		{
			err = r.red.Simple().Create().Element(desObj(upd.Desc), musStr(upd))
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
