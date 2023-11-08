package eventstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) DeleteEvnt(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		// Delete the the user specific mappings for user specific search queries.
		{
			err = r.red.Sorted().Delete().Score(eveUse(inp[i].User), inp[i].Evnt.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the label specific mappings for label specific search queries.
		for _, x := range append(inp[i].Cate, inp[i].Host...) {
			err = r.red.Sorted().Delete().Score(eveLab(x), inp[i].Evnt.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the time specific mappings for time specific search queries. Note
		// that the time specific mapping is created via Sorted.Create.Score, using
		// the event time as score. Here we want to remove a single specific event
		// object reference. So we use Sorted.Delete.Value to remove a single event
		// from the given list. Otherwise we would remove all event object
		// references happening at the same time.
		{
			err = r.red.Sorted().Delete().Value(keyfmt.EventTime, inp[i].Evnt.String())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Since the deletion process starts with the normalized key-value pair in
		// the handler, we delete it as the very last step, so the operation can
		// eventually be retried.
		{
			clk := clkEve(inp[i].Evnt)
			obj := eveObj(inp[i].Evnt)

			_, err = r.red.Simple().Delete().Multi(clk, obj)
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
		// ensuring the deletion of the event object and all of its associated data
		// structures.
		{
			err = r.emi.DeleteEvnt(inp[i].Evnt)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Mark the event object as deleted.
		{
			inp[i].Dltd = time.Now().UTC()
		}

		// Update the event object with the deletion timestamp.
		{
			err = r.red.Simple().Create().Element(eveObj(inp[i].Evnt), musStr(inp[i]))
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
