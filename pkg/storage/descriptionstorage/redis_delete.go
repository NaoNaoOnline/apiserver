package descriptionstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) DeleteDesc(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		// Delete the the user specific mappings for user specific search queries.
		{
			err = r.red.Sorted().Delete().Score(desUse(inp[i].User), inp[i].Desc.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the the event specific mappings for event specific search queries.
		{
			err = r.red.Sorted().Delete().Score(desEve(inp[i].Evnt), inp[i].Desc.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Since the deletion process starts with the normalized key-value pair in
		// the handler, we delete it as the very last step, so the operation can
		// eventually be retried. Here we also delete the description ID mappings
		// tracking all users that have reacted to it in one go.
		{
			lik := likDes(inp[i].Desc)
			obj := desObj(inp[i].Desc)

			_, err = r.red.Simple().Delete().Multi(lik, obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			out = append(out, objectstate.Deleted)
		}
	}

	return out, nil
}

func (r *Redis) DeleteLike(des objectid.ID, use []objectid.ID) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range use {
		{
			_, err = r.red.Simple().Delete().Multi(likMap(use[i], des))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			err = r.red.Sorted().Delete().Score(likUse(use[i]), des.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			out = append(out, objectstate.Deleted)
		}
	}

	return out, nil
}

func (r *Redis) DeleteWrkr(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		// Before deleting a nested structure, we need to create a worker task for
		// ensuring the deletion of the description object and all of its associated
		// data structures.
		{
			err = r.emi.DeleteDesc(inp[i].Desc)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Mark the description object as deleted.
		{
			inp[i].Dltd = time.Now().UTC()
		}

		// Update the description object with the deletion timestamp.
		{
			err = r.red.Simple().Create().Element(desObj(inp[i].Desc), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			out = append(out, objectstate.Deleted)
		}
	}

	return out, nil
}
