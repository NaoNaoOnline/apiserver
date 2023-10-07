package descriptionstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/rescue/task"
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
		// eventually be retried.
		{
			_, err = r.red.Simple().Delete().Multi(desObj(inp[i].Desc))
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
		var tas *task.Task
		{
			tas = &task.Task{
				Meta: &task.Meta{
					objectlabel.DescAction: objectlabel.ActionDelete,
					objectlabel.DescObject: inp[i].Desc.String(),
					objectlabel.DescOrigin: objectlabel.OriginCustom,
				},
			}
		}

		// Submit the task to the worker queue.
		{
			err = r.res.Create(tas)
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
