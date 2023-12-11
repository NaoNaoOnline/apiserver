package descriptionstorage

import (
	"encoding/json"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) UpdateLike(use objectid.ID, pre bool, obj []*Object, inc []bool) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range obj {
		// Ensure that descriptions can only be liked and unliked once by the same user.
		if obj[i].Mtrc.User[objectlabel.DescriptionMetricUser] && inc[i] {
			return nil, tracer.Mask(descriptionLikeAlreadyExistsError)
		}
		if !obj[i].Mtrc.User[objectlabel.DescriptionMetricUser] && !inc[i] {
			return nil, tracer.Mask(descriptionUnlikeAlreadyExistsError)
		}

		// Only track premium comment likes if the given user has a premium
		// subscription.
		if pre {
			if inc[i] {
				obj[i].Mtrc.Data[objectlabel.DescriptionMetricPrem]++
			} else {
				obj[i].Mtrc.Data[objectlabel.DescriptionMetricPrem]--
			}
		}

		// Always track user comment likes on the description object by incrementing
		// its public internal counter.
		if inc[i] {
			obj[i].Mtrc.Data[objectlabel.DescriptionMetricUser]++
		} else {
			obj[i].Mtrc.Data[objectlabel.DescriptionMetricUser]--
		}

		// Track the time of the last updated like.
		{
			obj[i].Mtrc.Time = time.Now().UTC()
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

		// Persist the like indicators for the calling user and the given description
		// ID, if a like happened. Othwerwise remove the like indicators.
		if inc[i] {
			// We use a simple key-value pair for a user-description relationship so
			// we can lookup all the likes a user made on a list of descriptions. This
			// internal data structure is used in the Description.Search endpoints.
			{
				err = r.red.Simple().Create().Element(likMap(use, obj[i].Desc), "1")
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			// We use a sorted set to store all the user IDs that have reacted to a
			// particular description in the form of a like. This internal data
			// structure is used to find and cleanup the other keys and values that we
			// use for tracking user likes on descriptions.
			{
				err = r.red.Sorted().Create().Score(likDes(obj[i].Desc), use.String(), use.Float())
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			// We use a sorted set for all the events that a user reacted to in the
			// form of a description like. This internal data structure is used in
			// Event.SearchLike and Event.SearchRule. Note that there should not be a
			// need to verify the integrity of the input object's event-description
			// relationship, because the RPC update handler should only provide data
			// from our internal storage, which should always be properly persisted.
			{
				err = r.red.Sorted().Create().Score(likUse(use), objectid.Pair(obj[i].Evnt, obj[i].Desc), obj[i].Desc.Float())
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}
		} else {
			// Just reverse the operation from above.
			{
				_, err = r.red.Simple().Delete().Multi(likMap(use, obj[i].Desc))
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			// Just reverse the operation from above.
			{
				err = r.red.Sorted().Delete().Score(likDes(obj[i].Desc), use.Float())
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			// Just reverse the operation from above.
			{
				err = r.red.Sorted().Delete().Score(likUse(use), obj[i].Desc.Float())
				if err != nil {
					return nil, tracer.Mask(err)
				}
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

		var now time.Time
		{
			now = time.Now().UTC()
		}

		{
			obj[i].Text.Time = now
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

		// Once we know the modified description object is still valid after
		// applying the JSON-Patch, we update its normalized key-value pair. Note
		// that we use the resource IDs from the given input in order to construct
		// the storage key for the description object. This input data should come
		// from our internal storage. If we were to use the updated state to
		// construct the storage keys, and if the input validation were to fail for
		// any reason, a potential attack vector would open, because an attacker
		// could choose to overwrite any description object.
		{
			err = r.red.Simple().Create().Element(desObj(obj[i].Desc), musStr(upd))
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
