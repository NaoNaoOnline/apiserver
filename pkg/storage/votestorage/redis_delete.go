package votestorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Delete(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		var eve *eventstorage.Object
		{
			eve, err = r.searchEvnt(inp[i].Evnt)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Ensure votes cannot be removed from events that have already happened.
		{
			if eve.Happnd() {
				return nil, tracer.Mask(eventAlreadyHappenedError)
			}
		}

		// Delete the user/event specific mappings for user/event specific search
		// queries.
		{
			err = r.red.Sorted().Delete().Value(votEve(inp[i].User, inp[i].Evnt), inp[i].Vote.String())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the user specific mappings for user specific search queries.
		{
			err = r.red.Sorted().Delete().Value(votUse(inp[i].User), inp[i].Vote.String())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the vote description mappings for vote description search queries.
		{
			err = r.red.Sorted().Delete().Value(votDes(inp[i].Desc), inp[i].Vote.String())
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
