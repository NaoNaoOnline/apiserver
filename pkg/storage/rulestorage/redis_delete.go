package rulestorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Delete(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		// Delete the the user specific mappings for user specific search queries.
		{
			err = r.red.Sorted().Delete().Score(rulUse(inp[i].User), inp[i].Rule.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the the list specific mappings for list specific search queries.
		{
			err = r.red.Sorted().Delete().Score(rulLis(inp[i].List), inp[i].Rule.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Since the deletion process starts with the normalized key-value pair in
		// the handler, we delete it as the very last step, so the operation can
		// eventually be retried.
		{
			_, err = r.red.Simple().Delete().Multi(rulObj(inp[i].Rule))
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
