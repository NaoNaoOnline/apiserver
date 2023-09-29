package votestorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Delete(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		// Delete the user/event specific mappings for user/event specific search
		// queries.
		{
			err = r.red.Sorted().Delete().Score(votEve(inp[i].User, inp[i].Evnt), inp[i].Vote.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the user specific mappings for user specific search queries. Note
		// that the user specific mapping is created via Sorted.Create.Score, using
		// the event ID as score. Here we want to remove a single specific vote
		// object reference. So we use Sorted.Delete.Value to remove a single vote
		// from the given event. Otherwise we would remove all vote object
		// references from an event.
		{
			err = r.red.Sorted().Delete().Value(votUse(inp[i].User), inp[i].Vote.String())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the vote description mappings for vote description search queries.
		{
			err = r.red.Sorted().Delete().Score(votDes(inp[i].Desc), inp[i].Vote.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Since the deletion process starts with the normalized key-value pair in
		// the handler, we delete it as the very last step, so the operation can
		// eventually be retried.
		{
			_, err = r.red.Simple().Delete().Multi(votObj(inp[i].Vote))
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
