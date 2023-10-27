package liststorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) DeleteList(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		// Delete the the user specific mappings for user specific search queries.
		{
			err = r.red.Sorted().Delete().Score(lisUse(inp[i].User), inp[i].List.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Since the deletion process starts with the normalized key-value pair in
		// the handler, we delete it as the very last step, so the operation can
		// eventually be retried.
		{
			_, err = r.red.Simple().Delete().Multi(lisObj(inp[i].List))
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
		// ensuring the deletion of the list object and all of its associated data
		// structures.
		{
			err = r.emi.Delete(inp[i].List)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Mark the list object as deleted.
		{
			inp[i].Dltd = time.Now().UTC()
		}

		// Update the list object with the deletion timestamp.
		{
			err = r.red.Simple().Create().Element(lisObj(inp[i].List), musStr(inp[i]))
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
