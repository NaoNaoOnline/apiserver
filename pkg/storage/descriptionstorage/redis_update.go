package descriptionstorage

import (
	"encoding/json"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) UpdateLike(use objectid.ID, obj []*Object, inc []bool) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range obj {
		// Ensure that descriptions can only be liked and unliked once by the same user.
		if obj[i].Like.User && inc[i] {
			return nil, tracer.Mask(descriptionLikeAlreadyExistsError)
		}
		if !obj[i].Like.User && !inc[i] {
			return nil, tracer.Mask(descriptionUnlikeAlreadyExistsError)
		}

		// Track the new like or unlike on the description object by incrementing or
		// decrementing its internal counter.
		if inc[i] {
			obj[i].Like.Data++
		} else {
			obj[i].Like.Data--
		}

		// Track the time of the last updated like.
		{
			obj[i].Like.Time = time.Now().UTC()
		}

		// Verify the modified description object to ensure the applied changes are
		// not rendering it invalid.
		{
			err := obj[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Once we know the modified description object is still valid after
		// tracking the new like, we update its normalized key-value pair.
		{
			err = r.red.Simple().Create().Element(desObj(obj[i].Desc), musStr(obj[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Persist the like indicator for the calling user and the given description
		// ID, if a like happened. Othwerwise remove the like indicator.
		if inc[i] {
			err = r.red.Simple().Create().Element(desLik(use, obj[i].Desc), "1")
			if err != nil {
				return nil, tracer.Mask(err)
			}
		} else {
			_, err = r.red.Simple().Delete().Multi(desLik(use, obj[i].Desc))
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

func (r *Redis) UpdatePtch(obj []*Object, pat [][]*Patch) ([]objectstate.String, error) {
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

		var dec jsonpatch.Patch
		{
			dec, err = jsonpatch.DecodePatch(musByt(pat[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Now apply the valid JSON-Patches to the internal description object.
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
